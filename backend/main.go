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

// ctrl + c to terminate the server after using command go run .

package main

import (
	"log"
	"net/http"
	"strconv"

	//packages added from tutorial
	"main/api"
	"main/database"

	//packeges from online
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

	//initalizing an HTTP request multiplexer- this can check to see if any of the incoming url match
	//those we load it with and then run the appropriate functions
	router := mux.NewRouter()

	router.Use(loggingMiddleware)
	//router.Use(optionsMiddleware)

	router.HandleFunc("/home-page", api.GoHome)
	router.HandleFunc("/calendar", api.DisplayCalendar)
	router.HandleFunc("/login", api.LoginFunc).Methods("POST", "OPTIONS")
	router.HandleFunc("/register", api.RegisterFunc).Methods("POST")

	//trying to add in a handler for all cases where URL does NOT match one of the above linked to the mux
	// defaultRouter := router.PathPrefix("").Subrouter()
	//router.HandleFunc("/", api.Options).Methods("OPTIONS")

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
