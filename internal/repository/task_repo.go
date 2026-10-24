package repository

import (
	"errors"

	"github.com/Rodco-off/checkCode/internal/domain"
	"gorm.io/gorm"
)

var ErrTaskNotFound error = errors.New("Задача не найдена")

type TaskRepository struct {
	db *gorm.DB
}

func NewTaskRepository(db *gorm.DB) *TaskRepository {
	return &TaskRepository{db: db}
}

func (repo *TaskRepository) GetAll(difficulty string) ([]domain.Task, error) {
	var tasks []domain.Task
	var err error

	if difficulty == "" {
		err = repo.db.Find(&tasks).Error
	} else {
		query := repo.db.Model(&domain.Task{})
		query = query.Where("difficulty = ?", difficulty)
		err = query.Find(&tasks).Error
	}

	return tasks, err
}

func (repo *TaskRepository) GetByID(id uint) (*domain.Task, error) {
	var task domain.Task

	err := repo.db.First(&task, id).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, ErrTaskNotFound
	}

	if err != nil {
		return nil, err
	}

	return &task, nil
}
