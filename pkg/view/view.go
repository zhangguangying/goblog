package view

import (
	"goblog/pkg/logger"
	"goblog/pkg/route"
	"html/template"
	"io"
	"path/filepath"
	"strings"
)

func Render(w io.Writer, name string, data interface{}) {
	viewDir := "resources/views/"

	name = strings.Replace(name, ".", "/", -1)

	files, err := filepath.Glob(viewDir + "/layouts/*.html")
	logger.LogError(err)

	newFiles := append(files, viewDir+name+".html")

	tpl, err := template.New(name + ".html").
		Funcs(template.FuncMap{
			"RouteName2URL": route.Name2URL,
		}).
		ParseFiles(newFiles...)
	logger.LogError(err)

	err = tpl.ExecuteTemplate(w, "app", data)
	logger.LogError(err)
}
