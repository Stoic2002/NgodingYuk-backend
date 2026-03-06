package domain

import (
	"time"

	"github.com/google/uuid"
)

type UserCourseEnrollment struct {
	ID          uuid.UUID  `gorm:"type:uuid;primaryKey;default:gen_random_uuid()" json:"id"`
	UserID      uuid.UUID  `gorm:"type:uuid;not null" json:"user_id"`
	CourseID    uuid.UUID  `gorm:"type:uuid;not null" json:"course_id"`
	EnrolledAt  time.Time  `gorm:"autoCreateTime" json:"enrolled_at"`
	CompletedAt *time.Time `json:"completed_at"`

	// Relations
	User   User   `gorm:"foreignKey:UserID" json:"-"`
	Course Course `gorm:"foreignKey:CourseID" json:"-"`
}

func (UserCourseEnrollment) TableName() string {
	return "user_course_enrollments"
}
