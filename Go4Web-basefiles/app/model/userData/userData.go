package userData

import (
	"encoding/hex"
	"net/http"

	"github.com/Luc-cpl/Go4Web/Go4Web-basefiles/app/model/database"

	"github.com/gorilla/securecookie"

	"crypto"
	"errors"
)

//Cookie é uma variavel para utilização e criação de cookies segundo uma chae aleatória criada
var Cookie = securecookie.New(securecookie.GenerateRandomKey(64), securecookie.GenerateRandomKey(32))

//securekey é uma chave de segurança para adicionar a senha do usuário na criação da hash, dificultando a quebra de senha por orça bruta
var secureKey = "j2jfU3rj9f"

//table é a variavel que indica a tabela de usuários no banco de dados
var table = "users"

//User pass user information to model
type User struct {
	Name         string
	Email        string
	EmailConf    string
	Password     string
	PasswordConf string
}

//Login retorna se o login é verdadeiro ou falso e o ID no banco de dados
func Login(user User) (userID string, auth string, login bool) {

	hash := crypto.SHA256.New()
	hash.Write([]byte(user.Password + secureKey))
	user.Password = hex.EncodeToString(hash.Sum(nil))

	id := []string{"email", "=", user.Email, "password", "=", user.Password}
	campos := []string{"user-id", "auth-level"}

	y, _ := database.DB.Get(table, id, campos)

	if y[0]["user-id"] != "" {
		login = true
		userID = y[0]["user-id"]
		auth = y[0]["auth-level"]
	}
	return
}

//GetUserID retorna o ID do usuário a partir do Cookie recebido na reuisição
func GetUserID(r *http.Request) (userID string) {
	if cookie, err := r.Cookie("session"); err == nil {
		value := make(map[string]string)
		if err = Cookie.Decode("session", cookie.Value, &value); err == nil {
			return value["userId"]
		}
	}
	return
}

func GetUserAuth(r *http.Request) (auth string) {

	if cookie, err := r.Cookie("session"); err == nil {
		value := make(map[string]string)
		if err = Cookie.Decode("session", cookie.Value, &value); err == nil {
			return value["userAuth"]
		}
	}
	return "999"
}

//NewUser cria o usuario no banco de dados e retorna sua cituação de Login (ou retorna erro se o usuário já existe)
func NewUser(user User) (userID string, auth string, err error) {

	if user.Email == user.EmailConf {
		if user.Password == user.PasswordConf {
			col := []string{"name", "email", "password"}

			hash := crypto.SHA256.New()
			hash.Write([]byte(user.Password + secureKey))
			password := hex.EncodeToString(hash.Sum(nil))

			val := []string{user.Name, user.Email, password}

			err := database.DB.Insert(table, col, val)

			if err != nil {
				err = errors.New("User already exist")
				return "", "", err
			}

			id := []string{"email", "=", user.Email}
			campos := []string{"user-id"}
			y, _ := database.DB.Get(table, id, campos)

			if y != nil {
				auth = y[0]["auth-level"]
				userID = y[0]["user-id"]
				return userID, auth, nil
			}

		} else {
			err = errors.New("Check your password")
		}
	} else {
		err = errors.New("Check your email")
	}
	return
}
