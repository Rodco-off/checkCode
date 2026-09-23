package handler

import (
	"net/http"

	"github.com/google/uuid"
	"github.com/labstack/echo"
)

func SessionMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(cont echo.Context) error {
		cookie, err := cont.Cookie("session_id")
		if err != nil {
			sessionID := uuid.New().String()
			cont.SetCookie(&http.Cookie{
				Name: "session_id", Value: sessionID,
				MaxAge: 86400 * 30, HttpOnly: true, Path: "/",
			})
			cont.Set("session_id", sessionID)
		} else {
			cont.Set("session_id", cookie.Value)
		}
		return next(cont)
	}
}
