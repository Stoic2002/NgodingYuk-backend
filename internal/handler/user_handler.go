package handler

import (
	"github.com/gofiber/fiber/v2"

	"github.com/arulkarim/ngodingyuk-server/internal/service"
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
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "unauthorized"})
	}

	user, err := h.userService.GetProfile(userID)
	if err != nil {
		if err.Error() == "user not found" {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": err.Error()})
		}
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(user)
}

// GetXPHistory handles GET /api/users/me/xp-history
func (h *UserHandler) GetXPHistory(c *fiber.Ctx) error {
	userID, ok := getUserID(c)
	if !ok {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "unauthorized"})
	}

	limit := c.QueryInt("limit", 20)
	offset := c.QueryInt("offset", 0)

	items, err := h.userService.GetXPHistory(userID, limit, offset)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	// Make sure we always return an array, even if empty
	if items == nil {
		items = make([]service.XPItemResp, 0)
	}

	return c.JSON(fiber.Map{"data": items})
}

// GetChallengeStats handles GET /api/users/me/challenge-stats
func (h *UserHandler) GetChallengeStats(c *fiber.Ctx) error {
	userID, ok := getUserID(c)
	if !ok {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "unauthorized"})
	}

	stats, err := h.userService.GetChallengeStats(userID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(stats)
}

// UpdateMe handles PATCH /api/users/me
func (h *UserHandler) UpdateMe(c *fiber.Ctx) error {
	userID, ok := getUserID(c)
	if !ok {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "unauthorized"})
	}

	var req service.UpdateUserReq
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid request body"})
	}

	user, err := h.userService.UpdateProfile(userID, req)
	if err != nil {
		if err.Error() == "username already taken" || err.Error() == "locale must be 'id' or 'en'" || err.Error() == "no fields to update" {
			// Handle user errors with Bad Request or Conflict depending on case,
			// for simplicity let's handle Conflict for username specifically.
			if err.Error() == "username already taken" {
				return c.Status(fiber.StatusConflict).JSON(fiber.Map{"error": err.Error()})
			}
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
		}
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "failed to update"})
	}

	return c.JSON(user)
}

// GetCertificates handles GET /api/users/me/certificates
func (h *UserHandler) GetCertificates(c *fiber.Ctx) error {
	userID, ok := getUserID(c)
	if !ok {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "unauthorized"})
	}

	items, err := h.userService.GetCertificates(userID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	// Make sure we always return an array, even if empty
	if items == nil {
		items = make([]service.CertItemResp, 0)
	}

	return c.JSON(fiber.Map{"data": items})
}
