package view

import (
	"goblog/pkg/logger"
	"goblog/pkg/route"
	"html/template"
	"io"
	"path/filepath"
	"strings"
)

type D map[string]interface{}

func Render(w io.Writer, data interface{}, tplFiles ...string) {
	RenderTemplate(w, "app", data, tplFiles...)
}

func RenderSimple(w io.Writer, data interface{}, tplFiles ...string) {
	RenderTemplate(w, "simple", data, tplFiles...)
}

func RenderTemplate(w io.Writer, name string, data interface{}, tplFiles ...string) {
	viewDir := "resources/views/"

	for i, f := range tplFiles {
		tplFiles[i] = viewDir + strings.Replace(f, ".", "/", -1) + ".html"
	}

	layoutFiles, err := filepath.Glob(viewDir + "/layouts/*.html")
	logger.LogError(err)

	allFiles := append(layoutFiles, tplFiles...)

	tpl, err := template.New("").
		Funcs(template.FuncMap{
			"RouteName2URL": route.Name2URL,
		}).
		ParseFiles(allFiles...)
	logger.LogError(err)

	err = tpl.ExecuteTemplate(w, name, data)
	logger.LogError(err)
}
