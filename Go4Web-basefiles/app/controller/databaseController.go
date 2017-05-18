package controller

import (
	"net/http"
	"strconv"
	"strings"

	"github.com/Luc-cpl/Go4Web/Go4Web-basefiles/app/model/database"
	"github.com/Luc-cpl/Go4Web/Go4Web-basefiles/app/model/userData"

	"encoding/json"

	"github.com/gorilla/mux"
)

//access 0 = root
//access 1 = can take info from users with access equal or greater than 1
//access 2 = can take info from users with access equal or greater than 2
//...

//  table = table name
//	allUsers = (true = all info/ false = user info)
//	query1 = follows databese function
//	query2 = follows databese function

var splitVar1 = "/"
var splitVar2 = "&"

//DatabaseGet  /api/db/get/table/allUsers/query1/query2
func DatabaseGet(w http.ResponseWriter, r *http.Request) {
	urlValues := mux.Vars(r)["rest"]

	if strings.Contains(urlValues, "user-id") {
		json.NewEncoder(w).Encode("Unauthorized request")
		return
	} else if strings.Contains(urlValues, "password") {
		json.NewEncoder(w).Encode("Unauthorized request")
		return
	}

	if strings.Count(urlValues, "/") != 3 {
		json.NewEncoder(w).Encode("The path don't matches with necessary fields")
		return
	}

	s := strings.Split(urlValues, splitVar1)
	table := s[0]
	allUsers := s[1]
	var query1 []string
	if s[2] != "&" {
		query1 = strings.Split(s[2], splitVar2)
	}
	query2 := strings.Split(s[3], splitVar2)

	if allUsers == "false" {
		userID := userData.GetUserID(r)
		newQuery1 := make([]string, len(query1)+3)
		newQuery1[0] = "user-id"
		newQuery1[1] = "="
		newQuery1[2] = userID
		for i := 3; i < len(newQuery1); i++ {
			newQuery1[i] = query1[i-3]
		}
		query1 = newQuery1
	}
	var auth authorization

	if (len(query1)%3) == 0 && len(query1) > 0 {

		userAuth, _ := strconv.Atoi(userData.GetUserAuth(r))

		auth = authCheck(userAuth, table)

		if auth.AllContent == false && query1[0] != "user-id" {
			auth.Auth = false
		}
	} else {
		json.NewEncoder(w).Encode("The request values don´t match with necessary number of values")
		return
	}

	if auth.Auth == true {
		response, err := database.DB.Get(table, query1, query2)
		if err != nil {
			json.NewEncoder(w).Encode(err)
			return
		}

		json.NewEncoder(w).Encode(response)
	} else {
		json.NewEncoder(w).Encode("You don't have autorization for this request")
		return
	}

}

func DatabasePut(w http.ResponseWriter, r *http.Request) {

}

func DatabaseUpdate(w http.ResponseWriter, r *http.Request) {

}

func DatabaseDelete(w http.ResponseWriter, r *http.Request) {

}

type authorization struct {
	Auth       bool
	AllContent bool
	Create     bool
	Change     bool
	Delete     bool
}

func authCheck(userAuth int, table string) (auth authorization) {

	checkQuery1 := []string{"table-name", "=", table}

	checkQuery2 := []string{"auth", "auth-level", "regressive-auth", "allcontent-auth", "create-auth", "change-auth", "delete-auth"}

	check, _ := database.DB.Get("db-auth", checkQuery1, checkQuery2)
	authLevel, _ := strconv.Atoi(check[0]["auth-level"])
	regressiveAuth, _ := strconv.Atoi(check[0]["regressive-auth"])
	AllcontentAuth, _ := strconv.Atoi(check[0]["allcontent-auth"])
	createAuth, _ := strconv.Atoi(check[0]["create-auth"])
	changeAuth, _ := strconv.Atoi(check[0]["change-auth"])
	deleteAuth, _ := strconv.Atoi(check[0]["delete-auth"])

	//check for basic authorization
	if check[0]["auth"] == "0" {
		auth.Auth = true
		auth.AllContent = true
	} else if authLevel >= userAuth && regressiveAuth <= userAuth {
		auth.Auth = true
	} else if userAuth == 0 {
		auth.Auth = true
	} else {
		return
	}

	if AllcontentAuth >= userAuth {
		auth.AllContent = true
	}
	if createAuth >= userAuth {
		auth.Create = true
	}
	if changeAuth >= userAuth {
		auth.Change = true
	}
	if deleteAuth >= userAuth {
		auth.Delete = true
	}
	return
}
