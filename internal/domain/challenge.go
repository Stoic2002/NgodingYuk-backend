package domain

import (
	"encoding/json"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Challenge struct {
	ID         uuid.UUID `gorm:"type:uuid;primaryKey;default:gen_random_uuid()" json:"id"`
	Slug       string    `gorm:"uniqueIndex;size:255;not null" json:"slug"`
	Language   string    `gorm:"size:20;not null" json:"language"`   // 'sql' | 'golang'
	Difficulty string    `gorm:"size:10;not null" json:"difficulty"` // 'easy' | 'medium' | 'hard'

	// Bilingual content fields (suffix _id = Indonesia, _en = English)
	// _en fields can be NULL → fallback to _id in backend
	TitleID string  `gorm:"size:255;not null" json:"title_id"`
	TitleEN *string `gorm:"size:255" json:"title_en"`
	StoryID string  `gorm:"type:text;not null" json:"story_id"`
	StoryEN *string `gorm:"type:text" json:"story_en"`
	TaskID  string  `gorm:"type:text;not null" json:"task_id"`
	TaskEN  *string `gorm:"type:text" json:"task_en"`
	HintID  *string `gorm:"type:text" json:"hint_id"`
	HintEN  *string `gorm:"type:text" json:"hint_en"`

	// Non-text fields (no bilingual needed)
	SchemaInfo     json.RawMessage `gorm:"type:jsonb" json:"schema_info"` // SQL table info; null for golang
	ExpectedOutput json.RawMessage `gorm:"type:jsonb;not null" json:"expected_output"`
	StarterCode    *string         `gorm:"type:text" json:"starter_code"`
	SolutionCode   *string         `gorm:"type:text" json:"-"` // NEVER sent to client
	TestCases      json.RawMessage `gorm:"type:jsonb;not null" json:"test_cases"`
	XPReward       int64           `gorm:"default:10" json:"xp_reward"`
	OrderIndex     int             `gorm:"default:0" json:"order_index"`
	CreatedAt      time.Time       `json:"created_at"`
	UpdatedAt      time.Time       `json:"updated_at"`
	DeletedAt      gorm.DeletedAt  `gorm:"index" json:"-"`
}
