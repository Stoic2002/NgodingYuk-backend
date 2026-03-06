package domain

import (
	"time"

	"github.com/google/uuid"
)

type Certificate struct {
	ID             uuid.UUID `gorm:"type:uuid;primaryKey;default:gen_random_uuid()" json:"id"`
	UserID         uuid.UUID `gorm:"type:uuid;not null" json:"user_id"`
	CourseID       uuid.UUID `gorm:"type:uuid;not null" json:"course_id"`
	Score          int       `gorm:"not null" json:"score"`
	TotalQuestions int       `gorm:"not null" json:"total_questions"`
	PassedAt       time.Time `gorm:"autoCreateTime" json:"passed_at"`

	// Relations
	User   User   `gorm:"foreignKey:UserID" json:"-"`
	Course Course `gorm:"foreignKey:CourseID" json:"-"`
}

func (Certificate) TableName() string {
	return "certificates"
}
