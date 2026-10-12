package domain

import "encoding/json"

type Task struct {
	TaskID      int    `gorm:"primaryKey"`
	Title       string `gorm:"not null"`
	Description string `gorm:"not null"`
	Difficulty  string `gorm:"default:easy"`
	TestCases   []byte `gorm:"type:jsonb;not null"`
}

func (task *Task) ParseTestCase() ([]TestCase, error) {
	var cases []TestCase
	if err := json.Unmarshal(task.TestCases, &cases); err != nil {
		return nil, err
	}

	return cases, nil
}
