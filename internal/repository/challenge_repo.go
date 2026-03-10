package repository

import (
	"github.com/arulkarim/ngodingyuk-server/internal/domain"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type ChallengeRepository struct {
	db *gorm.DB
}

func NewChallengeRepository(db *gorm.DB) *ChallengeRepository {
	return &ChallengeRepository{db: db}
}

// List returns challenges with optional filtering by language, difficulty, and search query.
// Supports pagination via limit/offset.
func (r *ChallengeRepository) List(language, difficulty, search string, limit, offset int) ([]domain.Challenge, int64, error) {
	query := r.db.Model(&domain.Challenge{})

	if language != "" {
		query = query.Where("language = ?", language)
	}
	if difficulty != "" {
		query = query.Where("difficulty = ?", difficulty)
	}
	if search != "" {
		query = query.Where("title_id ILIKE ? OR title_en ILIKE ?", "%"+search+"%", "%"+search+"%")
	}

	var total int64
	if err := query.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	var challenges []domain.Challenge
	if err := query.Order("order_index ASC, created_at ASC").
		Limit(limit).Offset(offset).
		Find(&challenges).Error; err != nil {
		return nil, 0, err
	}

	return challenges, total, nil
}

// FindBySlug returns a single challenge by its slug.
func (r *ChallengeRepository) FindBySlug(slug string) (*domain.Challenge, error) {
	var challenge domain.Challenge
	if err := r.db.Where("slug = ?", slug).First(&challenge).Error; err != nil {
		return nil, err
	}
	return &challenge, nil
}

// FindByID returns a single challenge by its UUID.
func (r *ChallengeRepository) FindByID(id uuid.UUID) (*domain.Challenge, error) {
	var challenge domain.Challenge
	if err := r.db.Where("id = ?", id).First(&challenge).Error; err != nil {
		return nil, err
	}
	return &challenge, nil
}

// Create inserts a new challenge.
func (r *ChallengeRepository) Create(challenge *domain.Challenge) error {
	return r.db.Create(challenge).Error
}

// Update saves all fields of an existing challenge.
func (r *ChallengeRepository) Update(challenge *domain.Challenge) error {
	return r.db.Save(challenge).Error
}

// Delete soft-deletes a challenge.
func (r *ChallengeRepository) Delete(id uuid.UUID) error {
	return r.db.Where("id = ?", id).Delete(&domain.Challenge{}).Error
}
