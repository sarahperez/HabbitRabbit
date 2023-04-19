//references:
//https://golang.ch/which-golang-router-to-use-for-what-cases/
// https://github.com/gorilla/mux
// https://medium.com/@anshap1719/getting-started-with-angular-and-go-setting-up-a-boilerplate-project-8c273b81aa6
// https://pkg.go.dev/net/http#ServeMux

//after starting the server, open a new terminal and run these comands to test get http requests:
// curl.exe -v -X GET http://localhost:3000/home-page
// curl.exe -v -X GET http://localhost:3000/calendar
//curl.exe -v -X POST http://localhost:3000/home-page -H 'Content-Type: application/json' -d "@userInfo.json"

//curl commands to test login and register
//curl.exe -v -X POST http://localhost:3000/login -H 'Content-Type: application/json' -d "@userInfo.json"
//curl.exe -v -X POST http://localhost:3000/register -H 'Content-Type: application/json' -d "@userInfo.json"

//curl commands to test adding to to do list, completeing on the to do list and getting the status
//curl.exe -v -X POST http://localhost:3000/EditToDo -H 'Content-Type: application/json' -d "@userInfo.json"
//curl.exe -v -X DELETE http://localhost:3000/EditToDo -H 'Content-Type: application/json' -d "@userInfo.json"
//curl.exe -v -X PUT http://localhost:3000/EditToDo -H 'Content-Type: application/json' -d "@userInfo.json"
//curl.exe -v -X POST http://localhost:3000/ToDoStatus -H 'Content-Type: application/json' -d "@userInfo.json"

//curl commands to calendar
//curl.exe -v -X POST http://localhost:3000/EditCal -H 'Content-Type: application/json' -d "@userInfo.json"
//curl.exe -v -X DELETE http://localhost:3000/EditCal -H 'Content-Type: application/json' -d "@userInfo.json"
//curl.exe -v -X POST http://localhost:3000/CalStatus -H 'Content-Type: application/json' -d "@userInfo.json"

//curl commands to friend table
//curl.exe -v -X POST http://localhost:3000/RequestFriend -H 'Content-Type: application/json' -d "@userInfo.json"
//curl.exe -v -X POST http://localhost:3000/AcceptFriend -H 'Content-Type: application/json' -d "@userInfo.json"
//curl.exe -v -X POST http://localhost:3000/BlockFriend -H 'Content-Type: application/json' -d "@userInfo.json"
//curl.exe -v -X POST http://localhost:3000/FriendStatus -H 'Content-Type: application/json' -d "@userInfo.json"

// ctrl + c to terminate the server after using command go run .

package main

import (
	"log"
	"net/http"
	"strconv"

	//packages added from tutorial
	"main/api"
	"main/database"
	"main/migrations"

	//this may change, I believe they just want us to reference main/customevents in our
	//own files, however im gonna leave it like this until im sure
	//"github.com/dipeshdulal/event-scheduling/customevents"

	"github.com/gorilla/mux"
)

// "middleware" is a function can is called before the final handling function, in this case, this function is called on each incoming
// request before that request is handled by the handling function loaded onto the router, this allows us to keep track of what is called
// for testing purpouses
// function from: https://pkg.go.dev/github.com/gorilla/mux#section-readme
func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Do stuff here
		log.Printf("uri: %s method: %s", r.RequestURI, r.Method)
		// Call the next handler, which can be another middleware in the chain, or the final handler.
		next.ServeHTTP(w, r)
	})
}

// Main+functions were modified from: https://medium.com/@anshap1719/getting-started-with-angular-and-go-setting-up-a-boilerplate-project-8c273b81aa6
// the main function start the server
func main() {

	database.InitDatabase()
	migrations.MigrateFriends()

	//initalizing an HTTP request multiplexer- this can check to see if any of the incoming url match
	//those we load it with and then run the appropriate functions
	router := mux.NewRouter()

	router.Use(loggingMiddleware)
	//router.Use(optionsMiddleware)

	router.HandleFunc("/login", api.LoginFunc).Methods("POST", "OPTIONS")
	router.HandleFunc("/register", api.RegisterFunc).Methods("OPTIONS", "POST")

	router.HandleFunc("/EditToDo", api.EditToDo).Methods("POST", "PUT", "OPTIONS")
	router.HandleFunc("/ToDoStatus", api.ToDoStatus).Methods("POST", "OPTIONS")
	router.HandleFunc("/DeleteToDo", api.DeleteToDo).Methods("POST", "OPTIONS")

	router.HandleFunc("/EditCal", api.EditCal).Methods("POST", "OPTIONS")
	router.HandleFunc("/CalStatus", api.CalStatus).Methods("POST", "OPTIONS")
	router.HandleFunc("/CalDelete", api.DeleteCal).Methods("POST", "OPTIONS")

	router.HandleFunc("/RequestFriend", api.RequestFriend).Methods("POST", "OPTIONS")
	router.HandleFunc("/AcceptFriend", api.AcceptFriend).Methods("POST", "OPTIONS")
	router.HandleFunc("/BlockFriend", api.BlockFriend).Methods("POST", "OPTIONS")
	router.HandleFunc("/FriendStatus", api.FriendStat).Methods("POST", "OPTIONS")

	router.HandleFunc("/DeleteUser", api.DeleteUser).Methods("POST", "OPTIONS")

	//add default handler

	//set port (backend), server will run on local host (your pc address)
	const port = 3000
	const server = "localhost"

	//https://pkg.go.dev/github.com/gorilla/mux#CORSMethodMiddleware
	router.Use(mux.CORSMethodMiddleware(router))

	//log.Printf shows date and time- could also just use Printf, but log better practice
	log.Printf("starting server on http://%s:%d", server, port)
	//start the web server, listen for requests sent to the server
	err := http.ListenAndServe(server+":"+strconv.Itoa(port), router)
	//if something does not work, (exit status 1) ie. if someone tries to use the same port
	log.Fatal(err)
}
