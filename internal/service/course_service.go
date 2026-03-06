package service

import (
	"encoding/json"
	"errors"
	"time"

	"github.com/arulkarim/ngodingyuk-server/internal/domain"
	"github.com/arulkarim/ngodingyuk-server/internal/repository"
	"github.com/arulkarim/ngodingyuk-server/pkg/i18n"
	"github.com/google/uuid"
)

type CourseService struct {
	courseRepo   *repository.CourseRepository
	progressRepo *repository.ProgressRepository
	userRepo     *repository.UserRepository
}

func NewCourseService(
	courseRepo *repository.CourseRepository,
	progressRepo *repository.ProgressRepository,
	userRepo *repository.UserRepository,
) *CourseService {
	return &CourseService{
		courseRepo:   courseRepo,
		progressRepo: progressRepo,
		userRepo:     userRepo,
	}
}

// === Response DTOs ===

type CourseListItem struct {
	ID          string `json:"id"`
	Slug        string `json:"slug"`
	Language    string `json:"language"`
	Level       string `json:"level"`
	Title       string `json:"title"`
	Description string `json:"description,omitempty"`
	Thumbnail   string `json:"thumbnail_url,omitempty"`
	OrderIndex  int    `json:"order_index"`
}

type CourseDetailResponse struct {
	ID          string           `json:"id"`
	Slug        string           `json:"slug"`
	Language    string           `json:"language"`
	Level       string           `json:"level"`
	Title       string           `json:"title"`
	Description string           `json:"description,omitempty"`
	Thumbnail   string           `json:"thumbnail_url,omitempty"`
	IsEnrolled  bool             `json:"is_enrolled"`
	HasCert     bool             `json:"has_certificate"`
	Lessons     []LessonListItem `json:"lessons"`
}

type LessonListItem struct {
	ID         string `json:"id"`
	Title      string `json:"title"`
	OrderIndex int    `json:"order_index"`
	XPReward   int64  `json:"xp_reward"`
	IsLocked   bool   `json:"is_locked"`
	Completed  bool   `json:"completed"`
}

type LessonDetailResponse struct {
	ID              string         `json:"id"`
	CourseID        string         `json:"course_id"`
	Title           string         `json:"title"`
	ContentMarkdown string         `json:"content_markdown"`
	OrderIndex      int            `json:"order_index"`
	XPReward        int64          `json:"xp_reward"`
	Quizzes         []QuizResponse `json:"quizzes"`
}

type QuizResponse struct {
	ID         string          `json:"id"`
	Question   string          `json:"question"`
	Options    json.RawMessage `json:"options"`
	OrderIndex int             `json:"order_index"`
	// correct_index is NOT sent to client
}

type QuizSubmitRequest struct {
	Answers []int `json:"answers"` // user's answer_index for each quiz, in order
}

type QuizSubmitResponse struct {
	Score    int              `json:"score"`
	Total    int              `json:"total"`
	XPGained int64            `json:"xp_gained"`
	Details  []QuizResultItem `json:"details"`
}

type QuizResultItem struct {
	QuizID       string `json:"quiz_id"`
	Correct      bool   `json:"correct"`
	CorrectIndex int    `json:"correct_index"`
	Explanation  string `json:"explanation,omitempty"`
}

type CourseProgressResponse struct {
	CompletedLessons int64                `json:"completed_lessons"`
	TotalLessons     int64                `json:"total_lessons"`
	Lessons          []LessonProgressItem `json:"lessons"`
}

type LessonProgressItem struct {
	LessonID      string `json:"lesson_id"`
	QuizCompleted bool   `json:"quiz_completed"`
	QuizScore     int    `json:"quiz_score"`
	CompletedAt   string `json:"completed_at,omitempty"`
}

// === Exam DTOs ===

type ExamResponse struct {
	CourseTitle    string         `json:"course_title"`
	Quizzes        []QuizResponse `json:"quizzes"`
	TotalQuestions int            `json:"total_questions"`
	TimeLimitSec   int            `json:"time_limit_sec"` // seconds
}

