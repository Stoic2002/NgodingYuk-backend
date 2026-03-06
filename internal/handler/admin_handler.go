package handler

import (
	"encoding/json"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"

	"github.com/arulkarim/ngodingyuk-server/internal/domain"
	"github.com/arulkarim/ngodingyuk-server/internal/repository"
)

type AdminHandler struct {
	challengeRepo *repository.ChallengeRepository
	courseRepo    *repository.CourseRepository
}

func NewAdminHandler(challengeRepo *repository.ChallengeRepository, courseRepo *repository.CourseRepository) *AdminHandler {
	return &AdminHandler{
		challengeRepo: challengeRepo,
		courseRepo:    courseRepo,
	}
}

// === Challenge Admin ===

// CreateChallenge handles POST /api/admin/challenges
func (h *AdminHandler) CreateChallenge(c *fiber.Ctx) error {
	var body struct {
		Slug           string          `json:"slug"`
		Language       string          `json:"language"`
		Difficulty     string          `json:"difficulty"`
		TitleID        string          `json:"title_id"`
		TitleEN        *string         `json:"title_en"`
		StoryID        string          `json:"story_id"`
		StoryEN        *string         `json:"story_en"`
		TaskID         string          `json:"task_id"`
		TaskEN         *string         `json:"task_en"`
		HintID         *string         `json:"hint_id"`
		HintEN         *string         `json:"hint_en"`
		SchemaInfo     json.RawMessage `json:"schema_info"`
		ExpectedOutput json.RawMessage `json:"expected_output"`
		StarterCode    *string         `json:"starter_code"`
		SolutionCode   *string         `json:"solution_code"`
		TestCases      json.RawMessage `json:"test_cases"`
		XPReward       int64           `json:"xp_reward"`
		OrderIndex     int             `json:"order_index"`
	}
	if err := c.BodyParser(&body); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid request body"})
	}

	challenge := &domain.Challenge{
		Slug:           body.Slug,
		Language:       body.Language,
		Difficulty:     body.Difficulty,
		TitleID:        body.TitleID,
		TitleEN:        body.TitleEN,
		StoryID:        body.StoryID,
		StoryEN:        body.StoryEN,
		TaskID:         body.TaskID,
		TaskEN:         body.TaskEN,
		HintID:         body.HintID,
		HintEN:         body.HintEN,
		SchemaInfo:     body.SchemaInfo,
		ExpectedOutput: body.ExpectedOutput,
		StarterCode:    body.StarterCode,
		SolutionCode:   body.SolutionCode,
		TestCases:      body.TestCases,
		XPReward:       body.XPReward,
		OrderIndex:     body.OrderIndex,
	}

	if err := h.challengeRepo.Create(challenge); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.Status(fiber.StatusCreated).JSON(fiber.Map{"id": challenge.ID.String()})
}

// UpdateChallenge handles PUT /api/admin/challenges/:id
func (h *AdminHandler) UpdateChallenge(c *fiber.Ctx) error {
	id, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid id"})
	}

	challenge, err := h.challengeRepo.FindByID(id)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "challenge not found"})
	}

	if err := c.BodyParser(challenge); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid request body"})
	}
	challenge.ID = id // ensure ID doesn't change

	if err := h.challengeRepo.Update(challenge); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(fiber.Map{"message": "updated"})
}

// DeleteChallenge handles DELETE /api/admin/challenges/:id
func (h *AdminHandler) DeleteChallenge(c *fiber.Ctx) error {
	id, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid id"})
	}

	if err := h.challengeRepo.Delete(id); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(fiber.Map{"message": "deleted"})
}

// === Course Admin ===

