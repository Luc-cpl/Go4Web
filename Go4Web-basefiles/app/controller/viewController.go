package controller

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"

	"github.com/Luc-cpl/Go4Web/Go4Web-basefiles/app/model/userData"

	"github.com/Luc-cpl/Go4Web/Go4Web-basefiles/render"
	"github.com/gorilla/mux"
)

//The URL to redirect if the JSON redirect is true
var redirectURL = "/"

type viewMap struct {
	URL      string `json:"url"`      //the access url
	HTML     string `json:"html"`     //the html to pass in view folder
	Template string `json:"template"` //the html template to pass in view folder
	CSS      string `json:"css"`      //the css to pass in view folder
	JS       string `json:"js"`       //the javascript to pass in view folder
	Auth     bool   `json:"auth"`     //if login is necessary
	Redirect bool   `json:"redirect"` //to redirect if is logged
}

//Views controll all the url views in webapp
func Views(w http.ResponseWriter, r *http.Request) {
	url := mux.Vars(r)["rest"]

	raw, err := ioutil.ReadFile("./views/viewmap.json")

	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	var decodedJSON []viewMap
	var view viewMap

	err = json.Unmarshal(raw, &decodedJSON)
	if err != nil {
		fmt.Println(err.Error())
	}
	exist := false
	auth := false

	for _, element := range decodedJSON {
		if element.URL == url {
			exist = true

			if element.Auth == false {

				if element.Redirect == true && userData.GetUserID(r) != "0" {
					if ("/" + url) != redirectURL {
						http.Redirect(w, r, redirectURL, http.StatusSeeOther)
						break
					} else {
						auth = false
					}
				} else {
					auth = true
					view = element
					break
				}

			} else if userData.GetUserID(r) != "0" {
				auth = true
				view = element
				break
			}

		}
	}
	if auth == true && exist == true {
		render.Render(w, nil, view.Template, view.HTML, view.CSS, view.JS)
	} else if exist == true {
		http.Redirect(w, r, "/", http.StatusSeeOther)
	} else {
		if strings.ContainsAny(url, ".") {
			http.ServeFile(w, r, "./public/"+url)
		} else {
			render.Render(w, nil, "template.html", "error404.html")
		}
	}

}
