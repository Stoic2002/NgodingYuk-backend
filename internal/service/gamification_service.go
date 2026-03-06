package service

import (
	"time"

	"github.com/arulkarim/ngodingyuk-server/internal/domain"
)

// Level thresholds per spec: 0 → 100 → 250 → 500 → 1000 → 2000 → ...
// After the defined thresholds, each level requires +1000 XP more.
var levelThresholds = []int64{0, 100, 250, 500, 1000, 2000, 3500, 5500, 8000, 11000, 15000}

// CalculateLevel determines the user's level based on total XP.
func CalculateLevel(xp int64) int64 {
	var level int64 = 1
	for i, threshold := range levelThresholds {
		if xp >= threshold {
			level = int64(i) + 1
		} else {
			break
		}
	}
	// Beyond defined thresholds: each additional level = last_threshold + N*1000
	if xp >= levelThresholds[len(levelThresholds)-1] {
		extra := (xp - levelThresholds[len(levelThresholds)-1]) / 1000
		level = int64(len(levelThresholds)) + extra
	}
	return level
}

// XPForLevel returns the total XP required to reach a given level.
func XPForLevel(level int64) int64 {
	if level <= 1 {
		return 0
	}
	idx := int(level) - 1
	if idx < len(levelThresholds) {
		return levelThresholds[idx]
	}
	// Beyond known thresholds
	excess := int64(idx - len(levelThresholds) + 1)
	return levelThresholds[len(levelThresholds)-1] + excess*1000
}

// UpdateStreak checks if the user's streak should be incremented or reset.
// Call this whenever the user performs an activity that counts as "active today".
func UpdateStreak(user *domain.User) {
	now := time.Now()
	today := now.Truncate(24 * time.Hour)

	if user.LastActiveDate == nil {
		user.StreakCount = 1
		user.LastActiveDate = &now
		return
	}

	lastActive := user.LastActiveDate.Truncate(24 * time.Hour)

	if today.Equal(lastActive) {
		// Same day — no streak change, just update timestamp
		return
	}

	if today.Equal(lastActive.Add(24 * time.Hour)) {
		// Next day — increment streak
		user.StreakCount++
	} else {
		// Gap > 1 day — reset streak
		user.StreakCount = 1
	}

	user.LastActiveDate = &now
}