type ExamSubmitRequest struct {
	Answers []int `json:"answers"` // answer indices in quiz order
}

type ExamSubmitResponse struct {
	Score    int              `json:"score"`
	Total    int              `json:"total"`
	Passed   bool             `json:"passed"`
	XPGained int64            `json:"xp_gained"`
	CertID   string           `json:"certificate_id,omitempty"`
	Details  []QuizResultItem `json:"details"`
}

// === Service Methods ===

// ListCourses returns all courses filtered by language/level.
func (s *CourseService) ListCourses(language, level, locale string) ([]CourseListItem, error) {
	courses, err := s.courseRepo.ListCourses(language, level)
	if err != nil {
		return nil, err
	}

	var items []CourseListItem
	for _, c := range courses {
		items = append(items, CourseListItem{
			ID:          c.ID.String(),
			Slug:        c.Slug,
			Language:    c.Language,
			Level:       c.Level,
			Title:       i18n.Resolve(locale, c.TitleID, c.TitleEN),
			Description: i18n.ResolveOptional(locale, c.DescriptionID, c.DescriptionEN),
			Thumbnail:   ptrToString(c.ThumbnailURL),
			OrderIndex:  c.OrderIndex,
		})
	}
	return items, nil
}

// GetCourseBySlug returns course detail with lesson list.
// userID is optional (zero = not logged in).
func (s *CourseService) GetCourseBySlug(slug, locale string, userID uuid.UUID) (*CourseDetailResponse, error) {
	course, err := s.courseRepo.FindCourseBySlug(slug, true) // preload lessons
	if err != nil {
		return nil, errors.New("course not found")
	}

	// Check enrollment
	isEnrolled := false
	hasCert := false
	var completedLessonIDs map[string]bool
	if userID != uuid.Nil {
		isEnrolled, _ = s.courseRepo.IsEnrolled(userID, course.ID)
		_, certErr := s.courseRepo.GetCertificate(userID, course.ID)
		hasCert = certErr == nil

		// Get completed lessons for locking
		completedLessonIDs = make(map[string]bool)
		progressList, _ := s.progressRepo.GetUserLessonProgressByCourse(userID, course.ID)
		for _, p := range progressList {
			if p.CompletedAt != nil {
				completedLessonIDs[p.LessonID.String()] = true
			}
		}
	}

	var lessons []LessonListItem
	for i, l := range course.Lessons {
		// Lesson locking: lesson 0 always unlocked, lesson N locked if lesson N-1 not completed
		isLocked := false
		completed := false
		if userID != uuid.Nil && isEnrolled {
			completed = completedLessonIDs[l.ID.String()]
			if i > 0 {
				prevID := course.Lessons[i-1].ID.String()
				if !completedLessonIDs[prevID] {
					isLocked = true
				}
			}
		} else if userID != uuid.Nil && !isEnrolled {
			// Not enrolled = all locked
			isLocked = true
		}

		lessons = append(lessons, LessonListItem{
			ID:         l.ID.String(),
			Title:      i18n.Resolve(locale, l.TitleID, l.TitleEN),
			OrderIndex: l.OrderIndex,
			XPReward:   l.XPReward,
			IsLocked:   isLocked,
			Completed:  completed,
		})
	}

	return &CourseDetailResponse{
		ID:          course.ID.String(),
		Slug:        course.Slug,
		Language:    course.Language,
		Level:       course.Level,
		Title:       i18n.Resolve(locale, course.TitleID, course.TitleEN),
		Description: i18n.ResolveOptional(locale, course.DescriptionID, course.DescriptionEN),
		Thumbnail:   ptrToString(course.ThumbnailURL),
		IsEnrolled:  isEnrolled,
		HasCert:     hasCert,
		Lessons:     lessons,
	}, nil
}

