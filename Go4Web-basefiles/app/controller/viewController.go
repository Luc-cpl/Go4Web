package controller

import (
	"myprojects/Go4Web/render"
	"net/http"
)

// Index é a tela de entrada do webApp
func Index(w http.ResponseWriter, r *http.Request) {
	render.Render(w, "view/home.tmpl", nil)
}
