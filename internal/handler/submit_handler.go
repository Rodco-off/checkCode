package handler

import (
	"net/http"

	"github.com/Rodco-off/checkCode/internal/domain"
	"github.com/Rodco-off/checkCode/internal/repository"
	"github.com/Rodco-off/checkCode/internal/service"
	"github.com/labstack/echo"
)

type SubmitHandler struct {
	subRepo *repository.SubmissionRepository
	checker *service.Checker
}

func NewSubmissionHandler(subRepo *repository.SubmissionRepository, checker *service.Checker) *SubmitHandler {
	return &SubmitHandler{subRepo: subRepo, checker: checker}
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

	if err := handl.checker.Check(cont.Request().Context(), sub); err != nil {
		return cont.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	return cont.JSON(http.StatusOK, map[string]any{
		"id":            sub.ID,
		"status":        sub.Status,
		"output":        sub.Output,
		"expected":      sub.Expected,
		"error_message": sub.ErrorMessage,
	})
}
