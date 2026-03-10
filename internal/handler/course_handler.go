package handler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"

	"github.com/arulkarim/ngodingyuk-server/internal/service"
	"github.com/arulkarim/ngodingyuk-server/pkg/response"
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
	search := c.Query("search")
	locale := resolveLocale(c)
	limit := c.QueryInt("limit", 10)
	offset := c.QueryInt("offset", 0)

	courses, total, err := h.svc.ListCourses(language, level, search, locale, limit, offset)
	if err != nil {
		return response.Error(c, fiber.StatusInternalServerError, err.Error())
	}

	totalPages := 0
	if limit > 0 {
		totalPages = int((total + int64(limit) - 1) / int64(limit))
	}

	return response.Success(c, fiber.StatusOK, "success", courses, response.WithPagination(totalPages, int(total)))
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
		return response.Error(c, fiber.StatusNotFound, err.Error())
	}
	return response.Success(c, fiber.StatusOK, "success", detail)
}

// EnrollCourse handles POST /api/courses/:slug/enroll
func (h *CourseHandler) EnrollCourse(c *fiber.Ctx) error {
	slug := c.Params("slug")
	userID, ok := getUserID(c)
	if !ok {
		return response.Error(c, fiber.StatusUnauthorized, "unauthorized")
	}

	if err := h.svc.EnrollCourse(userID, slug); err != nil {
		return response.Error(c, fiber.StatusBadRequest, err.Error())
	}
	return response.Success(c, fiber.StatusOK, "enrolled successfully", nil)
}

// GetLessonDetail handles GET /api/courses/:slug/lessons/:lesson_id
func (h *CourseHandler) GetLessonDetail(c *fiber.Ctx) error {
	lessonIDStr := c.Params("lesson_id")
	locale := resolveLocale(c)

	lessonID, err := uuid.Parse(lessonIDStr)
	if err != nil {
		return response.Error(c, fiber.StatusBadRequest, "invalid lesson_id")
	}

	// Check lesson locking if user is authenticated
	if userID, ok := getUserID(c); ok {
		accessible, _ := h.svc.IsLessonAccessible(userID, lessonID)
		if !accessible {
			return response.Error(c, fiber.StatusForbidden, "lesson is locked")
		}
	}

	detail, err := h.svc.GetLessonDetail(lessonID, locale)
	if err != nil {
		return response.Error(c, fiber.StatusNotFound, err.Error())
	}
	return response.Success(c, fiber.StatusOK, "success", detail)
}

// CompleteLesson handles POST /api/courses/:slug/lessons/:lesson_id/complete
func (h *CourseHandler) CompleteLesson(c *fiber.Ctx) error {
	lessonIDStr := c.Params("lesson_id")
	userID, ok := getUserID(c)
	if !ok {
		return response.Error(c, fiber.StatusUnauthorized, "unauthorized")
	}

	lessonID, err := uuid.Parse(lessonIDStr)
	if err != nil {
		return response.Error(c, fiber.StatusBadRequest, "invalid lesson_id")
	}

	if err := h.svc.CompleteLesson(userID, lessonID); err != nil {
		return response.Error(c, fiber.StatusInternalServerError, err.Error())
	}

	return response.Success(c, fiber.StatusOK, "lesson completed", nil)
}

// SubmitQuiz handles POST /api/courses/:slug/lessons/:lesson_id/quiz
func (h *CourseHandler) SubmitQuiz(c *fiber.Ctx) error {
	lessonIDStr := c.Params("lesson_id")
	userID, ok := getUserID(c)
	if !ok {
		return response.Error(c, fiber.StatusUnauthorized, "unauthorized")
	}

	lessonID, err := uuid.Parse(lessonIDStr)
	if err != nil {
		return response.Error(c, fiber.StatusBadRequest, "invalid lesson_id")
	}

	var req service.QuizSubmitRequest
	if err := c.BodyParser(&req); err != nil {
		return response.Error(c, fiber.StatusBadRequest, "invalid request body")
	}

	result, err := h.svc.SubmitQuiz(userID, lessonID, req)
	if err != nil {
		return response.Error(c, fiber.StatusBadRequest, err.Error())
	}
	return response.Success(c, fiber.StatusOK, "success", result)
}

// GetExam handles GET /api/courses/:slug/exam
func (h *CourseHandler) GetExam(c *fiber.Ctx) error {
	slug := c.Params("slug")
	locale := resolveLocale(c)
	userID, ok := getUserID(c)
	if !ok {
		return response.Error(c, fiber.StatusUnauthorized, "unauthorized")
	}

	exam, err := h.svc.GetExam(userID, slug, locale)
	if err != nil {
		return response.Error(c, fiber.StatusBadRequest, err.Error())
	}
	return response.Success(c, fiber.StatusOK, "success", exam)
}

// SubmitExam handles POST /api/courses/:slug/exam/submit
func (h *CourseHandler) SubmitExam(c *fiber.Ctx) error {
	slug := c.Params("slug")
	userID, ok := getUserID(c)
	if !ok {
		return response.Error(c, fiber.StatusUnauthorized, "unauthorized")
	}

	var req service.ExamSubmitRequest
	if err := c.BodyParser(&req); err != nil {
		return response.Error(c, fiber.StatusBadRequest, "invalid request body")
	}

	result, err := h.svc.SubmitExam(userID, slug, req)
	if err != nil {
		return response.Error(c, fiber.StatusBadRequest, err.Error())
	}
	return response.Success(c, fiber.StatusOK, "success", result)
}

// GetMyProgress handles GET /api/courses/:slug/my-progress
func (h *CourseHandler) GetMyProgress(c *fiber.Ctx) error {
	slug := c.Params("slug")
	userID, ok := getUserID(c)
	if !ok {
		return response.Error(c, fiber.StatusUnauthorized, "unauthorized")
	}

	progress, err := h.svc.GetMyProgress(slug, userID)
	if err != nil {
		return response.Error(c, fiber.StatusInternalServerError, err.Error())
	}
	return response.Success(c, fiber.StatusOK, "success", progress)
}
