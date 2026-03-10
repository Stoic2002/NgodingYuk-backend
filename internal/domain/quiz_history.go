package domain

import (
	"encoding/json"
	"time"
)

// QuizHistoryItem represents a combined view of quiz and exam results for a user.
type QuizHistoryItem struct {
	Type           string          `json:"type"` // "lesson_quiz" or "course_exam"
	CourseTitle    string          `json:"course_title"`
	LessonTitle    string          `json:"lesson_title,omitempty"`
	Score          int             `json:"score"`
	TotalQuestions int             `json:"total_questions"`
	ResultDetails  json.RawMessage `json:"result_details,omitempty"`
	PassedAt       time.Time       `json:"passed_at"`
}
