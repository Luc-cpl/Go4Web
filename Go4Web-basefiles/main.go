package main

import (
	"log"
	"myprojects/Go4Web/app"
	"myprojects/Go4Web/app/model/database"
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

	http.ListenAndServe(":7012", router)
}
