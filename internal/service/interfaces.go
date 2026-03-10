package service

import (
	"github.com/arulkarim/ngodingyuk-server/internal/domain"
	"github.com/google/uuid"
)

// ChallengeRepository defines the interface for challenge data storage.
type ChallengeRepository interface {
	List(language, difficulty, search string, limit, offset int) ([]domain.Challenge, int64, error)
	FindBySlug(slug string) (*domain.Challenge, error)
	Create(challenge *domain.Challenge) error
	Update(challenge *domain.Challenge) error
	Delete(id uuid.UUID) error
}

// CourseRepository defines the interface for course data storage.
type CourseRepository interface {
	// Courses
	ListCourses(language, level, search string, limit, offset int) ([]domain.Course, int64, error)
	FindCourseBySlug(slug string, preloadLessons bool) (*domain.Course, error)
	FindCourseByID(id uuid.UUID) (*domain.Course, error)
	CreateCourse(course *domain.Course) error
	UpdateCourse(course *domain.Course) error
	DeleteCourse(id uuid.UUID) error

	// Lessons
	FindLessonByID(id uuid.UUID) (*domain.Lesson, error)
	ListLessonsByCourse(courseID uuid.UUID) ([]domain.Lesson, error)
	CreateLesson(lesson *domain.Lesson) error
	UpdateLesson(lesson *domain.Lesson) error
	DeleteLesson(id uuid.UUID) error
	CountLessonsInCourse(courseID uuid.UUID) (int64, error)

	// Lesson Quizzes
	FindQuizzesByLesson(lessonID uuid.UUID) ([]domain.LessonQuiz, error)
	FindQuizByID(id uuid.UUID) (*domain.LessonQuiz, error)
	CreateQuiz(quiz *domain.LessonQuiz) error

	// Exam
	FindAllQuizzesByCourse(courseID uuid.UUID) ([]domain.LessonQuiz, error)

	// Enrollment & Certificates
	EnrollUser(userID, courseID uuid.UUID) error
	IsEnrolled(userID, courseID uuid.UUID) (bool, error)
	GetEnrollment(userID, courseID uuid.UUID) (*domain.UserCourseEnrollment, error)
	CreateCertificate(cert *domain.Certificate) error
	GetCertificate(userID, courseID uuid.UUID) (*domain.Certificate, error)
	GetUserCertificates(userID uuid.UUID) ([]domain.Certificate, error)
}

// ProgressRepository defines the interface for user progress storage.
type ProgressRepository interface {
	// Challenge Progress
	FindChallengeProgress(userID, challengeID uuid.UUID) (*domain.UserChallengeProgress, error)
	GetCompletedChallengeIDs(userID uuid.UUID) ([]uuid.UUID, error)
	FindOrCreateChallengeProgress(userID, challengeID uuid.UUID) (*domain.UserChallengeProgress, error)
	UpdateChallengeProgress(progress *domain.UserChallengeProgress) error
	GetUserChallengeStats(userID uuid.UUID) ([]struct {
		Difficulty string
		Status     string
		Count      int64
	}, error)

	// Lesson Progress
	FindLessonProgress(userID, lessonID uuid.UUID) (*domain.UserLessonProgress, error)
	FindOrCreateLessonProgress(userID, lessonID uuid.UUID) (*domain.UserLessonProgress, error)
	UpdateLessonProgress(progress *domain.UserLessonProgress) error
	CountCompletedLessonsInCourse(userID, courseID uuid.UUID) (int64, error)
	GetUserLessonProgressByCourse(userID, courseID uuid.UUID) ([]domain.UserLessonProgress, error)

	// XP & Gamification
	CreateXPHistory(history *domain.UserXPHistory) error
	GetUserXPHistory(userID uuid.UUID, limit, offset int) ([]domain.UserXPHistory, error)
	GetUserQuizHistory(userID uuid.UUID) ([]domain.QuizHistoryItem, error)
	GetWeeklyLeaderboard(limit int) ([]struct {
		UserID   uuid.UUID
		Username string
		TotalXP  int64
		Level    int64
	}, error)
	GetAllTimeLeaderboard(limit int) ([]domain.User, error)
}

// UserRepository defines the interface for user data storage.
type UserRepository interface {
	Create(user *domain.User) error
	FindByID(id uuid.UUID) (*domain.User, error)
	FindByEmail(email string) (*domain.User, error)
	FindByUsername(username string) (*domain.User, error)
	Update(user *domain.User) error
	UpdateFields(id uuid.UUID, fields map[string]interface{}) error
}
