package handler

type SubmitRequest struct {
	TaskID uint   `json:"task_id"`
	Code   string `json:"code"`
}
