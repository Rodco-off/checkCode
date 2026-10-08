package handler

import (
	"net/http"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func NewRouter(taskHandler TaskHandler, submitHandler SubmitHandler) *echo.Echo {
	e := echo.New()

	e.Renderer = NewTemplateRenderer("templates/*.html")
	e.Static("/static", "static")

	e.Use(middleware.Recover())
	e.Use(middleware.Logger())
	e.Use(SessionMiddleware)

	api := e.Group("/api")
	{
		api.GET("/tasks", taskHandler.GetAll)
		api.GET("/tasks/:id", taskHandler.GetByID)
		api.POST("/submit", submitHandler.Submit)
	}

	e.GET("/health", func(cont echo.Context) error {
		return cont.JSON(http.StatusOK, map[string]string{"status": "ok"})
	})

	e.GET("/", taskHandler.TasksPage)
	e.GET("/tasks/:id", taskHandler.TaskPage)

	return e
}
