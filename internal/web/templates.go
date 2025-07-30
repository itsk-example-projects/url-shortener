package web

import (
	"html/template"
	"path/filepath"
)

func RegisterTemplates() (*template.Template, error) {
	return template.ParseGlob(filepath.Join("./internal/web/templates", "*.html"))
}
