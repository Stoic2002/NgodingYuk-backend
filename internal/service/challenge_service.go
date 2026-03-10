package service

import (
	"encoding/json"
	"errors"
	"time"

	"github.com/arulkarim/ngodingyuk-server/internal/domain"
	"github.com/arulkarim/ngodingyuk-server/pkg/i18n"
	"github.com/google/uuid"
)

type ChallengeService struct {
	challengeRepo ChallengeRepository
	progressRepo  ProgressRepository
	userRepo      UserRepository
	executeSvc    *ExecuteService
}

func NewChallengeService(
	challengeRepo ChallengeRepository,
	progressRepo ProgressRepository,
	userRepo UserRepository,
	executeSvc *ExecuteService,
) *ChallengeService {
	return &ChallengeService{
		challengeRepo: challengeRepo,
		progressRepo:  progressRepo,
		userRepo:      userRepo,
		executeSvc:    executeSvc,
	}
}

// === Response DTOs ===

type ChallengeListItem struct {
	ID          string `json:"id"`
	Slug        string `json:"slug"`
	Language    string `json:"language"`
	Difficulty  string `json:"difficulty"`
	Title       string `json:"title"`
	XPReward    int64  `json:"xp_reward"`
	OrderIndex  int    `json:"order_index"`
	IsCompleted bool   `json:"is_completed"`
}

type ChallengeDetail struct {
	ID             string          `json:"id"`
	Slug           string          `json:"slug"`
	Language       string          `json:"language"`
	Difficulty     string          `json:"difficulty"`
	Title          string          `json:"title"`
	Story          string          `json:"story"`
	Task           string          `json:"task"`
	Hint           string          `json:"hint,omitempty"`
	SchemaInfo     json.RawMessage `json:"schema_info,omitempty"`
	ExpectedOutput json.RawMessage `json:"expected_output"`
	StarterCode    string          `json:"starter_code,omitempty"`
	TestCases      json.RawMessage `json:"test_cases"`
	XPReward       int64           `json:"xp_reward"`
}

type ChallengeProgressResponse struct {
	Status   string `json:"status"`
	Attempts int    `json:"attempts"`
	SolvedAt string `json:"solved_at,omitempty"`
}

type SubmitResponse struct {
	AllPassed bool        `json:"all_passed"`
	Results   []RunResult `json:"results"`
	XPGained  int64       `json:"xp_gained"`
}

// === Service Methods ===

// List returns a paginated list of challenges, resolved to the given locale.
func (s *ChallengeService) List(language, difficulty, search, locale string, limit, offset int, userID uuid.UUID) ([]ChallengeListItem, int64, error) {
	if limit <= 0 {
		limit = 50
	}
	challenges, total, err := s.challengeRepo.List(language, difficulty, search, limit, offset)
	if err != nil {
		return nil, 0, err
	}

	completedMap := make(map[uuid.UUID]bool)
	if userID != uuid.Nil {
		completedIDs, _ := s.progressRepo.GetCompletedChallengeIDs(userID)
		for _, id := range completedIDs {
			completedMap[id] = true
		}
	}

	var items []ChallengeListItem
	for _, c := range challenges {
		items = append(items, ChallengeListItem{
			ID:          c.ID.String(),
			Slug:        c.Slug,
			Language:    c.Language,
			Difficulty:  c.Difficulty,
			Title:       i18n.Resolve(locale, c.TitleID, c.TitleEN),
			XPReward:    c.XPReward,
			OrderIndex:  c.OrderIndex,
			IsCompleted: completedMap[c.ID],
		})
	}
	return items, total, nil
}

