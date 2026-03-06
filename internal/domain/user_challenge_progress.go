package domain

import (
	"time"

	"github.com/google/uuid"
)

type UserChallengeProgress struct {
	ID          uuid.UUID  `gorm:"type:uuid;primaryKey;default:gen_random_uuid()" json:"id"`
	UserID      uuid.UUID  `gorm:"type:uuid;not null" json:"user_id"`
	ChallengeID uuid.UUID  `gorm:"type:uuid;not null" json:"challenge_id"`
	Status      string     `gorm:"size:20;default:'attempted'" json:"status"` // 'attempted' | 'solved'
	Attempts    int        `gorm:"default:0" json:"attempts"`
	LastCode    *string    `gorm:"type:text" json:"last_code"`
	SolvedAt    *time.Time `json:"solved_at"`
	CreatedAt   time.Time  `json:"created_at"`
	UpdatedAt   time.Time  `json:"updated_at"`

	// Relations
	User      User      `gorm:"foreignKey:UserID" json:"-"`
	Challenge Challenge `gorm:"foreignKey:ChallengeID" json:"-"`
}

// Unique constraint on (user_id, challenge_id) is handled via GORM migration hook or manual SQL
func (UserChallengeProgress) TableName() string {
	return "user_challenge_progress"
}
