package views

import (
	"embed"
	"html/template"
	"io"
	"log/slog"

	"github.com/labstack/echo/v4"
)

//go:embed templates/*
var tmplFS embed.FS

type TemplateManager struct {
	templates *template.Template
}

func NewTemplateManager() (*TemplateManager, error) {
	templates, err := template.New("").ParseFS(tmplFS, "templates/*.html")
	if err != nil {
		return nil, err
	}
	return &TemplateManager{
		templates: templates,
	}, nil
}

func (t *TemplateManager) Render(
	w io.Writer,
	name string,
	data interface{},
	c echo.Context,
) error {
	// TODO: fix this up, it doesn't look right
	slog.Info("Rendering template: " + name)
	tmpl := template.Must(t.templates.Clone())
	tmpl = template.Must(tmpl.ParseFS(tmplFS, "templates/"+name))
	return tmpl.ExecuteTemplate(w, name, data)
}
