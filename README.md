Go4Web is a full MVC framework designed in GO.
---
Is just a simple start point to build web applications.

***

To use Go4Web framework, is necessary de follow packages:

*  github.com/gorilla/securecookie
```
      $ go get github.com/gorilla/securecookie
```
*  github.com/gorilla/mux
```
      $ go get github.com/gorilla/mux
```

---

To build a new project, just go to Go4Web folder and run the installer.
```
    $ cd GOPATH/src/github.com/Luc-cpl/Go4Web
    $ go run Go4Web.go
```
After start the installer, create your project, inserting the name and folder, after that, confirm the application path.
```
    $ Exemple-folder/Project
    $ yes
```
In this case, the installer will create the project in GOPATH/src/Exemple-folder/Project.

---
To run your application, go to your new application folder and run de main.go to start the server (is necessary restart the server every time that a .go file in the application is changed).
```
    $ cd GOPATH/src/Exemple-folder/Project
    $ go run main.go
```   
---  
To connect your MySQL database, just open the main.go and change "root:@/db_teste" to your MySQL access and database according to go-sql-driver package.
