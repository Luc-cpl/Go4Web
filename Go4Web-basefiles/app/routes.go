package routes

import (
	"net/http"

	"github.com/Luc-cpl/Go4Web/Go4Web-basefiles/app/controller"

	"github.com/gorilla/mux"
)

func NewRouter() *mux.Router {

	var routes = Routes{
		//API
		Route{"UserLogin", "POST", "/api/users/login", controller.UserLogin},
		Route{"NewUser", "POST", "/api/users/new", controller.NewUser},
		Route{"UpdateUser", "PUT", "/api/users/update", controller.UpdateUser},
		Route{"ExemploGet", "GET", "/exemploGet/{id}", controller.ExemploGet},

		//views manager
		Route{"Index", "GET", "/", controller.Views},
		Route{"Views", "GET", "/{rest:.*}", controller.Views},
	}

	router := mux.NewRouter().StrictSlash(true)

	for _, route := range routes {
		router.Methods(route.Method).Path(route.Pattern).Name(route.Name).Handler(route.HandlerFunc)
	}

	return router

}

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

type Routes []Route
