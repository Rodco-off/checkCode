package domain

import "encoding/json"

type Task struct {
	TaskID      uint   `gorm:"primaryKey"`
	Title       string `gorm:"not null"`
	Description string `gorm:"not null"`
	Difficulty  string `gorm:"default:easy"`
	Constraints string `gorm:"type:text"`
	TestCases   []byte `gorm:"type:jsonb;not null"`
}

func (task *Task) ParseTestCases() ([]TestCase, error) {
	var cases []TestCase
	if err := json.Unmarshal(task.TestCases, &cases); err != nil {
		return nil, err
	}

	return cases, nil
}