// CreateCourse handles POST /api/admin/courses
func (h *AdminHandler) CreateCourse(c *fiber.Ctx) error {
	var body struct {
		Slug          string  `json:"slug"`
		Language      string  `json:"language"`
		Level         string  `json:"level"`
		TitleID       string  `json:"title_id"`
		TitleEN       *string `json:"title_en"`
		DescriptionID *string `json:"description_id"`
		DescriptionEN *string `json:"description_en"`
		ThumbnailURL  *string `json:"thumbnail_url"`
		OrderIndex    int     `json:"order_index"`
	}
	if err := c.BodyParser(&body); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid request body"})
	}

	course := &domain.Course{
		Slug:          body.Slug,
		Language:      body.Language,
		Level:         body.Level,
		TitleID:       body.TitleID,
		TitleEN:       body.TitleEN,
		DescriptionID: body.DescriptionID,
		DescriptionEN: body.DescriptionEN,
		ThumbnailURL:  body.ThumbnailURL,
		OrderIndex:    body.OrderIndex,
	}

	if err := h.courseRepo.CreateCourse(course); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.Status(fiber.StatusCreated).JSON(fiber.Map{"id": course.ID.String()})
}

// CreateLesson handles POST /api/admin/courses/:id/lessons
func (h *AdminHandler) CreateLesson(c *fiber.Ctx) error {
	courseID, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid course id"})
	}

	var body struct {
		TitleID           string  `json:"title_id"`
		TitleEN           *string `json:"title_en"`
		ContentMarkdownID string  `json:"content_markdown_id"`
		ContentMarkdownEN *string `json:"content_markdown_en"`
		OrderIndex        int     `json:"order_index"`
		XPReward          int64   `json:"xp_reward"`
	}
	if err := c.BodyParser(&body); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid request body"})
	}

	lesson := &domain.Lesson{
		CourseID:          courseID,
		TitleID:           body.TitleID,
		TitleEN:           body.TitleEN,
		ContentMarkdownID: body.ContentMarkdownID,
		ContentMarkdownEN: body.ContentMarkdownEN,
		OrderIndex:        body.OrderIndex,
		XPReward:          body.XPReward,
	}

	if err := h.courseRepo.CreateLesson(lesson); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.Status(fiber.StatusCreated).JSON(fiber.Map{"id": lesson.ID.String()})
}

// UpdateLesson handles PUT /api/admin/lessons/:id
func (h *AdminHandler) UpdateLesson(c *fiber.Ctx) error {
	id, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid id"})
	}

	lesson, err := h.courseRepo.FindLessonByID(id)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "lesson not found"})
	}

	if err := c.BodyParser(lesson); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid request body"})
	}
	lesson.ID = id

	if err := h.courseRepo.UpdateLesson(lesson); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(fiber.Map{"message": "updated"})
}

// DeleteLesson handles DELETE /api/admin/lessons/:id
func (h *AdminHandler) DeleteLesson(c *fiber.Ctx) error {
	id, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid id"})
	}

	if err := h.courseRepo.DeleteLesson(id); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(fiber.Map{"message": "deleted"})
}

// CreateQuiz handles POST /api/admin/lessons/:id/quizzes
func (h *AdminHandler) CreateQuiz(c *fiber.Ctx) error {
	lessonID, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid lesson id"})
	}

	var body struct {
		QuestionID    string          `json:"question_id"`
		QuestionEN    *string         `json:"question_en"`
		OptionsID     json.RawMessage `json:"options_id"`
		OptionsEN     json.RawMessage `json:"options_en"`
		CorrectIndex  int             `json:"correct_index"`
		ExplanationID *string         `json:"explanation_id"`
		ExplanationEN *string         `json:"explanation_en"`
		OrderIndex    int             `json:"order_index"`
		XPReward      int64           `json:"xp_reward"`
	}
	if err := c.BodyParser(&body); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid request body"})
	}

	quiz := &domain.LessonQuiz{
		LessonID:      lessonID,
		QuestionID:    body.QuestionID,
		QuestionEN:    body.QuestionEN,
		OptionsID:     body.OptionsID,
		OptionsEN:     body.OptionsEN,
		CorrectIndex:  body.CorrectIndex,
		ExplanationID: body.ExplanationID,
		ExplanationEN: body.ExplanationEN,
		OrderIndex:    body.OrderIndex,
		XPReward:      body.XPReward,
	}

	if err := h.courseRepo.CreateQuiz(quiz); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.Status(fiber.StatusCreated).JSON(fiber.Map{"id": quiz.ID.String()})
}