// EnrollCourse enrolls a user in a course.
func (s *CourseService) EnrollCourse(userID uuid.UUID, courseSlug string) error {
	course, err := s.courseRepo.FindCourseBySlug(courseSlug, false)
	if err != nil {
		return errors.New("course not found")
	}
	already, _ := s.courseRepo.IsEnrolled(userID, course.ID)
	if already {
		return errors.New("already enrolled")
	}
	return s.courseRepo.EnrollUser(userID, course.ID)
}

// GetExam returns all quizzes for a course's final exam.
func (s *CourseService) GetExam(userID uuid.UUID, courseSlug, locale string) (*ExamResponse, error) {
	course, err := s.courseRepo.FindCourseBySlug(courseSlug, false)
	if err != nil {
		return nil, errors.New("course not found")
	}

	// Must be enrolled
	enrolled, _ := s.courseRepo.IsEnrolled(userID, course.ID)
	if !enrolled {
		return nil, errors.New("not enrolled")
	}

	// Check if already has certificate
	_, certErr := s.courseRepo.GetCertificate(userID, course.ID)
	if certErr == nil {
		return nil, errors.New("already passed")
	}

	quizzes, err := s.courseRepo.FindAllQuizzesByCourse(course.ID)
	if err != nil || len(quizzes) == 0 {
		return nil, errors.New("no quizzes found for this course")
	}

	var items []QuizResponse
	for _, q := range quizzes {
		options := q.OptionsID
		if locale == "en" && q.OptionsEN != nil && len(q.OptionsEN) > 0 {
			options = q.OptionsEN
		}
		items = append(items, QuizResponse{
			ID:         q.ID.String(),
			Question:   i18n.Resolve(locale, q.QuestionID, q.QuestionEN),
			Options:    options,
			OrderIndex: q.OrderIndex,
		})
	}

	// Time limit: 1 minute per question
	timeLimit := len(quizzes) * 60

	return &ExamResponse{
		CourseTitle:    i18n.Resolve(locale, course.TitleID, course.TitleEN),
		Quizzes:        items,
		TotalQuestions: len(quizzes),
		TimeLimitSec:   timeLimit,
	}, nil
}

// SubmitExam validates exam answers, awards XP, and issues certificate if passed (>=70%).
func (s *CourseService) SubmitExam(userID uuid.UUID, courseSlug string, req ExamSubmitRequest) (*ExamSubmitResponse, error) {
	course, err := s.courseRepo.FindCourseBySlug(courseSlug, false)
	if err != nil {
		return nil, errors.New("course not found")
	}

	enrolled, _ := s.courseRepo.IsEnrolled(userID, course.ID)
	if !enrolled {
		return nil, errors.New("not enrolled")
	}

	// Already has certificate?
	_, certErr := s.courseRepo.GetCertificate(userID, course.ID)
	if certErr == nil {
		return nil, errors.New("already passed")
	}

	quizzes, err := s.courseRepo.FindAllQuizzesByCourse(course.ID)
	if err != nil {
		return nil, errors.New("no quizzes")
	}

	if len(req.Answers) != len(quizzes) {
		return nil, errors.New("answer count mismatch")
	}

	score := 0
	var details []QuizResultItem
	for i, quiz := range quizzes {
		correct := req.Answers[i] == quiz.CorrectIndex
		if correct {
			score++
		}
		explanation := i18n.ResolveOptional("id", quiz.ExplanationID, quiz.ExplanationEN)
		details = append(details, QuizResultItem{
			QuizID:       quiz.ID.String(),
			Correct:      correct,
			CorrectIndex: quiz.CorrectIndex,
			Explanation:  explanation,
		})
	}

	total := len(quizzes)
	passed := float64(score)/float64(total) >= 0.70

	var xpGained int64
	var certID string

	if passed {
		// Award XP for all quizzes combined
		for _, q := range quizzes {
			xpGained += q.XPReward
		}

		user, err := s.userRepo.FindByID(userID)
		if err == nil {
			user.XP += xpGained
			user.Level = CalculateLevel(user.XP)
			UpdateStreak(user)
			s.userRepo.Update(user)

			courseID := course.ID
			s.progressRepo.CreateXPHistory(&domain.UserXPHistory{
				UserID:     userID,
				XPGained:   xpGained,
				SourceType: "lesson_quiz",
				SourceID:   &courseID,
			})
		}

		// Create certificate
		cert := &domain.Certificate{
			UserID:         userID,
			CourseID:       course.ID,
			Score:          score,
			TotalQuestions: total,
		}
		s.courseRepo.CreateCertificate(cert)
		certID = cert.ID.String()
	}

	return &ExamSubmitResponse{
		Score:    score,
		Total:    total,
		Passed:   passed,
		XPGained: xpGained,
		CertID:   certID,
		Details:  details,
	}, nil
}

