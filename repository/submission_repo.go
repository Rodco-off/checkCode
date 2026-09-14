package repository

import (
	"github.com/Rodco-off/checkCode/domain"
	"gorm.io/gorm"
)

type SubmissionRepository struct {
	db *gorm.DB
}

func NewSubmissionRepository(db *gorm.DB) *SubmissionRepository {
	return &SubmissionRepository{db: db}
}

func (repo *SubmissionRepository) Create(sub *domain.Submission) error {
	return repo.db.Create(sub).Error
}

func (repo *SubmissionRepository) GetBySession(sessionID string) ([]domain.Submission, error) {
	var subs []domain.Submission

	err := repo.db.Where("session = ?", sessionID).
		Order("create_at DESC").Find(&subs).Error

	return subs, err
}
