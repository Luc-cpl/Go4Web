package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
)

func main() {

	searchDir := "Go4Web-basefiles"
	fileList := []string{}

	valid := false
	goPath := strings.Replace(os.Getenv("GOPATH"), `\`, `/`, -1)
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

	err := filepath.Walk(searchDir, func(path string, f os.FileInfo, err error) error {
		fileList = append(fileList, path)
		return nil
	})

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("Copying files to new aplication folder...")
	newPath := goPath + "/src/" + folder

	for _, file := range fileList {
		file = strings.Replace(file, `\`, `/`, -1)
		newFilePath := newPath + strings.Replace(file, searchDir, "", 1)
		if strings.ContainsAny(file, ".") == false {
			fmt.Println("folder: " + newFilePath)
			err := os.MkdirAll(newFilePath, os.FileMode(0775))
			if err != nil {
				fmt.Println(err)
				break
			}
		} else {
			s := makeFile(file, folder)
			fileName := filepath.Base(file)
			newFilePath = newFilePath[0 : len(newFilePath)-len(fileName)]
			ext := filepath.Ext(file)
			name := fileName[0 : len(fileName)-len(ext)]
			save(s, newFilePath, name, ext)
		}
	}

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

	f, _ := os.Open(file)
	io.Copy(buf, f)
	f.Close()

	s = string(buf.Bytes())

	s = strings.Replace(s, "github.com/Luc-cpl/Go4Web/Go4Web-basefiles", name, -1)
	return
}
