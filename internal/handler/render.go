package handler

import (
	"html/template"
	"io"

	"github.com/labstack/echo"
)

type TemplateRenderer struct {
	templates *template.Template
}

func NewTemplateRenderer(pattern string) *TemplateRenderer {
	return &TemplateRenderer{
		template.Must(template.ParseGlob(pattern)),
	}
}

func (templ *TemplateRenderer) Render(write io.Writer, name string, data interface{}, cont echo.Context) error {
	return templ.templates.ExecuteTemplate(write, name, data)
}
