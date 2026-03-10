package repository

import (
	"encoding/json"
	"time"

	"github.com/arulkarim/ngodingyuk-server/internal/domain"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type ProgressRepository struct {
	db *gorm.DB
}

func NewProgressRepository(db *gorm.DB) *ProgressRepository {
	return &ProgressRepository{db: db}
}

// ============ CHALLENGE PROGRESS ============

// FindChallengeProgress returns the user's progress for a specific challenge.
func (r *ProgressRepository) FindChallengeProgress(userID, challengeID uuid.UUID) (*domain.UserChallengeProgress, error) {
	var progress domain.UserChallengeProgress
	if err := r.db.Where("user_id = ? AND challenge_id = ?", userID, challengeID).
		First(&progress).Error; err != nil {
		return nil, err
	}
	return &progress, nil
}

// GetCompletedChallengeIDs returns a list of challenge IDs the user has solved.
func (r *ProgressRepository) GetCompletedChallengeIDs(userID uuid.UUID) ([]uuid.UUID, error) {
	var ids []uuid.UUID
	err := r.db.Model(&domain.UserChallengeProgress{}).
		Where("user_id = ? AND status = ?", userID, "solved").
		Pluck("challenge_id", &ids).Error
	return ids, err
}

// FindOrCreateChallengeProgress finds existing progress or creates a new one.
func (r *ProgressRepository) FindOrCreateChallengeProgress(userID, challengeID uuid.UUID) (*domain.UserChallengeProgress, error) {
	var progress domain.UserChallengeProgress
	err := r.db.Where("user_id = ? AND challenge_id = ?", userID, challengeID).First(&progress).Error
	if err == gorm.ErrRecordNotFound {
		progress = domain.UserChallengeProgress{
			UserID:      userID,
			ChallengeID: challengeID,
			Status:      "attempted",
			Attempts:    0,
		}
		if err := r.db.Create(&progress).Error; err != nil {
			return nil, err
		}
	} else if err != nil {
		return nil, err
	}
	return &progress, nil
}

// UpdateChallengeProgress saves the challenge progress record.
func (r *ProgressRepository) UpdateChallengeProgress(progress *domain.UserChallengeProgress) error {
	return r.db.Save(progress).Error
}

// GetUserChallengeStats returns solved/attempted counts grouped by difficulty.
func (r *ProgressRepository) GetUserChallengeStats(userID uuid.UUID) ([]struct {
	Difficulty string
	Status     string
	Count      int64
}, error) {
	type result struct {
		Difficulty string
		Status     string
		Count      int64
	}
	var results []result

	err := r.db.Table("user_challenge_progress").
		Select("challenges.difficulty, user_challenge_progress.status, COUNT(*) as count").
		Joins("JOIN challenges ON challenges.id = user_challenge_progress.challenge_id").
		Where("user_challenge_progress.user_id = ?", userID).
		Group("challenges.difficulty, user_challenge_progress.status").
		Scan(&results).Error
	if err != nil {
		return nil, err
	}

	// Convert to the expected return type
	var out []struct {
		Difficulty string
		Status     string
		Count      int64
	}
	for _, r := range results {
		out = append(out, struct {
			Difficulty string
			Status     string
			Count      int64
		}{r.Difficulty, r.Status, r.Count})
	}
	return out, nil
}

// ============ LESSON PROGRESS ============

// FindLessonProgress returns the user's progress for a specific lesson.
func (r *ProgressRepository) FindLessonProgress(userID, lessonID uuid.UUID) (*domain.UserLessonProgress, error) {
	var progress domain.UserLessonProgress
	if err := r.db.Where("user_id = ? AND lesson_id = ?", userID, lessonID).
		First(&progress).Error; err != nil {
		return nil, err
	}
	return &progress, nil
}

// FindOrCreateLessonProgress finds existing progress or creates a new one.
func (r *ProgressRepository) FindOrCreateLessonProgress(userID, lessonID uuid.UUID) (*domain.UserLessonProgress, error) {
	var progress domain.UserLessonProgress
	err := r.db.Where("user_id = ? AND lesson_id = ?", userID, lessonID).First(&progress).Error
	if err == gorm.ErrRecordNotFound {
		progress = domain.UserLessonProgress{
			UserID:   userID,
			LessonID: lessonID,
		}
		if err := r.db.Create(&progress).Error; err != nil {
			return nil, err
		}
	} else if err != nil {
		return nil, err
	}
	return &progress, nil
}

// UpdateLessonProgress saves the lesson progress record.
func (r *ProgressRepository) UpdateLessonProgress(progress *domain.UserLessonProgress) error {
	return r.db.Save(progress).Error
}

// CountCompletedLessonsInCourse counts how many lessons the user completed in a course.
func (r *ProgressRepository) CountCompletedLessonsInCourse(userID, courseID uuid.UUID) (int64, error) {
	var count int64
	err := r.db.Table("user_lesson_progress").
		Joins("JOIN lessons ON lessons.id = user_lesson_progress.lesson_id").
		Where("user_lesson_progress.user_id = ? AND lessons.course_id = ? AND user_lesson_progress.completed_at IS NOT NULL", userID, courseID).
		Count(&count).Error
	return count, err
}

// GetUserLessonProgressByCourse returns all lesson progress records for a user in a course.
func (r *ProgressRepository) GetUserLessonProgressByCourse(userID, courseID uuid.UUID) ([]domain.UserLessonProgress, error) {
	var progress []domain.UserLessonProgress
	err := r.db.Joins("JOIN lessons ON lessons.id = user_lesson_progress.lesson_id").
		Where("user_lesson_progress.user_id = ? AND lessons.course_id = ?", userID, courseID).
		Find(&progress).Error
	return progress, err
}

// ============ XP HISTORY ============

// CreateXPHistory logs an XP gain event.
func (r *ProgressRepository) CreateXPHistory(history *domain.UserXPHistory) error {
	return r.db.Create(history).Error
}

// GetUserXPHistory returns paginated XP history for a user.
func (r *ProgressRepository) GetUserXPHistory(userID uuid.UUID, limit, offset int) ([]domain.UserXPHistory, error) {
	var history []domain.UserXPHistory
	if err := r.db.Where("user_id = ?", userID).
		Order("created_at DESC").
		Limit(limit).Offset(offset).
		Find(&history).Error; err != nil {
		return nil, err
	}
	return history, nil
}

// GetUserQuizHistory returns a combined history of completed lesson quizzes and course exams.
func (r *ProgressRepository) GetUserQuizHistory(userID uuid.UUID) ([]domain.QuizHistoryItem, error) {
	var results []domain.QuizHistoryItem

	// 1. Get Lesson Quizzes
	// Join user_lesson_progress (where quiz_completed = true) -> lessons -> courses
	// To get TotalQuestions, we can join lesson_quizzes or just return 0 if complicated.
	// Actually, let's use a subquery for TotalQuestions.
	type lessonResult struct {
		CourseTitle    string          `gorm:"column:course_title"`
		LessonTitle    string          `gorm:"column:lesson_title"`
		Score          int             `gorm:"column:score"`
		TotalQuestions int             `gorm:"column:total_questions"`
		ResultDetails  json.RawMessage `gorm:"column:result_details"`
		CompletedAt    time.Time       `gorm:"column:completed_at"`
	}
	var lessons []lessonResult
	err := r.db.Table("user_lesson_progress").
		Select("courses.title_en as course_title, lessons.title_en as lesson_title, user_lesson_progress.quiz_score as score, (SELECT COUNT(*) FROM lesson_quizzes WHERE lesson_quizzes.lesson_id = lessons.id) as total_questions, user_lesson_progress.result_details as result_details, user_lesson_progress.completed_at as completed_at").
		Joins("JOIN lessons ON lessons.id = user_lesson_progress.lesson_id").
		Joins("JOIN courses ON courses.id = lessons.course_id").
		Where("user_lesson_progress.user_id = ? AND user_lesson_progress.quiz_completed = ?", userID, true).
		Scan(&lessons).Error
	if err == nil {
		for _, l := range lessons {
			results = append(results, domain.QuizHistoryItem{
				Type:           "lesson_quiz",
				CourseTitle:    l.CourseTitle,
				LessonTitle:    l.LessonTitle,
				Score:          l.Score,
				TotalQuestions: l.TotalQuestions,
				ResultDetails:  l.ResultDetails,
				PassedAt:       l.CompletedAt,
			})
		}
	}

	// 2. Get Course Exams (Certificates)
	type examResult struct {
		CourseTitle    string          `gorm:"column:course_title"`
		Score          int             `gorm:"column:score"`
		TotalQuestions int             `gorm:"column:total_questions"`
		ResultDetails  json.RawMessage `gorm:"column:result_details"`
		PassedAt       time.Time       `gorm:"column:passed_at"`
	}
	var exams []examResult
	err = r.db.Table("certificates").
		Select("courses.title_en as course_title, certificates.score as score, certificates.total_questions as total_questions, certificates.result_details as result_details, certificates.passed_at as passed_at").
		Joins("JOIN courses ON courses.id = certificates.course_id").
		Where("certificates.user_id = ?", userID).
		Scan(&exams).Error
	if err == nil {
		for _, e := range exams {
			results = append(results, domain.QuizHistoryItem{
				Type:           "course_exam",
				CourseTitle:    e.CourseTitle,
				Score:          e.Score,
				TotalQuestions: e.TotalQuestions,
				ResultDetails:  e.ResultDetails,
				PassedAt:       e.PassedAt,
			})
		}
	}

	return results, nil
}

// ============ LEADERBOARD ============

// GetWeeklyLeaderboard returns top N users by XP gained in the last 7 days.
func (r *ProgressRepository) GetWeeklyLeaderboard(limit int) ([]struct {
	UserID   uuid.UUID
	Username string
	TotalXP  int64
	Level    int64
}, error) {
	type result struct {
		UserID   uuid.UUID
		Username string
		TotalXP  int64
		Level    int64
	}
	var results []result

	weekAgo := time.Now().AddDate(0, 0, -7)
	err := r.db.Table("user_xp_history").
		Select("user_xp_history.user_id, users.username, SUM(user_xp_history.xp_gained) as total_xp, users.level").
		Joins("JOIN users ON users.id = user_xp_history.user_id").
		Where("user_xp_history.created_at >= ?", weekAgo).
		Group("user_xp_history.user_id, users.username, users.level").
		Order("total_xp DESC").
		Limit(limit).
		Scan(&results).Error
	if err != nil {
		return nil, err
	}

	var out []struct {
		UserID   uuid.UUID
		Username string
		TotalXP  int64
		Level    int64
	}
	for _, r := range results {
		out = append(out, struct {
			UserID   uuid.UUID
			Username string
			TotalXP  int64
			Level    int64
		}{r.UserID, r.Username, r.TotalXP, r.Level})
	}
	return out, nil
}

// GetAllTimeLeaderboard returns top N users by total XP.
func (r *ProgressRepository) GetAllTimeLeaderboard(limit int) ([]domain.User, error) {
	var users []domain.User
	if err := r.db.Select("id, username, xp, level, streak_count").
		Order("xp DESC").
		Limit(limit).
		Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}
