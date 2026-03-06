package repository

import (
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
