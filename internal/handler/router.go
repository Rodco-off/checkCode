package handler

import (
	"net/http"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func NewRouter(taskHandler TaskHandler) *echo.Echo {
	e := echo.New()

	e.Use(middleware.Recover())
	e.Use(middleware.Logger())

	api := e.Group("/api")
	{
		api.GET("/tasks", taskHandler.GetAll)
		api.GET("/tasks/:id", taskHandler.GetByID)
	}

	e.GET("/health", func(cont echo.Context) error {
		return cont.JSON(http.StatusOK, map[string]string{"status": "ok"})
	})

	return e
}
