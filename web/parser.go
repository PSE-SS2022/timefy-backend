package web

import (
	"html/template"
)

func ParseFiles(filenames ...string) (*template.Template, error) {
	return template.ParseFiles(filenames...)
}
