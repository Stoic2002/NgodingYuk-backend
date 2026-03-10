package service

import (
	"encoding/json"
	"errors"
	"time"

	"github.com/arulkarim/ngodingyuk-server/internal/domain"
	"github.com/google/uuid"
)

// Define interface for repositories used by UserService to follow Dependency Injection principles.
type UserRepo interface {
	FindByID(id uuid.UUID) (*domain.User, error)
	FindByUsername(username string) (*domain.User, error)
	UpdateFields(id uuid.UUID, fields map[string]interface{}) error
}

type ProgressRepo interface {
	GetUserXPHistory(userID uuid.UUID, limit, offset int) ([]domain.UserXPHistory, error)
	GetUserChallengeStats(userID uuid.UUID) ([]struct {
		Difficulty string
		Status     string
		Count      int64
	}, error)
	GetUserQuizHistory(userID uuid.UUID) ([]domain.QuizHistoryItem, error)
}

type CourseRepo interface {
	GetUserCertificates(userID uuid.UUID) ([]domain.Certificate, error)
}

// Since domain models are returned, we imported domain at the top.
// We will refine these interfaces and models as we implement.

type UserService struct {
	userRepo     UserRepo
	progressRepo ProgressRepo
	courseRepo   CourseRepo
}

func NewUserService(userRepo UserRepo, progressRepo ProgressRepo, courseRepo CourseRepo) *UserService {
	return &UserService{
		userRepo:     userRepo,
		progressRepo: progressRepo,
		courseRepo:   courseRepo,
	}
}

// === Request/Response Types ===

type UpdateUserReq struct {
	Username *string `json:"username"`
	Locale   *string `json:"locale"`
}

type XPItemResp struct {
	ID         string    `json:"id"`
	XPGained   int64     `json:"xp_gained"`
	SourceType string    `json:"source_type"`
	SourceID   string    `json:"source_id,omitempty"`
	CreatedAt  time.Time `json:"created_at"`
}

type CertItemResp struct {
	ID             string    `json:"id"`
	CourseID       string    `json:"course_id"`
	CourseTitle    string    `json:"course_title"`
	CourseSlug     string    `json:"course_slug"`
	Score          int       `json:"score"`
	TotalQuestions int       `json:"total_questions"`
	PassedAt       time.Time `json:"passed_at"`
}

type QuizHistoryItemResp struct {
	Type           string          `json:"type"`
	CourseTitle    string          `json:"course_title"`
	LessonTitle    string          `json:"lesson_title,omitempty"`
	Score          int             `json:"score"`
	TotalQuestions int             `json:"total_questions"`
	ResultDetails  json.RawMessage `json:"result_details,omitempty"`
	PassedAt       time.Time       `json:"passed_at"`
}

// === Service Methods ===

func (s *UserService) GetProfile(userID uuid.UUID) (*UserResponse, error) {
	user, err := s.userRepo.FindByID(userID)
	if err != nil {
		return nil, errors.New("user not found")
	}

	resp := UserResponse{
		ID:          user.ID.String(),
		Username:    user.Username,
		Email:       user.Email,
		XP:          user.XP,
		Level:       user.Level,
		StreakCount: user.StreakCount,
		Locale:      user.Locale,
	}
	return &resp, nil
}

func (s *UserService) GetXPHistory(userID uuid.UUID, limit, offset int) ([]XPItemResp, error) {
	history, err := s.progressRepo.GetUserXPHistory(userID, limit, offset)
	if err != nil {
		return nil, err
	}

	var items []XPItemResp
	for _, h := range history {
		item := XPItemResp{
			ID:         h.ID.String(),
			XPGained:   h.XPGained,
			SourceType: h.SourceType,
			CreatedAt:  h.CreatedAt,
		}
		if h.SourceID != nil {
			item.SourceID = h.SourceID.String()
		}
		items = append(items, item)
	}

	return items, nil
}

func (s *UserService) GetChallengeStats(userID uuid.UUID) (map[string]map[string]int64, error) {
	stats, err := s.progressRepo.GetUserChallengeStats(userID)
	if err != nil {
		return nil, err
	}

	// Organize by difficulty
	result := map[string]map[string]int64{
		"easy":   {"solved": 0, "attempted": 0},
		"medium": {"solved": 0, "attempted": 0},
		"hard":   {"solved": 0, "attempted": 0},
	}
	for _, stat := range stats {
		if _, ok := result[stat.Difficulty]; ok {
			result[stat.Difficulty][stat.Status] = stat.Count
		}
	}

	return result, nil
}

func (s *UserService) UpdateProfile(userID uuid.UUID, req UpdateUserReq) (*UserResponse, error) {
	fields := make(map[string]interface{})

	if req.Username != nil && *req.Username != "" {
		// Check uniqueness
		existing, _ := s.userRepo.FindByUsername(*req.Username)
		if existing != nil && existing.ID != userID {
			return nil, errors.New("username already taken")
		}
		fields["username"] = *req.Username
	}

	if req.Locale != nil {
		if *req.Locale != "id" && *req.Locale != "en" {
			return nil, errors.New("locale must be 'id' or 'en'")
		}
		fields["locale"] = *req.Locale
	}

	if len(fields) == 0 {
		return nil, errors.New("no fields to update")
	}

	if err := s.userRepo.UpdateFields(userID, fields); err != nil {
		return nil, errors.New("failed to update profile")
	}

	// Return updated user
	return s.GetProfile(userID)
}

func (s *UserService) GetCertificates(userID uuid.UUID) ([]CertItemResp, error) {
	certs, err := s.courseRepo.GetUserCertificates(userID)
	if err != nil {
		return nil, err
	}

	var items []CertItemResp
	for _, cert := range certs {
		items = append(items, CertItemResp{
			ID:             cert.ID.String(),
			CourseID:       cert.CourseID.String(),
			CourseTitle:    cert.Course.TitleID, // default to ID
			CourseSlug:     cert.Course.Slug,
			Score:          cert.Score,
			TotalQuestions: cert.TotalQuestions,
			PassedAt:       cert.PassedAt,
		})
	}

	return items, nil
}

func (s *UserService) GetQuizHistory(userID uuid.UUID) ([]QuizHistoryItemResp, error) {
	history, err := s.progressRepo.GetUserQuizHistory(userID)
	if err != nil {
		return nil, err
	}

	var items []QuizHistoryItemResp
	for _, h := range history {
		items = append(items, QuizHistoryItemResp{
			Type:           h.Type,
			CourseTitle:    h.CourseTitle,
			LessonTitle:    h.LessonTitle,
			Score:          h.Score,
			TotalQuestions: h.TotalQuestions,
			ResultDetails:  h.ResultDetails,
			PassedAt:       h.PassedAt,
		})
	}

	if items == nil {
		items = make([]QuizHistoryItemResp, 0)
	}
	return items, nil
}
