package render

import (
	"fmt"
	"html/template"
	"net/http"
)

func Render(w http.ResponseWriter, filename string, data interface{}) {
	var err error
	tmpl := template.New("")

	if tmpl, err = template.ParseFiles("view/layout.tmpl", filename); err != nil {
		fmt.Println(err)
		return
	}

	if err = tmpl.Execute(w, data); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
