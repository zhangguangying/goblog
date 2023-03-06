package view

import (
	"fmt"
	"goblog/pkg/logger"
	"goblog/pkg/route"
	"html/template"
	"io"
	"path/filepath"
	"strings"
)

func Render(w io.Writer, data interface{}, tplFiles ...string) {
	viewDir := "resources/views/"

	for i, f := range tplFiles {
		tplFiles[i] = viewDir + strings.Replace(f, ".", "/", -1) + ".html"
	}

	layoutFiles, err := filepath.Glob(viewDir + "/layouts/*.html")
	logger.LogError(err)

	allFiles := append(layoutFiles, tplFiles...)
	fmt.Println(allFiles)

	tpl, err := template.New("").
		Funcs(template.FuncMap{
			"RouteName2URL": route.Name2URL,
		}).
		ParseFiles(allFiles...)
	logger.LogError(err)

	err = tpl.ExecuteTemplate(w, "app", data)
	logger.LogError(err)
}
