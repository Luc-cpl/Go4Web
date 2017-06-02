package controller

import (
	"errors"
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

//password and user-id are forbiden in all mettods

var splitVar1 = "/"
var splitVar2 = "&"

//DatabaseGet  /api/db/get/table/allUsers/query1/query2
func DatabaseGet(w http.ResponseWriter, r *http.Request) {
	urlValues := mux.Vars(r)["rest"]

	err := contentCheck(urlValues)
	if err != nil {
		json.NewEncoder(w).Encode(err)
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

		auth, err := authCheck(userAuth, table)

		if err != nil {
			json.NewEncoder(w).Encode("error requesting authorization")
			return
		}

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

//DatabasePost need keys passed on GET mettod and values on POST: /api/db/post/table/key1&key2&key3
//All posts are bound to a user-id (you can use an annonimous id = 0 if needed)
func DatabasePost(w http.ResponseWriter, r *http.Request) {
	urlValues := mux.Vars(r)["rest"]

	err := contentCheck(urlValues)
	if err != nil {
		json.NewEncoder(w).Encode(err)
		return
	}

	if strings.Count(urlValues, "/") != 1 {
		json.NewEncoder(w).Encode("The path don't matches with necessary fields")
		return
	}
	s := strings.Split(urlValues, splitVar1)
	table := s[0]

	userID := userData.GetUserID(r)

	query1 := strings.Split(s[1], splitVar2)
	newQuery1 := make([]string, len(query1)+1)
	newQuery1[0] = "user-id"
	for i := 1; i < len(newQuery1); i++ {
		newQuery1[i] = query1[i-1]
	}
	query1 = newQuery1

	query2 := make([]string, len(query1))
	query2[0] = userID
	for i := 1; i < len(query1); i++ {
		query2[i] = r.PostFormValue(query1[i])
	}

	userAuth, _ := strconv.Atoi(userData.GetUserAuth(r))

	auth, err := authCheck(userAuth, table)

	if err != nil {
		json.NewEncoder(w).Encode("error requesting authorization")
		return
	}

	if auth.Create == true {
		err = database.DB.Insert(table, query1, query2)
		if err != nil {
			json.NewEncoder(w).Encode("err")
			return
		}
		json.NewEncoder(w).Encode("data included to database")
	} else {
		json.NewEncoder(w).Encode("You don't have autorization for this request")
		return
	}
}

func DatabaseUpdate(w http.ResponseWriter, r *http.Request) {

}

func DatabaseDelete(w http.ResponseWriter, r *http.Request) {

}

type Authorization struct {
	Get       bool
	GetAll    bool
	Create    bool
	CreateAll bool
	Change    bool
	ChangeAll bool
	Delete    bool
	DeleteAll bool
}

func AuthCheck(userAuth int, table string, methods ...string) (auth Authorization, err error) {
	if userAuth == 0 {
		auth.Get = true
		auth.GetAll = true
		auth.Create = true
		auth.CreateAll = true
		auth.Change = true
		auth.ChangeAll = true
		auth.Delete = true
		auth.DeleteAll = true
		return
	}

	checkQuery1 := []string{"table-name", "=", table}
	for _, method := range methods {
		checkQuery2 := []string{}
		if method == "GET" {
			checkQuery2 = []string{
				"auth-level",
				"regressive-auth-level",
				"allcontent-auth",
				"regressive-allcontent-auth"}
		} else if method == "POST" {
			checkQuery2 = []string{
				"auth-level",
				"regressive-auth-level",
				"create-auth",
				"regressive-create-auth"}
		} else if method == "UPDATE" {
			checkQuery2 = []string{
				"change-auth",
				"regressive-change-auth",
				"allcontent-change-auth",
				"regressive-allcontent-change-auth"}
		} else if method == "DELETE" {
			checkQuery2 = []string{
				"delete-auth",
				"regressive-delete-auth",
				"allcontent-delete-auth",
				"regressive-allcontent-delete-auth"}
		}

		check, err := database.DB.Get("db-auth", checkQuery1, checkQuery2)
		if err != nil {
			return auth, err
		}
		authLevel := 999
		regLevel := 999
		allLevel := 999
		regAllLevel := 999
		for _, y := range check {
			s := make(map[int]int)
			n := 0
			for _, value := range y {
				s[n], _ = strconv.Atoi(value)
				n++
			}
			if s[0] < authLevel && s[0] > userAuth {
				authLevel = s[0]
				regLevel = s[1]
				allLevel = s[2]
				regAllLevel = s[3]
			} else if s[0] == userAuth {
				authLevel = s[0]
				regLevel = s[1]
				allLevel = s[2]
				regAllLevel = s[3]
				break
			}
		}
		if authLevel >= userAuth && regLevel <= userAuth {
			if method == "GET" {
				auth.GetAll = true
			} else if method == "POST" {
				auth.CreateAll = true
			} else if method == "UPDATE" {
				auth.ChangeAll = true
			} else if method == "DELETE" {
				auth.DeleteAll = true
			}
			if allLevel >= userAuth && regAllLevel <= userAuth {

			}

		}

	}

	return
}

func contentCheck(urlValues string) (err error) {
	if strings.Contains(urlValues, "user-id") {
		err = errors.New("Unauthorized request")
		return
	} else if strings.Contains(urlValues, "password") {
		err = errors.New("Unauthorized request")
		return
	} else if strings.Contains(urlValues, ";") {
		err = errors.New("Unauthorized request")
		return
	}
	return
}
