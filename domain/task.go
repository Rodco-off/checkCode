package domain

type Task struct {
	TaskID      int
	Title       string
	Description string
	Difficulty  string
	TestCase    []byte
}