// GetLessonDetail returns lesson content + quizzes (without correct_index).
func (s *CourseService) GetLessonDetail(lessonID uuid.UUID, locale string) (*LessonDetailResponse, error) {
	lesson, err := s.courseRepo.FindLessonByID(lessonID) // preloads quizzes
	if err != nil {
		return nil, errors.New("lesson not found")
	}

	var quizzes []QuizResponse
	for _, q := range lesson.Quizzes {
		// Resolve bilingual options
		options := q.OptionsID
		if locale == "en" && q.OptionsEN != nil && len(q.OptionsEN) > 0 {
			options = q.OptionsEN
		}

		quizzes = append(quizzes, QuizResponse{
			ID:         q.ID.String(),
			Question:   i18n.Resolve(locale, q.QuestionID, q.QuestionEN),
			Options:    options,
			OrderIndex: q.OrderIndex,
		})
	}

	return &LessonDetailResponse{
		ID:              lesson.ID.String(),
		CourseID:        lesson.CourseID.String(),
		Title:           i18n.Resolve(locale, lesson.TitleID, lesson.TitleEN),
		ContentMarkdown: i18n.Resolve(locale, lesson.ContentMarkdownID, lesson.ContentMarkdownEN),
		OrderIndex:      lesson.OrderIndex,
		XPReward:        lesson.XPReward,
		Quizzes:         quizzes,
	}, nil
}

// CompleteLesson marks a lesson as completed and awards XP.
func (s *CourseService) CompleteLesson(userID, lessonID uuid.UUID) error {
	lesson, err := s.courseRepo.FindLessonByID(lessonID)
	if err != nil {
		return errors.New("lesson not found")
	}

	progress, err := s.progressRepo.FindOrCreateLessonProgress(userID, lessonID)
	if err != nil {
		return err
	}

	if progress.CompletedAt == nil {
		now := time.Now()
		progress.CompletedAt = &now
		s.progressRepo.UpdateLessonProgress(progress)

		// Award XP
		user, err := s.userRepo.FindByID(userID)
		if err == nil {
			user.XP += lesson.XPReward
			user.Level = CalculateLevel(user.XP)
			UpdateStreak(user)
			s.userRepo.Update(user)

			s.progressRepo.CreateXPHistory(&domain.UserXPHistory{
				UserID:     userID,
				XPGained:   lesson.XPReward,
				SourceType: "lesson_quiz",
				SourceID:   &lesson.ID,
			})
		}
	}

	return nil
}

