package handler

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/Rodco-off/checkCode/repository"
	"github.com/labstack/echo"
)

type TaskHandler struct {
	repo *repository.TaskRepository
}

func NewTaskHandler(repo *repository.TaskRepository) *TaskHandler {
	return &TaskHandler{repo: repo}
}

func (handl *TaskHandler) GetAll(cont echo.Context) error {
	tasks, err := handl.repo.GetAll()
	if err != nil {
		return cont.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	return cont.JSON(http.StatusOK, tasks)
}

func (handl *TaskHandler) GetByID(cont echo.Context) error {
	id, err := strconv.ParseUint(cont.Param("id"), 10, 64)
	if err != nil {
		return cont.JSON(http.StatusBadRequest, map[string]string{"error": "Неправильный ID"})
	}

	task, err := handl.repo.GetByID(uint(id))
	if errors.Is(err, repository.ErrTaskNotFound) {
		cont.JSON(http.StatusNotFound, map[string]string{"error": "Не найдена задача с таким ID"})
	}

	if err != nil {
		cont.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	return cont.JSON(http.StatusOK, task)

}
