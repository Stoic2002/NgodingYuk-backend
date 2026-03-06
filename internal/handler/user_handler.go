package handler

import (
	"time"

	"github.com/gofiber/fiber/v2"

	"github.com/arulkarim/ngodingyuk-server/internal/repository"
	"github.com/arulkarim/ngodingyuk-server/internal/service"
)

type UserHandler struct {
	userRepo     *repository.UserRepository
	progressRepo *repository.ProgressRepository
	courseRepo   *repository.CourseRepository
}

func NewUserHandler(userRepo *repository.UserRepository, progressRepo *repository.ProgressRepository, courseRepo *repository.CourseRepository) *UserHandler {
	return &UserHandler{
		userRepo:     userRepo,
		progressRepo: progressRepo,
		courseRepo:   courseRepo,
	}
}

// GetMe handles GET /api/users/me
func (h *UserHandler) GetMe(c *fiber.Ctx) error {
	userID, ok := getUserID(c)
	if !ok {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "unauthorized"})
	}

	user, err := h.userRepo.FindByID(userID)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "user not found"})
	}

	return c.JSON(fiber.Map{
		"id":               user.ID.String(),
		"username":         user.Username,
		"email":            user.Email,
		"xp":               user.XP,
		"level":            user.Level,
		"streak_count":     user.StreakCount,
		"locale":           user.Locale,
		"last_active_date": user.LastActiveDate,
		"created_at":       user.CreatedAt,
	})
}

// GetXPHistory handles GET /api/users/me/xp-history
func (h *UserHandler) GetXPHistory(c *fiber.Ctx) error {
	userID, ok := getUserID(c)
	if !ok {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "unauthorized"})
	}

	limit := c.QueryInt("limit", 20)
	offset := c.QueryInt("offset", 0)

	history, err := h.progressRepo.GetUserXPHistory(userID, limit, offset)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	type xpItem struct {
		ID         string    `json:"id"`
		XPGained   int64     `json:"xp_gained"`
		SourceType string    `json:"source_type"`
		SourceID   string    `json:"source_id,omitempty"`
		CreatedAt  time.Time `json:"created_at"`
	}

	var items []xpItem
	for _, h := range history {
		item := xpItem{
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

	return c.JSON(fiber.Map{"data": items})
}

// GetChallengeStats handles GET /api/users/me/challenge-stats
func (h *UserHandler) GetChallengeStats(c *fiber.Ctx) error {
	userID, ok := getUserID(c)
	if !ok {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "unauthorized"})
	}

	stats, err := h.progressRepo.GetUserChallengeStats(userID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	// Organize by difficulty
	result := map[string]map[string]int64{
		"easy":   {"solved": 0, "attempted": 0},
		"medium": {"solved": 0, "attempted": 0},
		"hard":   {"solved": 0, "attempted": 0},
	}
	for _, s := range stats {
		if _, ok := result[s.Difficulty]; ok {
			result[s.Difficulty][s.Status] = s.Count
		}
	}

	return c.JSON(result)
}

// UpdateMe handles PATCH /api/users/me
func (h *UserHandler) UpdateMe(c *fiber.Ctx) error {
	userID, ok := getUserID(c)
	if !ok {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "unauthorized"})
	}

	var body struct {
		Username *string `json:"username"`
		Locale   *string `json:"locale"`
	}
	if err := c.BodyParser(&body); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid request body"})
	}

	fields := make(map[string]interface{})

	if body.Username != nil && *body.Username != "" {
		// Check uniqueness
		existing, _ := h.userRepo.FindByUsername(*body.Username)
		if existing != nil && existing.ID != userID {
			return c.Status(fiber.StatusConflict).JSON(fiber.Map{"error": "username already taken"})
		}
		fields["username"] = *body.Username
	}

	if body.Locale != nil {
		if *body.Locale != "id" && *body.Locale != "en" {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "locale must be 'id' or 'en'"})
		}
		fields["locale"] = *body.Locale
	}

	if len(fields) == 0 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "no fields to update"})
	}

	if err := h.userRepo.UpdateFields(userID, fields); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "failed to update"})
	}

	// Return updated user
	user, _ := h.userRepo.FindByID(userID)
	return c.JSON(service.UserResponse{
		ID:          user.ID.String(),
		Username:    user.Username,
		Email:       user.Email,
		XP:          user.XP,
		Level:       user.Level,
		StreakCount: user.StreakCount,
		Locale:      user.Locale,
	})
}

// GetCertificates handles GET /api/users/me/certificates
func (h *UserHandler) GetCertificates(c *fiber.Ctx) error {
	userID, ok := getUserID(c)
	if !ok {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "unauthorized"})
	}

	certs, err := h.courseRepo.GetUserCertificates(userID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	type certItem struct {
		ID             string    `json:"id"`
		CourseID       string    `json:"course_id"`
		CourseTitle    string    `json:"course_title"`
		CourseSlug     string    `json:"course_slug"`
		Score          int       `json:"score"`
		TotalQuestions int       `json:"total_questions"`
		PassedAt       time.Time `json:"passed_at"`
	}

	var items []certItem
	for _, cert := range certs {
		items = append(items, certItem{
			ID:             cert.ID.String(),
			CourseID:       cert.CourseID.String(),
			CourseTitle:    cert.Course.TitleID, // default to ID, ideally resolved by locale but this is fine for now
			CourseSlug:     cert.Course.Slug,
			Score:          cert.Score,
			TotalQuestions: cert.TotalQuestions,
			PassedAt:       cert.PassedAt,
		})
	}

	return c.JSON(fiber.Map{"data": items})
}
