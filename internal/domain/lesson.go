package domain

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Lesson struct {
	ID         uuid.UUID  `gorm:"type:uuid;primaryKey;default:gen_random_uuid()" json:"id"`
	CourseID   uuid.UUID  `gorm:"type:uuid;not null;index" json:"course_id"`
	ModuleID   *uuid.UUID `gorm:"type:uuid;index" json:"module_id"`
	OrderIndex int        `gorm:"not null" json:"order_index"`
	XPReward   int64      `gorm:"default:5" json:"xp_reward"`

	// Bilingual content
	TitleID           string  `gorm:"size:255;not null" json:"title_id"`
	TitleEN           *string `gorm:"size:255" json:"title_en"`
	ContentMarkdownID string  `gorm:"type:text;not null" json:"content_markdown_id"`
	ContentMarkdownEN *string `gorm:"type:text" json:"content_markdown_en"`

	// Relations
	Course  Course       `gorm:"foreignKey:CourseID" json:"-"`
	Module  *Module      `gorm:"foreignKey:ModuleID" json:"-"`
	Quizzes []LessonQuiz `gorm:"foreignKey:LessonID" json:"quizzes,omitempty"`

	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
}
