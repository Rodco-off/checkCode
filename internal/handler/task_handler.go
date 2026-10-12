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

func (handl *TaskHandler) TasksPage(cont echo.Context) error {
	tasks, err := handl.repo.GetAll()
	if err != nil {
		return cont.String(http.StatusInternalServerError, "error")
	}

	return cont.Render(http.StatusOK, "tasks.html", map[string]interface{}{
		"Tasks": tasks,
	})
}

func (handl *TaskHandler) TaskPage(cont echo.Context) error {
	id, err := strconv.ParseUint(cont.Param("id"), 10, 64)
	if err != nil {
		return cont.String(http.StatusBadRequest, "Неправильный id")
	}

	task, err := handl.repo.GetByID(uint(id))
	if errors.Is(err, repository.ErrTaskNotFound) {
		return cont.String(http.StatusNotFound, "Задача не найдена")
	}
	if err != nil {
		return cont.String(http.StatusBadRequest, "Ошибка")
	}

	return cont.Render(http.StatusOK, "task.html", map[string]interface{}{
		"Task": task,
	})
}
