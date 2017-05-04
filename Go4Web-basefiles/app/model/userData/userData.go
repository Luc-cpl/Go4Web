package userData

import (
	"encoding/hex"
	"myprojects/Go4Web/app/model/database"
	"net/http"

	"github.com/gorilla/securecookie"

	"crypto"

	"fmt"
)

//Cookie é uma variavel para utilização e criação de cookies segundo uma chae aleatória criada
var Cookie = securecookie.New(securecookie.GenerateRandomKey(64), securecookie.GenerateRandomKey(32))

//securekey é uma chave de segurança para adicionar a senha do usuário na criação da hash, dificultando a quebra de senha por orça bruta
var secureKey = "j2jfU3rj9f"

//table é a variavel que indica a tabela de usuários no banco de dados
var table = "users"

//Login retorna se o login é verdadeiro ou falso e o ID no banco de dados
func Login(w http.ResponseWriter, r *http.Request) (userID string, login bool) {

	email := r.PostFormValue("email")
	password := r.PostFormValue("password")

	hash := crypto.SHA256.New()
	hash.Write([]byte(password + secureKey))
	password = hex.EncodeToString(hash.Sum(nil))

	fmt.Println(password)

	id := []string{"email", email, "senha", password}
	campos := []string{"codigo"}

	y, _ := database.DB.Get(table, id, campos)

	if y != nil {
		login = true
		userID = y[0]["codigo"]
	}
	return
}

//GetUserID retorna o ID do usuário a partir do Cookie recebido na reuisição
func GetUserID(r *http.Request) (userID string) {
	if cookie, err := r.Cookie("session"); err == nil {
		value := make(map[string]string)
		if err = Cookie.Decode("session", cookie.Value, &value); err == nil {
			userID = value["userId"]
		}
	}
	return
}

//NewUser cria o usuario no banco de dados e retorna sua cituação de Login (ou retorna erro se o usuário já existe)
func NewUser(r *http.Request) (userID string, login bool, erro string) {
	email := r.PostFormValue("email")
	emailConf := r.PostFormValue("emailConf")
	password := r.PostFormValue("password")
	passwordConf := r.PostFormValue("passwordConf")
	name := r.PostFormValue("name")

	if email == emailConf {
		if password == passwordConf {
			col := []string{"nome", "email", "senha"}

			hash := crypto.SHA256.New()
			hash.Write([]byte(password + secureKey))
			password = hex.EncodeToString(hash.Sum(nil))

			val := []string{name, email, password}

			err := database.DB.Insert(table, col, val)
			if err != nil {
				erro = err.Error()
				login = false
				return
			} else {
				id := []string{"email", email}
				campos := []string{"codigo"}
				y, _ := database.DB.Get(table, id, campos)

				if y != nil {
					login = true
					userID = y[0]["codigo"]
				}
			}

		} else {
			erro = "senha"
		}
	} else {
		erro = "email"
	}
	return
}
