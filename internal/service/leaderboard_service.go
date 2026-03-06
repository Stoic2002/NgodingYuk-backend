package service

import (
	"github.com/arulkarim/ngodingyuk-server/internal/domain"
	"github.com/arulkarim/ngodingyuk-server/internal/repository"
)

type LeaderboardService struct {
	progressRepo *repository.ProgressRepository
}

func NewLeaderboardService(progressRepo *repository.ProgressRepository) *LeaderboardService {
	return &LeaderboardService{progressRepo: progressRepo}
}

// === Response DTOs ===

type LeaderboardEntry struct {
	Rank     int    `json:"rank"`
	UserID   string `json:"user_id"`
	Username string `json:"username"`
	XP       int64  `json:"xp"`
	Level    int64  `json:"level"`
}

// GetWeekly returns top 50 users by XP gained in the last 7 days.
func (s *LeaderboardService) GetWeekly() ([]LeaderboardEntry, error) {
	entries, err := s.progressRepo.GetWeeklyLeaderboard(50)
	if err != nil {
		return nil, err
	}

	var result []LeaderboardEntry
	for i, e := range entries {
		result = append(result, LeaderboardEntry{
			Rank:     i + 1,
			UserID:   e.UserID.String(),
			Username: e.Username,
			XP:       e.TotalXP,
			Level:    e.Level,
		})
	}
	return result, nil
}

// GetAllTime returns top 50 users by total XP.
func (s *LeaderboardService) GetAllTime() ([]LeaderboardEntry, error) {
	users, err := s.progressRepo.GetAllTimeLeaderboard(50)
	if err != nil {
		return nil, err
	}

	var result []LeaderboardEntry
	for i, u := range users {
		result = append(result, leaderboardFromUser(i+1, u))
	}
	return result, nil
}

func leaderboardFromUser(rank int, u domain.User) LeaderboardEntry {
	return LeaderboardEntry{
		Rank:     rank,
		UserID:   u.ID.String(),
		Username: u.Username,
		XP:       u.XP,
		Level:    u.Level,
	}
}
