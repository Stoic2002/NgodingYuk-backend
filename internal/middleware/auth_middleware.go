package middleware

import (
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"

	jwtpkg "github.com/arulkarim/ngodingyuk-server/pkg/jwt"
)

// extractToken helper to get token from cookie or header
func extractToken(c *fiber.Ctx) string {
	// 1. Try cookie first
	token := c.Cookies("access_token")
	if token != "" {
		return token
	}

	// 2. Fall back to Authorization header
	authHeader := c.Get("Authorization")
	if authHeader != "" {
		parts := strings.SplitN(authHeader, " ", 2)
		if len(parts) == 2 && strings.ToLower(parts[0]) == "bearer" {
			return parts[1]
		}
	}
	return ""
}

// AuthMiddleware validates the JWT token from Cookie or Authorization header
// and sets the userID (as uuid.UUID) in Fiber's context locals.
func AuthMiddleware() fiber.Handler {
	return func(c *fiber.Ctx) error {
		token := extractToken(c)
		if token == "" {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": "unauthorized, missing token",
			})
		}

		claims, err := jwtpkg.ValidateToken(token)
		if err != nil {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": "invalid or expired token",
			})
		}

		// Parse string user ID into UUID
		uid, err := uuid.Parse(claims.UserID)
		if err != nil {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": "invalid user ID in token",
			})
		}

		// Store parsed UUID in context for downstream handlers
		c.Locals("userID", uid)
		return c.Next()
	}
}

// OptionalAuthMiddleware attempts to parse the JWT but does not block the request
// if no token is provided. Useful for endpoints that behave differently for
// authenticated vs anonymous users.
func OptionalAuthMiddleware() fiber.Handler {
	return func(c *fiber.Ctx) error {
		token := extractToken(c)
		if token == "" {
			return c.Next()
		}

		claims, err := jwtpkg.ValidateToken(token)
		if err != nil {
			return c.Next()
		}

		uid, err := uuid.Parse(claims.UserID)
		if err != nil {
			return c.Next()
		}

		c.Locals("userID", uid)
		return c.Next()
	}
}
