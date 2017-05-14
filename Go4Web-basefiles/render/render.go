package render

import (
	"fmt"
	"html/template"
	"net/http"
)

func Render(w http.ResponseWriter, data interface{}, files ...string) {
	var err error
	tmpl := template.New("")

	if len(files) == 2 {
		tmpl, err = template.ParseFiles(files[0], files[1])
	}
	if len(files) == 3 {
		tmpl, err = template.ParseFiles(files[0], files[1], files[2])
	}
	if len(files) == 4 {
		tmpl, err = template.ParseFiles(files[0], files[1], files[2], files[3])
	}

	if err != nil {
		fmt.Println(err)
		return
	}

	if err = tmpl.Execute(w, data); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
