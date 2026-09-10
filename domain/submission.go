package domain

import "time"

type Submission struct {
	ID           int
	TaskID       int
	SessionID    int
	Code         string
	Status       string
	Output       string
	Expected     string
	ErrorMessage string
	// Hint string // запас под нейронку
	CreatedAt time.Time
}

// Для Поля Status в структуре
const (
	StatusError   = "Error"
	StatusPending = "Pending"
	StatusFailed  = "Failed"
	StatusSuccess = "Success"
)