// SubmitQuiz validates quiz answers and awards XP for correct answers.
func (s *CourseService) SubmitQuiz(userID, lessonID uuid.UUID, req QuizSubmitRequest) (*QuizSubmitResponse, error) {
	quizzes, err := s.courseRepo.FindQuizzesByLesson(lessonID)
	if err != nil {
		return nil, errors.New("quizzes not found")
	}

	if len(req.Answers) != len(quizzes) {
		return nil, errors.New("answer count must match quiz count")
	}

	score := 0
	var totalXP int64
	var details []QuizResultItem

	for i, quiz := range quizzes {
		correct := req.Answers[i] == quiz.CorrectIndex
		if correct {
			score++
			totalXP += quiz.XPReward
		}

		explanation := i18n.ResolveOptional("id", quiz.ExplanationID, quiz.ExplanationEN)

		details = append(details, QuizResultItem{
			QuizID:       quiz.ID.String(),
			Correct:      correct,
			CorrectIndex: quiz.CorrectIndex,
			Explanation:  explanation,
		})
	}

	// Update progress
	progress, err := s.progressRepo.FindOrCreateLessonProgress(userID, lessonID)
	if err == nil {
		progress.QuizCompleted = true
		progress.QuizScore = score

		if progress.CompletedAt == nil {
			now := time.Now()
			progress.CompletedAt = &now
		}
		s.progressRepo.UpdateLessonProgress(progress)
	}

	// Award XP
	if totalXP > 0 {
		user, err := s.userRepo.FindByID(userID)
		if err == nil {
			user.XP += totalXP
			user.Level = CalculateLevel(user.XP)
			UpdateStreak(user)
			s.userRepo.Update(user)

			s.progressRepo.CreateXPHistory(&domain.UserXPHistory{
				UserID:     userID,
				XPGained:   totalXP,
				SourceType: "lesson_quiz",
				SourceID:   &lessonID,
			})
		}
	}

	return &QuizSubmitResponse{
		Score:    score,
		Total:    len(quizzes),
		XPGained: totalXP,
		Details:  details,
	}, nil
}

// GetMyProgress returns the user's progress for a course.
func (s *CourseService) GetMyProgress(courseSlug string, userID uuid.UUID) (*CourseProgressResponse, error) {
	course, err := s.courseRepo.FindCourseBySlug(courseSlug, false)
	if err != nil {
		return nil, errors.New("course not found")
	}

	totalLessons, _ := s.courseRepo.CountLessonsInCourse(course.ID)
	completedLessons, _ := s.progressRepo.CountCompletedLessonsInCourse(userID, course.ID)

	progressList, _ := s.progressRepo.GetUserLessonProgressByCourse(userID, course.ID)

	var items []LessonProgressItem
	for _, p := range progressList {
		item := LessonProgressItem{
			LessonID:      p.LessonID.String(),
			QuizCompleted: p.QuizCompleted,
			QuizScore:     p.QuizScore,
		}
		if p.CompletedAt != nil {
			item.CompletedAt = p.CompletedAt.Format(time.RFC3339)
		}
		items = append(items, item)
	}

	return &CourseProgressResponse{
		CompletedLessons: completedLessons,
		TotalLessons:     totalLessons,
		Lessons:          items,
	}, nil
}

// Helper
func ptrToString(s *string) string {
	if s == nil {
		return ""
	}
	return *s
}

// IsLessonAccessible checks if a lesson is accessible (not locked) for a user.
func (s *CourseService) IsLessonAccessible(userID, lessonID uuid.UUID) (bool, error) {
	lesson, err := s.courseRepo.FindLessonByID(lessonID)
	if err != nil {
		return false, errors.New("lesson not found")
	}

	// Check enrollment
	enrolled, _ := s.courseRepo.IsEnrolled(userID, lesson.CourseID)
	if !enrolled {
		return false, nil
	}

	// Get all lessons in course
	lessons, err := s.courseRepo.ListLessonsByCourse(lesson.CourseID)
	if err != nil {
		return false, err
	}

	for i, l := range lessons {
		if l.ID == lessonID {
			if i == 0 {
				return true, nil // first lesson always accessible
			}
			// Check if previous lesson is completed
			prevProgress, err := s.progressRepo.FindLessonProgress(userID, lessons[i-1].ID)
			if err != nil || prevProgress.CompletedAt == nil {
				return false, nil
			}
			return true, nil
		}
	}
	return false, nil
}
