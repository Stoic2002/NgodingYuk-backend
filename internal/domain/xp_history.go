package domain

import (
	"time"

	"github.com/google/uuid"
)

type UserXPHistory struct {
	ID         uuid.UUID  `gorm:"type:uuid;primaryKey;default:gen_random_uuid()" json:"id"`
	UserID     uuid.UUID  `gorm:"type:uuid;not null;index" json:"user_id"`
	XPGained   int64      `gorm:"not null" json:"xp_gained"`
	SourceType string     `gorm:"size:30;not null" json:"source_type"` // 'challenge' | 'lesson_quiz'
	SourceID   *uuid.UUID `gorm:"type:uuid" json:"source_id"`
	CreatedAt  time.Time  `json:"created_at"`

	// Relations
	User User `gorm:"foreignKey:UserID" json:"-"`
}

func (UserXPHistory) TableName() string {
	return "user_xp_history"
}
