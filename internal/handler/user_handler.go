package handler

import (
	"github.com/gofiber/fiber/v2"

	"github.com/arulkarim/ngodingyuk-server/internal/service"
	"github.com/arulkarim/ngodingyuk-server/pkg/response"
)

type UserHandler struct {
	userService *service.UserService
}

func NewUserHandler(userService *service.UserService) *UserHandler {
	return &UserHandler{
		userService: userService,
	}
}

// GetMe handles GET /api/users/me
func (h *UserHandler) GetMe(c *fiber.Ctx) error {
	userID, ok := getUserID(c)
	if !ok {
		return response.Error(c, fiber.StatusUnauthorized, "unauthorized")
	}

	user, err := h.userService.GetProfile(userID)
	if err != nil {
		if err.Error() == "user not found" {
			return response.Error(c, fiber.StatusNotFound, err.Error())
		}
		return response.Error(c, fiber.StatusInternalServerError, err.Error())
	}

	return response.Success(c, fiber.StatusOK, "success", user)
}

// GetXPHistory handles GET /api/users/me/xp-history
func (h *UserHandler) GetXPHistory(c *fiber.Ctx) error {
	userID, ok := getUserID(c)
	if !ok {
		return response.Error(c, fiber.StatusUnauthorized, "unauthorized")
	}

	limit := c.QueryInt("limit", 20)
	offset := c.QueryInt("offset", 0)

	items, err := h.userService.GetXPHistory(userID, limit, offset)
	if err != nil {
		return response.Error(c, fiber.StatusInternalServerError, err.Error())
	}

	// Make sure we always return an array, even if empty
	if items == nil {
		items = make([]service.XPItemResp, 0)
	}

	return response.Success(c, fiber.StatusOK, "success", items)
}

// GetChallengeStats handles GET /api/users/me/challenge-stats
func (h *UserHandler) GetChallengeStats(c *fiber.Ctx) error {
	userID, ok := getUserID(c)
	if !ok {
		return response.Error(c, fiber.StatusUnauthorized, "unauthorized")
	}

	stats, err := h.userService.GetChallengeStats(userID)
	if err != nil {
		return response.Error(c, fiber.StatusInternalServerError, err.Error())
	}

	return response.Success(c, fiber.StatusOK, "success", stats)
}

// UpdateMe handles PATCH /api/users/me
func (h *UserHandler) UpdateMe(c *fiber.Ctx) error {
	userID, ok := getUserID(c)
	if !ok {
		return response.Error(c, fiber.StatusUnauthorized, "unauthorized")
	}

	var req service.UpdateUserReq
	if err := c.BodyParser(&req); err != nil {
		return response.Error(c, fiber.StatusBadRequest, "invalid request body")
	}

	user, err := h.userService.UpdateProfile(userID, req)
	if err != nil {
		if err.Error() == "username already taken" || err.Error() == "locale must be 'id' or 'en'" || err.Error() == "no fields to update" {
			if err.Error() == "username already taken" {
				return response.Error(c, fiber.StatusConflict, err.Error())
			}
			return response.Error(c, fiber.StatusBadRequest, err.Error())
		}
		return response.Error(c, fiber.StatusInternalServerError, "failed to update")
	}

	return response.Success(c, fiber.StatusOK, "success", user)
}

// GetCertificates handles GET /api/users/me/certificates
func (h *UserHandler) GetCertificates(c *fiber.Ctx) error {
	userID, ok := getUserID(c)
	if !ok {
		return response.Error(c, fiber.StatusUnauthorized, "unauthorized")
	}

	items, err := h.userService.GetCertificates(userID)
	if err != nil {
		return response.Error(c, fiber.StatusInternalServerError, err.Error())
	}

	// Make sure we always return an array, even if empty
	if items == nil {
		items = make([]service.CertItemResp, 0)
	}

	return response.Success(c, fiber.StatusOK, "success", items)
}

// GetQuizHistory handles GET /api/users/me/quiz-history
func (h *UserHandler) GetQuizHistory(c *fiber.Ctx) error {
	userID, ok := getUserID(c)
	if !ok {
		return response.Error(c, fiber.StatusUnauthorized, "unauthorized")
	}

	items, err := h.userService.GetQuizHistory(userID)
	if err != nil {
		return response.Error(c, fiber.StatusInternalServerError, err.Error())
	}

	return response.Success(c, fiber.StatusOK, "success", items)
}
