package controller

import (
	"encoding/json"
	"myprojects/Go4Web/app/model/userData"
	"net/http"
)

//O Json "logged" indica que o cliente fez login, se falhar fazer verificação pela existência do cookie
//A função do model userData "GetUserID"

//UserLogin faz o login do usuário depois de passar as informações "email" e "password" por uma requisição POST
func UserLogin(w http.ResponseWriter, r *http.Request) {
	userID, login := userData.Login(w, r)
	value := map[string]string{"userId": userID}
	if login == false {
		json.NewEncoder(w).Encode("email or password invalid")
	} else if encoded, err := userData.Cookie.Encode("session", value); err == nil {
		cookie := &http.Cookie{
			Name:   "session",
			Value:  encoded,
			Path:   "/",
			MaxAge: 1800, //duração do cookie
		}
		http.SetCookie(w, cookie)
		json.NewEncoder(w).Encode("logged")
	}
}

//NewUser cria um novo usuário no banco de dados
func NewUser(w http.ResponseWriter, r *http.Request) {
	userID, login, err := userData.NewUser(r)
	value := map[string]string{"userId": userID}
	if login == false {
		json.NewEncoder(w).Encode(err)
	} else if encoded, err := userData.Cookie.Encode("session", value); err == nil {
		cookie := &http.Cookie{
			Name:   "session",
			Value:  encoded,
			Path:   "/",
			MaxAge: 1800, //duração do cookie
		}
		http.SetCookie(w, cookie)
		json.NewEncoder(w).Encode("logged")
	}
}
