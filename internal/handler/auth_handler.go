package handler

import (
	"time"

	"github.com/gofiber/fiber/v2"

	"github.com/arulkarim/ngodingyuk-server/internal/service"
)

type AuthHandler struct {
	svc *service.AuthService
}

func NewAuthHandler(svc *service.AuthService) *AuthHandler {
	return &AuthHandler{svc: svc}
}

// setAuthCookies helper to set access & refresh tokens in cookies
func setAuthCookies(c *fiber.Ctx, accessToken, refreshToken string) {
	c.Cookie(&fiber.Cookie{
		Name:     "access_token",
		Value:    accessToken,
		Expires:  time.Now().Add(15 * time.Minute), // Match JWT expiry
		HTTPOnly: true,
		Secure:   true,   // Require HTTPS for cross-origin
		SameSite: "None", // Required for cross-origin cookies
		Path:     "/",
	})
	c.Cookie(&fiber.Cookie{
		Name:     "refresh_token",
		Value:    refreshToken,
		Expires:  time.Now().Add(7 * 24 * time.Hour), // 7 days
		HTTPOnly: true,
		Secure:   true,   // Require HTTPS for cross-origin
		SameSite: "None", // Required for cross-origin cookies
		Path:     "/",
	})
}

// clearAuthCookies helper to clear auth cookies on logout
func clearAuthCookies(c *fiber.Ctx) {
	c.Cookie(&fiber.Cookie{
		Name:     "access_token",
		Value:    "",
		Expires:  time.Now().Add(-1 * time.Hour), // Expire immediately
		HTTPOnly: true,
		Secure:   true,
		SameSite: "None",
		Path:     "/",
	})
	c.Cookie(&fiber.Cookie{
		Name:     "refresh_token",
		Value:    "",
		Expires:  time.Now().Add(-1 * time.Hour), // Expire immediately
		HTTPOnly: true,
		Secure:   true,
		SameSite: "None",
		Path:     "/",
	})
}

// Register handles POST /api/auth/register
func (h *AuthHandler) Register(c *fiber.Ctx) error {
	var req service.RegisterRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid request body"})
	}

	resp, err := h.svc.Register(req)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	// Set cookies rather than strictly relying on JSON response
	setAuthCookies(c, resp.AccessToken, resp.RefreshToken)

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"user":    resp.User,
		"message": "registration successful",
	})
}

// Login handles POST /api/auth/login
func (h *AuthHandler) Login(c *fiber.Ctx) error {
	var req service.LoginRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid request body"})
	}

	resp, err := h.svc.Login(req)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": err.Error()})
	}

	setAuthCookies(c, resp.AccessToken, resp.RefreshToken)

	return c.JSON(fiber.Map{
		"user":    resp.User,
		"message": "login successful",
	})
}

// Refresh handles POST /api/auth/refresh
func (h *AuthHandler) Refresh(c *fiber.Ctx) error {
	// First check cookie, then fallback to JSON body for backwards compatibility
	refreshToken := c.Cookies("refresh_token")
	if refreshToken == "" {
		var req service.RefreshRequest
		if err := c.BodyParser(&req); err == nil {
			refreshToken = req.RefreshToken
		}
	}

	if refreshToken == "" {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "no refresh token provided"})
	}

	resp, err := h.svc.RefreshToken(service.RefreshRequest{RefreshToken: refreshToken})
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": err.Error()})
	}

	setAuthCookies(c, resp.AccessToken, resp.RefreshToken)

	return c.JSON(fiber.Map{
		"user":    resp.User,
		"message": "token refreshed successfully",
	})
}

// Logout handles POST /api/auth/logout
func (h *AuthHandler) Logout(c *fiber.Ctx) error {
	clearAuthCookies(c)
	return c.JSON(fiber.Map{"message": "logged out successfully"})
}

// GetProfile handles GET /api/auth/me
func (h *AuthHandler) GetProfile(c *fiber.Ctx) error {
	userID, ok := getUserID(c)
	if !ok {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "unauthorized"})
	}

	profile, err := h.svc.GetProfile(userID)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(profile)
}
