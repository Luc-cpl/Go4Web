package routes

import (
	"myprojects/Go4Web/app/controller"
	"net/http"

	"github.com/gorilla/mux"
)

func NewRouter() *mux.Router {

	var routes = Routes{
		Route{"Index", "GET", "/", controller.Index},

		Route{"UserLogin", "POST", "/userLogin", controller.UserLogin},
		Route{"NewUser", "POST", "/newUser", controller.NewUser},

		Route{"ExemploGet", "GET", "/exemploGet/{id}", controller.ExemploGet},
		//Route{"User", "POST", "/user", controller.User},
	}

	router := mux.NewRouter().StrictSlash(true)

	for _, route := range routes {
		router.Methods(route.Method).Path(route.Pattern).Name(route.Name).Handler(route.HandlerFunc)
	}
	router.PathPrefix("/").Handler(http.FileServer(http.Dir("./public/")))

	return router

}

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

type Routes []Route
