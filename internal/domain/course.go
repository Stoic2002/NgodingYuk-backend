package domain

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Course struct {
	ID       uuid.UUID `gorm:"type:uuid;primaryKey;default:gen_random_uuid()" json:"id"`
	Slug     string    `gorm:"uniqueIndex;size:255;not null" json:"slug"`
	Language string    `gorm:"size:20;not null" json:"language"` // 'sql' | 'golang'
	Level    string    `gorm:"size:20;not null" json:"level"`    // 'beginner' | 'intermediate' | 'advanced'

	// Bilingual content
	TitleID       string  `gorm:"size:255;not null" json:"title_id"`
	TitleEN       *string `gorm:"size:255" json:"title_en"`
	DescriptionID *string `gorm:"type:text" json:"description_id"`
	DescriptionEN *string `gorm:"type:text" json:"description_en"`

	ThumbnailURL *string `gorm:"size:255" json:"thumbnail_url"`
	OrderIndex   int     `gorm:"default:0" json:"order_index"`

	// Relations
	Modules []Module `gorm:"foreignKey:CourseID" json:"modules,omitempty"`
	Lessons []Lesson `gorm:"foreignKey:CourseID" json:"lessons,omitempty"`

	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
}
