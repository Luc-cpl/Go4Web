package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {
	valid := false
	goPath := os.Getenv("GOPATH")
	scanner := bufio.NewScanner(os.Stdin)
	var value string
	var folder string
	for valid == false {
		fmt.Println("Set your path/application name:")
		for scanner.Scan() {
			folder = scanner.Text()
			break
		}
		fmt.Println()
		fmt.Println("The aplication folder '" + goPath + "/src/" + folder + "' is correct?")
		for scanner.Scan() {
			value = scanner.Text()
			break
		}
		fmt.Println()
		if value == "yes" {
			valid = true
		} else if value == "y" {
			valid = true
		}
	}
	path := goPath + "/src/" + folder + "/"

	fmt.Println("Copying files to new aplication folder...")

	os.Mkdir(path, os.FileMode(0775))
	s := makeFile("Go4Web-basefiles/main.go", folder)
	save(s, path, "main", ".go")

	os.Mkdir(path+"app", os.FileMode(0775))
	s = makeFile("Go4Web-basefiles/app/routes.go", folder)
	save(s, path+"app/", "routes", ".go")

	os.Mkdir(path+"app/controller", os.FileMode(0775))
	s = makeFile("Go4Web-basefiles/app/controller/controller.go", folder)
	save(s, path+"app/controller/", "controller", ".go")
	s = makeFile("Go4Web-basefiles/app/controller/userController.go", folder)
	save(s, path+"app/controller/", "userController", ".go")
	s = makeFile("Go4Web-basefiles/app/controller/viewController.go", folder)
	save(s, path+"app/controller/", "viewController", ".go")

	os.Mkdir(path+"app/model", os.FileMode(0775))

	os.Mkdir(path+"app/model/userData", os.FileMode(0775))
	s = makeFile("Go4Web-basefiles/app/model/userData/userData.go", folder)
	save(s, path+"app/model/userData/", "userData", ".go")

	os.Mkdir(path+"app/model/database", os.FileMode(0775))
	s = makeFile("Go4Web-basefiles/app/model/database/database.go", folder)
	save(s, path+"app/model/database/", "database", ".go")

	os.Mkdir(path+"public", os.FileMode(0775))

	os.Mkdir(path+"public/css", os.FileMode(0775))
	s = makeFile("Go4Web-basefiles/public/css/style.css", folder)
	save(s, path+"public/css/", "style", ".css")

	os.Mkdir(path+"public/js", os.FileMode(0775))

	os.Mkdir(path+"render", os.FileMode(0775))
	s = makeFile("Go4Web-basefiles/render/render.go", folder)
	save(s, path+"render/", "render", ".go")

	os.Mkdir(path+"view", os.FileMode(0775))
	s = makeFile("Go4Web-basefiles/view/home.tmpl", folder)
	save(s, path+"view/", "home", ".tmpl")
	s = makeFile("Go4Web-basefiles/view/layout.tmpl", folder)
	save(s, path+"view/", "layout", ".tmpl")
	s = makeFile("Go4Web-basefiles/view/login.tmpl", folder)
	save(s, path+"view/", "login", ".tmpl")

	fmt.Println()
	fmt.Println("You are read to go!")
	fmt.Println()

}

func save(s string, caminho string, nomeArquivo string, tipoArquivo string) {
	l := caminho + nomeArquivo + tipoArquivo
	f, _ := os.Create(l)
	w := bufio.NewWriter(f)
	w.WriteString(s)
	w.Flush()
	defer f.Close()
}

func makeFile(file string, name string) (s string) {
	buf := bytes.NewBuffer(nil)

	f, _ := os.Open(file) // Error handling elided for brevity.
	io.Copy(buf, f)       // Error handling elided for brevity.
	f.Close()

	s = string(buf.Bytes())

	s = strings.Replace(s, "myprojects/Go4Web", name, -1)
	return
}
