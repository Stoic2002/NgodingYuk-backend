package domain

import (
	"encoding/json"
	"time"

	"github.com/google/uuid"
)

type LessonQuiz struct {
	ID           uuid.UUID `gorm:"type:uuid;primaryKey;default:gen_random_uuid()" json:"id"`
	LessonID     uuid.UUID `gorm:"type:uuid;not null;index" json:"lesson_id"`
	CorrectIndex int       `gorm:"not null" json:"correct_index"`
	OrderIndex   int       `gorm:"default:0" json:"order_index"`
	XPReward     int64     `gorm:"default:5" json:"xp_reward"`

	// Bilingual content
	QuestionID    string          `gorm:"type:text;not null" json:"question_id"`
	QuestionEN    *string         `gorm:"type:text" json:"question_en"`
	OptionsID     json.RawMessage `gorm:"type:jsonb;not null" json:"options_id"` // ["opsi a","opsi b","opsi c","opsi d"]
	OptionsEN     json.RawMessage `gorm:"type:jsonb" json:"options_en"`          // English version (NULL = fallback)
	ExplanationID *string         `gorm:"type:text" json:"explanation_id"`
	ExplanationEN *string         `gorm:"type:text" json:"explanation_en"`

	// Relations
	Lesson Lesson `gorm:"foreignKey:LessonID" json:"-"`

	CreatedAt time.Time `json:"created_at"`
}
