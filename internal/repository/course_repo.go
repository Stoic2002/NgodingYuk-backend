package repository

import (
	"github.com/arulkarim/ngodingyuk-server/internal/domain"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type CourseRepository struct {
	db *gorm.DB
}

func NewCourseRepository(db *gorm.DB) *CourseRepository {
	return &CourseRepository{db: db}
}

// ============ COURSES ============

// ListCourses returns courses with optional filtering by language, level, and search query.
func (r *CourseRepository) ListCourses(language, level, search string, limit, offset int) ([]domain.Course, int64, error) {
	query := r.db.Model(&domain.Course{})

	if language != "" {
		query = query.Where("language = ?", language)
	}
	if level != "" {
		query = query.Where("level = ?", level)
	}
	if search != "" {
		query = query.Where("title_id ILIKE ? OR title_en ILIKE ?", "%"+search+"%", "%"+search+"%")
	}

	var total int64
	if err := query.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	if limit > 0 {
		query = query.Limit(limit).Offset(offset)
	}

	var courses []domain.Course
	if err := query.Order("order_index ASC, created_at ASC").Find(&courses).Error; err != nil {
		return nil, 0, err
	}
	return courses, total, nil
}

// FindCourseBySlug returns a course by slug, optionally preloading lessons.
func (r *CourseRepository) FindCourseBySlug(slug string, preloadLessons bool) (*domain.Course, error) {
	query := r.db.Where("slug = ?", slug)
	if preloadLessons {
		query = query.Preload("Modules", func(db *gorm.DB) *gorm.DB {
			return db.Order("order_index ASC")
		}).Preload("Modules.Lessons", func(db *gorm.DB) *gorm.DB {
			return db.Order("order_index ASC")
		}).Preload("Lessons", func(db *gorm.DB) *gorm.DB {
			return db.Order("order_index ASC")
		})
	}
	var course domain.Course
	if err := query.First(&course).Error; err != nil {
		return nil, err
	}
	return &course, nil
}

// FindCourseByID returns a course by UUID.
func (r *CourseRepository) FindCourseByID(id uuid.UUID) (*domain.Course, error) {
	var course domain.Course
	if err := r.db.Where("id = ?", id).First(&course).Error; err != nil {
		return nil, err
	}
	return &course, nil
}

// CreateCourse inserts a new course.
func (r *CourseRepository) CreateCourse(course *domain.Course) error {
	return r.db.Create(course).Error
}

// UpdateCourse saves all fields of an existing course.
func (r *CourseRepository) UpdateCourse(course *domain.Course) error {
	return r.db.Save(course).Error
}

// DeleteCourse soft-deletes a course.
func (r *CourseRepository) DeleteCourse(id uuid.UUID) error {
	return r.db.Where("id = ?", id).Delete(&domain.Course{}).Error
}

// ============ LESSONS ============

// FindLessonByID returns a lesson by UUID, preloading quizzes.
func (r *CourseRepository) FindLessonByID(id uuid.UUID) (*domain.Lesson, error) {
	var lesson domain.Lesson
	if err := r.db.Preload("Quizzes", func(db *gorm.DB) *gorm.DB {
		return db.Order("order_index ASC")
	}).Where("id = ?", id).First(&lesson).Error; err != nil {
		return nil, err
	}
	return &lesson, nil
}

// ListLessonsByCourse returns all lessons for a course, ordered by order_index.
func (r *CourseRepository) ListLessonsByCourse(courseID uuid.UUID) ([]domain.Lesson, error) {
	var lessons []domain.Lesson
	if err := r.db.Where("course_id = ?", courseID).
		Order("order_index ASC").
		Find(&lessons).Error; err != nil {
		return nil, err
	}
	return lessons, nil
}

// CreateLesson inserts a new lesson.
func (r *CourseRepository) CreateLesson(lesson *domain.Lesson) error {
	return r.db.Create(lesson).Error
}

// UpdateLesson saves all fields of an existing lesson.
func (r *CourseRepository) UpdateLesson(lesson *domain.Lesson) error {
	return r.db.Save(lesson).Error
}

// DeleteLesson soft-deletes a lesson.
func (r *CourseRepository) DeleteLesson(id uuid.UUID) error {
	return r.db.Where("id = ?", id).Delete(&domain.Lesson{}).Error
}

// ============ LESSON QUIZZES ============

// FindQuizzesByLesson returns all quizzes for a lesson, ordered by order_index.
func (r *CourseRepository) FindQuizzesByLesson(lessonID uuid.UUID) ([]domain.LessonQuiz, error) {
	var quizzes []domain.LessonQuiz
	if err := r.db.Where("lesson_id = ?", lessonID).
		Order("order_index ASC").
		Find(&quizzes).Error; err != nil {
		return nil, err
	}
	return quizzes, nil
}

// FindQuizByID returns a single quiz by ID.
func (r *CourseRepository) FindQuizByID(id uuid.UUID) (*domain.LessonQuiz, error) {
	var quiz domain.LessonQuiz
	if err := r.db.Where("id = ?", id).First(&quiz).Error; err != nil {
		return nil, err
	}
	return &quiz, nil
}

// CreateQuiz inserts a new lesson quiz.
func (r *CourseRepository) CreateQuiz(quiz *domain.LessonQuiz) error {
	return r.db.Create(quiz).Error
}

// CountLessonsInCourse counts total lessons in a course.
func (r *CourseRepository) CountLessonsInCourse(courseID uuid.UUID) (int64, error) {
	var count int64
	if err := r.db.Model(&domain.Lesson{}).Where("course_id = ?", courseID).Count(&count).Error; err != nil {
		return 0, err
	}
	return count, nil
}

// ============ ENROLLMENT ============

// EnrollUser creates an enrollment record for (user, course).
func (r *CourseRepository) EnrollUser(userID, courseID uuid.UUID) error {
	enrollment := domain.UserCourseEnrollment{
		UserID:   userID,
		CourseID: courseID,
	}
	return r.db.Create(&enrollment).Error
}

// IsEnrolled checks if a user is enrolled in a course.
func (r *CourseRepository) IsEnrolled(userID, courseID uuid.UUID) (bool, error) {
	var count int64
	err := r.db.Model(&domain.UserCourseEnrollment{}).
		Where("user_id = ? AND course_id = ?", userID, courseID).
		Count(&count).Error
	return count > 0, err
}

// GetEnrollment returns the enrollment record.
func (r *CourseRepository) GetEnrollment(userID, courseID uuid.UUID) (*domain.UserCourseEnrollment, error) {
	var e domain.UserCourseEnrollment
	err := r.db.Where("user_id = ? AND course_id = ?", userID, courseID).First(&e).Error
	if err != nil {
		return nil, err
	}
	return &e, nil
}

// ============ CERTIFICATES ============

// CreateCertificate inserts a new certificate.
func (r *CourseRepository) CreateCertificate(cert *domain.Certificate) error {
	return r.db.Create(cert).Error
}

// GetCertificate returns a certificate for user + course.
func (r *CourseRepository) GetCertificate(userID, courseID uuid.UUID) (*domain.Certificate, error) {
	var cert domain.Certificate
	err := r.db.Where("user_id = ? AND course_id = ?", userID, courseID).First(&cert).Error
	if err != nil {
		return nil, err
	}
	return &cert, nil
}

// GetUserCertificates returns all certificates for a user.
func (r *CourseRepository) GetUserCertificates(userID uuid.UUID) ([]domain.Certificate, error) {
	var certs []domain.Certificate
	err := r.db.Preload("Course").Where("user_id = ?", userID).Order("passed_at DESC").Find(&certs).Error
	return certs, err
}

// ============ EXAM (all quizzes for a course) ============

// FindAllQuizzesByCourse returns all quizzes from all lessons in a course.
func (r *CourseRepository) FindAllQuizzesByCourse(courseID uuid.UUID) ([]domain.LessonQuiz, error) {
	var quizzes []domain.LessonQuiz
	err := r.db.Joins("JOIN lessons ON lessons.id = lesson_quizzes.lesson_id").
		Where("lessons.course_id = ? AND lessons.deleted_at IS NULL", courseID).
		Order("lessons.order_index ASC, lesson_quizzes.order_index ASC").
		Find(&quizzes).Error
	return quizzes, err
}
