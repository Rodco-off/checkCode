package handler

import (
	"net/http"

	"github.com/Rodco-off/checkCode/domain"
	"github.com/Rodco-off/checkCode/repository"
	"github.com/labstack/echo"
)

type SubmitHandler struct {
	subRepo *repository.SubmissionRepository
}

func NewSubmissionHandler(subRepo *repository.SubmissionRepository) *SubmitHandler {
	return &SubmitHandler{subRepo: subRepo}
}

func (handl *SubmitHandler) Submit(cont echo.Context) error {
	var request SubmitRequest
	if err := cont.Bind(&request); err != nil {
		return cont.JSON(http.StatusBadRequest, map[string]string{"error": "Неправильный формат запроса"})
	}

	sessionID, _ := cont.Get("session_id").(uint)

	sub := &domain.Submission{
		TaskID:    request.TaskID,
		SessionID: sessionID,
		Code:      request.Code,
		Status:    domain.StatusPending,
	}

	if err := handl.subRepo.Create(sub); err != nil {
		return cont.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	return cont.JSON(http.StatusOK, map[string]interface{}{
		"id":     sub.ID,
		"status": sub.Status,
	})
}
