package render

import (
	"html/template"
	"io/ioutil"
	"net/http"
	"path/filepath"
)

var path = "views/files-min/"

func Render(w http.ResponseWriter, data interface{}, files ...string) {
	var err error
	tmpl := template.New("")

	var html string
	for _, element := range files {
		file, _ := ioutil.ReadFile(path + element)
		s := string(file)

		if filepath.Ext(path+element) == ".css" {
			s = `{{define "css"}}` + s + `{{end}}`
		} else if filepath.Ext(path+element) == ".js" {
			s = `{{define "js"}}` + s + `{{end}}`
		}
		html = html + s
	}
	tmpl, _ = tmpl.Parse(html)

	if err = tmpl.Execute(w, data); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

}
