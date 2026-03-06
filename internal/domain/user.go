package domain

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	ID             uuid.UUID      `gorm:"type:uuid;primaryKey;default:gen_random_uuid()" json:"id"`
	Username       string         `gorm:"uniqueIndex;size:50;not null" json:"username"`
	Email          string         `gorm:"uniqueIndex;size:255;not null" json:"email"`
	PasswordHash   string         `gorm:"not null" json:"-"`
	XP             int64          `gorm:"default:0" json:"xp"`
	Level          int64          `gorm:"default:1" json:"level"`
	StreakCount    int64          `gorm:"default:0" json:"streak_count"`
	Locale         string         `gorm:"size:5;default:'id'" json:"locale"`
	LastActiveDate *time.Time     `json:"last_active_date"`
	CreatedAt      time.Time      `json:"created_at"`
	UpdatedAt      time.Time      `json:"updated_at"`
	DeletedAt      gorm.DeletedAt `gorm:"index" json:"-"`
}
