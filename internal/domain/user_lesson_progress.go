package domain

import (
	"time"

	"github.com/google/uuid"
)

type UserLessonProgress struct {
	ID            uuid.UUID  `gorm:"type:uuid;primaryKey;default:gen_random_uuid()" json:"id"`
	UserID        uuid.UUID  `gorm:"type:uuid;not null" json:"user_id"`
	LessonID      uuid.UUID  `gorm:"type:uuid;not null" json:"lesson_id"`
	QuizCompleted bool       `gorm:"default:false" json:"quiz_completed"`
	QuizScore     int        `gorm:"default:0" json:"quiz_score"`
	CompletedAt   *time.Time `json:"completed_at"`
	CreatedAt     time.Time  `json:"created_at"`
	UpdatedAt     time.Time  `json:"updated_at"`

	// Relations
	User   User   `gorm:"foreignKey:UserID" json:"-"`
	Lesson Lesson `gorm:"foreignKey:LessonID" json:"-"`
}

// Unique constraint on (user_id, lesson_id)
func (UserLessonProgress) TableName() string {
	return "user_lesson_progress"
}
