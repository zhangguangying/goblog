package view

import (
	"fmt"
	"github.com/zhangguangying/goblog/pkg/logger"
	"github.com/zhangguangying/goblog/pkg/route"
	"html/template"
	"net/http"
	"path/filepath"
	"strings"
)

type D map[string]interface{}

func RenderSimple(w http.ResponseWriter, data interface{}, tplFiles ...string) {
	RenderTemplate(w, "simple", data, tplFiles...)
}

func Render(w http.ResponseWriter, data interface{}, tplFiles ...string) {
	RenderTemplate(w, "app", data, tplFiles...)
}

func RenderTemplate(w http.ResponseWriter, name string, data interface{}, tplFiles ...string) {
	viewDir := "resources/views/"
	for i, f := range tplFiles {
		tplFiles[i] = viewDir + strings.Replace(f, ".", "/", -1) + ".gohtml"
	}
	files, err := filepath.Glob(viewDir + "layouts/*.gohtml")
	logger.LogError(err)
	files = append(files, tplFiles...)
	fmt.Println(files)
	tmpl, err := template.New("").
		Funcs(template.FuncMap{
			"RouteName2URL": route.Name2URL,
		}).ParseFiles(files...)
	logger.LogError(err)

	tmpl.ExecuteTemplate(w, name, data)
}
