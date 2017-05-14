package controller

import (
	"encoding/json"
	"net/http"

	"github.com/Luc-cpl/Go4Web/Go4Web-basefiles/app/model/userData"
)

//O Json "logged" indica que o cliente fez login, se falhar fazer verificação pela existência do cookie
//A função do model userData "GetUserID"

type Response struct {
	Login bool   `json:"login"`
	Err   string `json:"err"`
}

var response Response
var user userData.User

//UserLogin faz o login do usuário depois de passar as informações "email" e "password" por uma requisição POST
func UserLogin(w http.ResponseWriter, r *http.Request) {

	user.Email = r.PostFormValue("email")
	user.Password = r.PostFormValue("password")

	userID, login := userData.Login(user)
	value := map[string]string{"userId": userID}
	if login == false {
		response.Login = false
		json.NewEncoder(w).Encode(response)
	} else if encoded, err := userData.Cookie.Encode("session", value); err == nil {
		cookie := &http.Cookie{
			Name:   "session",
			Value:  encoded,
			Path:   "/",
			MaxAge: 1800, //duração do cookie
		}
		http.SetCookie(w, cookie)
		response.Login = true
		json.NewEncoder(w).Encode(response)
	}
}

//NewUser cria um novo usuário no banco de dados
func NewUser(w http.ResponseWriter, r *http.Request) {

	user.Email = r.PostFormValue("email")
	user.EmailConf = r.PostFormValue("emailConf")
	user.Password = r.PostFormValue("password")
	user.PasswordConf = r.PostFormValue("passwordConf")
	user.Name = r.PostFormValue("name")

	userID, _, err := userData.NewUser(user)

	value := map[string]string{"userId": userID}
	if err != nil {
		response.Login = false
		response.Err = err.Error()
		json.NewEncoder(w).Encode(response)
	} else if encoded, err := userData.Cookie.Encode("session", value); err == nil {
		cookie := &http.Cookie{
			Name:   "session",
			Value:  encoded,
			Path:   "/",
			MaxAge: 1800, //duração do cookie
		}
		http.SetCookie(w, cookie)
		response.Login = true
		json.NewEncoder(w).Encode(response)
	}
}

//UserUpdate updates a user account on database
func UpdateUser(w http.ResponseWriter, r *http.Request) {

}
