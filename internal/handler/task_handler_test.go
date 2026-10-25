package handler

import (
	"errors"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo"
	"github.com/stretchr/testify/assert"

	"github.com/Rodco-off/checkCode/internal/domain"
	"github.com/Rodco-off/checkCode/internal/repository"
)

type MockTaskRepo struct{}

func (m *MockTaskRepo) GetAll(difficulty string) ([]domain.Task, error) {
	if difficulty == "error" {
		return nil, errors.New("fake database error")
	}
	return []domain.Task{
		{TaskID: 1, Title: "Тестовая задача", Difficulty: "easy"},
		{TaskID: 2, Title: "Вторая задача", Difficulty: "medium"},
	}, nil
}

func (m *MockTaskRepo) GetByID(id uint) (*domain.Task, error) {
	if id == 999 {
		return nil, repository.ErrTaskNotFound
	}
	return &domain.Task{TaskID: id, Title: "Задача по ID", Difficulty: "easy"}, nil
}

type MockRenderer struct{}

func (m *MockRenderer) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	_, err := w.Write([]byte("<html>Фейковый HTML с задачей</html>"))
	return err
}

func TestTasksPage_Success(t *testing.T) {
	e := echo.New()
	e.Renderer = &MockRenderer{}

	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	h := NewTaskHandler(&MockTaskRepo{})
	err := h.TasksPage(c)

	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, rec.Code)
	assert.Contains(t, rec.Body.String(), "Фейковый HTML")
}

func TestGetByID_Success(t *testing.T) {
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/api/tasks/1", nil)
	rec := httptest.NewRecorder()
	cont := e.NewContext(req, rec)

	cont.SetParamNames("id")
	cont.SetParamValues("1")

	h := NewTaskHandler(&MockTaskRepo{})
	err := h.GetByID(cont)

	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, rec.Code)
	assert.Contains(t, rec.Body.String(), "Задача по ID")
}

func TestGetByID_NotFound(t *testing.T) {
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/api/tasks/999", nil)
	rec := httptest.NewRecorder()
	cont := e.NewContext(req, rec)

	cont.SetParamNames("id")
	cont.SetParamValues("999")

	h := NewTaskHandler(&MockTaskRepo{})
	err := h.GetByID(cont)

	assert.NoError(t, err)
	assert.Equal(t, http.StatusNotFound, rec.Code)

	assert.Contains(t, rec.Body.String(), "Не найдена задача с таким ID")
}

func TestGetByID_BadRequest(t *testing.T) {
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/api/tasks/abc", nil)
	rec := httptest.NewRecorder()
	cont := e.NewContext(req, rec)

	cont.SetParamNames("id")
	cont.SetParamValues("abc")

	h := NewTaskHandler(&MockTaskRepo{})
	err := h.GetByID(cont)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusBadRequest, rec.Code)
	assert.Contains(t, rec.Body.String(), "Неправильный ID")
}
