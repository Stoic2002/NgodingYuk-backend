package middleware

import (
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"

	jwtpkg "github.com/arulkarim/ngodingyuk-server/pkg/jwt"
)

// AuthMiddleware validates the JWT token from the Authorization header
// and sets the userID (as uuid.UUID) in Fiber's context locals.
func AuthMiddleware() fiber.Handler {
	return func(c *fiber.Ctx) error {
		authHeader := c.Get("Authorization")
		if authHeader == "" {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": "missing authorization header",
			})
		}

		parts := strings.SplitN(authHeader, " ", 2)
		if len(parts) != 2 || strings.ToLower(parts[0]) != "bearer" {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": "invalid authorization format",
			})
		}

		claims, err := jwtpkg.ValidateToken(parts[1])
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
		authHeader := c.Get("Authorization")
		if authHeader == "" {
			return c.Next()
		}

		parts := strings.SplitN(authHeader, " ", 2)
		if len(parts) != 2 || strings.ToLower(parts[0]) != "bearer" {
			return c.Next()
		}

		claims, err := jwtpkg.ValidateToken(parts[1])
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
