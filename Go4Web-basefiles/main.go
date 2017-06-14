package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/Luc-cpl/Go4Web/Go4Web-basefiles/app"
	mgoS "github.com/Luc-cpl/mgoSimpleCRUD"
	mgo "gopkg.in/mgo.v2"
)

var serverPort = ":3000"

func main() {

	//Database connection info (eddit if necessary)
	maxWait := time.Duration(5 * time.Second)
	local, err := mgo.DialWithTimeout("localhost:27017", maxWait)
	if err == nil {
		fmt.Println("Connected to MongoDB.")
		defer local.Close()
		mgoS.DB.Session = local.Clone()
		mgoS.DB.Database = "teste"
		mgoS.DB.UserIdentityValue = "email"
		mgoS.DB.AuthMap, err = mgoS.LoadAuthMap("./authMap.json")
		if err != nil {
			fmt.Println(err)
			return
		}
	} else {
		fmt.Println("Unable to connect to local mongo instance!")
	}

	router := routes.NewRouter()

	fmt.Println("started server at port", serverPort)
	http.ListenAndServe(serverPort, router)
}
