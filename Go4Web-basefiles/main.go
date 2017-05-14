package main

import (
	"log"
	"github.com/Luc-cpl/Go4Web/Go4Web-basefiles/app"
	"github.com/Luc-cpl/Go4Web/Go4Web-basefiles/app/model/database"
	"net/http"
)

func main() {

	db, err := database.NewOpen("root:@/db_teste")
	if err != nil {
		log.Println(err)
	}
	defer db.Close()
	database.DB = db

	router := routes.NewRouter()

	http.ListenAndServe(":8080", router)
}
