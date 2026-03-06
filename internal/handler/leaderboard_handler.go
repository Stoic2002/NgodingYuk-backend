package handler

import (
	"github.com/gofiber/fiber/v2"

	"github.com/arulkarim/ngodingyuk-server/internal/service"
)

type LeaderboardHandler struct {
	svc *service.LeaderboardService
}

func NewLeaderboardHandler(svc *service.LeaderboardService) *LeaderboardHandler {
	return &LeaderboardHandler{svc: svc}
}

// GetWeekly handles GET /api/leaderboard/weekly
func (h *LeaderboardHandler) GetWeekly(c *fiber.Ctx) error {
	entries, err := h.svc.GetWeekly()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(fiber.Map{"data": entries})
}

// GetAllTime handles GET /api/leaderboard/all-time
func (h *LeaderboardHandler) GetAllTime(c *fiber.Ctx) error {
	entries, err := h.svc.GetAllTime()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(fiber.Map{"data": entries})
}