// GetBySlug returns the full detail of a challenge (without solution_code).
func (s *ChallengeService) GetBySlug(slug, locale string) (*ChallengeDetail, error) {
	c, err := s.challengeRepo.FindBySlug(slug)
	if err != nil {
		return nil, errors.New("challenge not found")
	}

	starterCode := ""
	if c.StarterCode != nil {
		starterCode = *c.StarterCode
	}

	return &ChallengeDetail{
		ID:             c.ID.String(),
		Slug:           c.Slug,
		Language:       c.Language,
		Difficulty:     c.Difficulty,
		Title:          i18n.Resolve(locale, c.TitleID, c.TitleEN),
		Story:          i18n.Resolve(locale, c.StoryID, c.StoryEN),
		Task:           i18n.Resolve(locale, c.TaskID, c.TaskEN),
		Hint:           i18n.ResolveOptional(locale, c.HintID, c.HintEN),
		SchemaInfo:     c.SchemaInfo,
		ExpectedOutput: c.ExpectedOutput,
		StarterCode:    starterCode,
		TestCases:      c.TestCases,
		XPReward:       c.XPReward,
	}, nil
}

// Run executes user code without saving progress. Used for the "Run" button.
func (s *ChallengeService) Run(slug string, code string) (*ExecuteResponse, error) {
	c, err := s.challengeRepo.FindBySlug(slug)
	if err != nil {
		return nil, errors.New("challenge not found")
	}

	switch c.Language {
	case "golang", "go":
		return s.executeSvc.RunGo(code, c.TestCases)
	case "sql":
		return s.executeSvc.RunSQL(code, c.SchemaInfo, c.TestCases, c.ExpectedOutput)
	default:
		return nil, errors.New("unsupported language: " + c.Language)
	}
}

// Submit executes user code AND saves progress + awards XP if all passed.
func (s *ChallengeService) Submit(slug string, userID uuid.UUID, code string) (*SubmitResponse, error) {
	c, err := s.challengeRepo.FindBySlug(slug)
	if err != nil {
		return nil, errors.New("challenge not found")
	}

	// Execute code
	var execResult *ExecuteResponse
	switch c.Language {
	case "golang", "go":
		execResult, err = s.executeSvc.RunGo(code, c.TestCases)
	case "sql":
		execResult, err = s.executeSvc.RunSQL(code, c.SchemaInfo, c.TestCases, c.ExpectedOutput)
	default:
		return nil, errors.New("unsupported language: " + c.Language)
	}
	if err != nil {
		return nil, err
	}

	// Update progress
	progress, err := s.progressRepo.FindOrCreateChallengeProgress(userID, c.ID)
	if err != nil {
		return nil, err
	}

	progress.Attempts++
	progress.LastCode = &code

	var xpGained int64

	if execResult.AllPassed && progress.Status != "solved" {
		progress.Status = "solved"
		now := time.Now()
		progress.SolvedAt = &now
		xpGained = c.XPReward

		// Award XP
		user, err := s.userRepo.FindByID(userID)
		if err == nil {
			user.XP += xpGained
			user.Level = CalculateLevel(user.XP)
			UpdateStreak(user)
			s.userRepo.Update(user)

			s.progressRepo.CreateXPHistory(&domain.UserXPHistory{
				UserID:     userID,
				XPGained:   xpGained,
				SourceType: "challenge",
				SourceID:   &c.ID,
			})
		}
	} else if !execResult.AllPassed {
		progress.Status = "attempted"
	}

	s.progressRepo.UpdateChallengeProgress(progress)

	return &SubmitResponse{
		AllPassed: execResult.AllPassed,
		Results:   execResult.Results,
		XPGained:  xpGained,
	}, nil
}

// GetMyProgress returns the user's progress on a specific challenge.
func (s *ChallengeService) GetMyProgress(slug string, userID uuid.UUID) (*ChallengeProgressResponse, error) {
	c, err := s.challengeRepo.FindBySlug(slug)
	if err != nil {
		return nil, errors.New("challenge not found")
	}

	progress, err := s.progressRepo.FindChallengeProgress(userID, c.ID)
	if err != nil {
		return &ChallengeProgressResponse{
			Status:   "not_started",
			Attempts: 0,
		}, nil
	}

	resp := &ChallengeProgressResponse{
		Status:   progress.Status,
		Attempts: progress.Attempts,
	}
	if progress.SolvedAt != nil {
		resp.SolvedAt = progress.SolvedAt.Format(time.RFC3339)
	}
	return resp, nil
}
