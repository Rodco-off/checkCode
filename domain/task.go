package domain

type Task struct {
	TaskID      int    `gorm:"primaryKey"`
	Title       string `gorm:"not null"`
	Description string `gorm:"not null"`
	Difficulty  string `gorm:"default:easy"`
	TestCase    []byte `gorm:"type:jsonb;not null"`
}
