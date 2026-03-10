package handler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"

	"github.com/arulkarim/ngodingyuk-server/internal/service"
	"github.com/arulkarim/ngodingyuk-server/pkg/response"
)

type ChallengeHandler struct {
	svc *service.ChallengeService
}

func NewChallengeHandler(svc *service.ChallengeService) *ChallengeHandler {
	return &ChallengeHandler{svc: svc}
}

// resolveLocale extracts locale from query param, defaulting to "id".
func resolveLocale(c *fiber.Ctx) string {
	locale := c.Query("locale", "id")
	if locale != "en" && locale != "id" {
		locale = "id"
	}
	return locale
}

// getUserID extracts the authenticated user's UUID from Fiber Locals.
func getUserID(c *fiber.Ctx) (uuid.UUID, bool) {
	uid, ok := c.Locals("userID").(uuid.UUID)
	return uid, ok
}

// ListChallenges handles GET /api/challenges
func (h *ChallengeHandler) ListChallenges(c *fiber.Ctx) error {
	language := c.Query("language")
	difficulty := c.Query("difficulty")
	search := c.Query("search")
	locale := resolveLocale(c)
	limit := c.QueryInt("limit", 50)
	offset := c.QueryInt("offset", 0)

	userID, _ := getUserID(c)

	items, total, err := h.svc.List(language, difficulty, search, locale, limit, offset, userID)
	if err != nil {
		return response.Error(c, fiber.StatusInternalServerError, err.Error())
	}

	totalPages := 0
	if limit > 0 {
		totalPages = int((total + int64(limit) - 1) / int64(limit))
	}

	return response.Success(c, fiber.StatusOK, "success", items, response.WithPagination(totalPages, int(total)))
}

// GetChallenge handles GET /api/challenges/:slug
func (h *ChallengeHandler) GetChallenge(c *fiber.Ctx) error {
	slug := c.Params("slug")
	locale := resolveLocale(c)

	detail, err := h.svc.GetBySlug(slug, locale)
	if err != nil {
		return response.Error(c, fiber.StatusNotFound, err.Error())
	}
	return response.Success(c, fiber.StatusOK, "success", detail)
}

// RunChallenge handles POST /api/challenges/:slug/run
func (h *ChallengeHandler) RunChallenge(c *fiber.Ctx) error {
	slug := c.Params("slug")

	var body struct {
		Code string `json:"code"`
	}
	if err := c.BodyParser(&body); err != nil {
		return response.Error(c, fiber.StatusBadRequest, "invalid request body")
	}
	if body.Code == "" {
		return response.Error(c, fiber.StatusBadRequest, "code is required")
	}

	result, err := h.svc.Run(slug, body.Code)
	if err != nil {
		return response.Error(c, fiber.StatusBadRequest, err.Error())
	}
	return response.Success(c, fiber.StatusOK, "success", result)
}

// SubmitChallenge handles POST /api/challenges/:slug/submit
func (h *ChallengeHandler) SubmitChallenge(c *fiber.Ctx) error {
	slug := c.Params("slug")
	userID, ok := getUserID(c)
	if !ok {
		return response.Error(c, fiber.StatusUnauthorized, "unauthorized")
	}

	var body struct {
		Code string `json:"code"`
	}
	if err := c.BodyParser(&body); err != nil {
		return response.Error(c, fiber.StatusBadRequest, "invalid request body")
	}
	if body.Code == "" {
		return response.Error(c, fiber.StatusBadRequest, "code is required")
	}

	result, err := h.svc.Submit(slug, userID, body.Code)
	if err != nil {
		return response.Error(c, fiber.StatusInternalServerError, err.Error())
	}
	return response.Success(c, fiber.StatusOK, "success", result)
}

// GetMyProgress handles GET /api/challenges/:slug/my-progress
func (h *ChallengeHandler) GetMyProgress(c *fiber.Ctx) error {
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
