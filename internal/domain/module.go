package domain

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Module struct {
	ID         uuid.UUID `gorm:"type:uuid;primaryKey;default:gen_random_uuid()" json:"id"`
	CourseID   uuid.UUID `gorm:"type:uuid;not null;index" json:"course_id"`
	TitleID    string    `gorm:"size:255;not null" json:"title_id"`
	TitleEN    *string   `gorm:"size:255" json:"title_en"`
	OrderIndex int       `gorm:"default:0" json:"order_index"`

	// Relations
	Lessons []Lesson `gorm:"foreignKey:ModuleID" json:"lessons,omitempty"`
	Course  *Course  `gorm:"foreignKey:CourseID" json:"-"`

	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
}
