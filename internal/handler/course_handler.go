package handler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"

	"github.com/arulkarim/ngodingyuk-server/internal/service"
)

type CourseHandler struct {
	svc *service.CourseService
}

func NewCourseHandler(svc *service.CourseService) *CourseHandler {
	return &CourseHandler{svc: svc}
}

// ListCourses handles GET /api/courses
func (h *CourseHandler) ListCourses(c *fiber.Ctx) error {
	language := c.Query("language")
	level := c.Query("level")
	locale := resolveLocale(c)

	courses, err := h.svc.ListCourses(language, level, locale)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(fiber.Map{"data": courses})
}

// GetCourse handles GET /api/courses/:slug
func (h *CourseHandler) GetCourse(c *fiber.Ctx) error {
	slug := c.Params("slug")
	locale := resolveLocale(c)

	// Optional user ID for enrollment/locking info
	userID := uuid.Nil
	if uid, ok := getUserID(c); ok {
		userID = uid
	}

	detail, err := h.svc.GetCourseBySlug(slug, locale, userID)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(detail)
}

// EnrollCourse handles POST /api/courses/:slug/enroll
func (h *CourseHandler) EnrollCourse(c *fiber.Ctx) error {
	slug := c.Params("slug")
	userID, ok := getUserID(c)
	if !ok {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "unauthorized"})
	}

	if err := h.svc.EnrollCourse(userID, slug); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(fiber.Map{"message": "enrolled successfully"})
}

// GetLessonDetail handles GET /api/courses/:slug/lessons/:lesson_id
func (h *CourseHandler) GetLessonDetail(c *fiber.Ctx) error {
	lessonIDStr := c.Params("lesson_id")
	locale := resolveLocale(c)

	lessonID, err := uuid.Parse(lessonIDStr)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid lesson_id"})
	}

	// Check lesson locking if user is authenticated
	if userID, ok := getUserID(c); ok {
		accessible, _ := h.svc.IsLessonAccessible(userID, lessonID)
		if !accessible {
			return c.Status(fiber.StatusForbidden).JSON(fiber.Map{"error": "lesson is locked"})
		}
	}

	detail, err := h.svc.GetLessonDetail(lessonID, locale)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(detail)
}

// CompleteLesson handles POST /api/courses/:slug/lessons/:lesson_id/complete
func (h *CourseHandler) CompleteLesson(c *fiber.Ctx) error {
	lessonIDStr := c.Params("lesson_id")
	userID, ok := getUserID(c)
	if !ok {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "unauthorized"})
	}

	lessonID, err := uuid.Parse(lessonIDStr)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid lesson_id"})
	}

	if err := h.svc.CompleteLesson(userID, lessonID); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(fiber.Map{"message": "lesson completed"})
}

// SubmitQuiz handles POST /api/courses/:slug/lessons/:lesson_id/quiz
func (h *CourseHandler) SubmitQuiz(c *fiber.Ctx) error {
	lessonIDStr := c.Params("lesson_id")
	userID, ok := getUserID(c)
	if !ok {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "unauthorized"})
	}

	lessonID, err := uuid.Parse(lessonIDStr)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid lesson_id"})
	}

	var req service.QuizSubmitRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid request body"})
	}

	result, err := h.svc.SubmitQuiz(userID, lessonID, req)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(result)
}

// GetExam handles GET /api/courses/:slug/exam
func (h *CourseHandler) GetExam(c *fiber.Ctx) error {
	slug := c.Params("slug")
	locale := resolveLocale(c)
	userID, ok := getUserID(c)
	if !ok {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "unauthorized"})
	}

	exam, err := h.svc.GetExam(userID, slug, locale)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(exam)
}

// SubmitExam handles POST /api/courses/:slug/exam/submit
func (h *CourseHandler) SubmitExam(c *fiber.Ctx) error {
	slug := c.Params("slug")
	userID, ok := getUserID(c)
	if !ok {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "unauthorized"})
	}

	var req service.ExamSubmitRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid request body"})
	}

	result, err := h.svc.SubmitExam(userID, slug, req)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(result)
}

// GetMyProgress handles GET /api/courses/:slug/my-progress
func (h *CourseHandler) GetMyProgress(c *fiber.Ctx) error {
	slug := c.Params("slug")
	userID, ok := getUserID(c)
	if !ok {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "unauthorized"})
	}

	progress, err := h.svc.GetMyProgress(slug, userID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(progress)
}
