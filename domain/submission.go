package domain

import "time"

type Submission struct {
	ID           uint   `gorm:"primaryKey"`
	TaskID       uint   `gorm:"not null;index"`
	SessionID    uint   `gorm:"not null;index"`
	Code         string `gorm:"type:text;not null"`
	Status       string `gorm:"default:pending"`
	Output       string `gorm:"type:text"`
	Expected     string `gorm:"type:text"`
	ErrorMessage string `gorm:"type:text"`
	// Hint string // запас под нейронку
	CreatedAt time.Time `gorm:"autoCreateTime"`
}

// Для Поля Status в структуре
const (
	StatusError   = "Error"
	StatusPending = "Pending"
	StatusFailed  = "Failed"
	StatusSuccess = "Success"
)
