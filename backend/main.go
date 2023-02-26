//references:
//https://golang.ch/which-golang-router-to-use-for-what-cases/
// https://github.com/gorilla/mux
// https://medium.com/@anshap1719/getting-started-with-angular-and-go-setting-up-a-boilerplate-project-8c273b81aa6
// https://pkg.go.dev/net/http#ServeMux

// !!!!!!!!!!!!!!!!run go get github.com/rs/cors in terminal before running code!!!!!!!!!!!!!!!!

//after starting the server, open a new terminal and run these comands to test get http requests:
// curl.exe -v -X GET http://localhost:3000/home-page
// curl.exe -v -X GET http://localhost:3000/calendar

//run these commands to test post http requests:
//curl.exe -v -X POST http://localhost:3000/home-page -H 'Content-Type: application/json' -d "@userInfo.json"

//curl commands to test login and register
//curl.exe -v -X POST http://localhost:3000/login -H 'Content-Type: application/json' -d "@userInfo.json"
//curl.exe -v -X POST http://localhost:3000/register -H 'Content-Type: application/json' -d "@userInfo.json"

// ctrl + c to terminate the server after using command go run .

package main

import (
	"log"
	"strconv"

	//web server router package- up to date (made by GO)
	"net/http"
	//packages added from tutorial

	"main/api"

	"github.com/gorilla/mux"
)

func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Do stuff here
		log.Printf("uri: %s method: %s", r.RequestURI, r.Method)
		// Call the next handler, which can be another middleware in the chain, or the final handler.
		next.ServeHTTP(w, r)
	})
}

// func optionsMiddleware(next http.Handler) http.Handler {
// 	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
// 		w.Header().Set("Access-Control-Allow-Origin", "*")
// 		w.Header().Set("Access-Control-Allow-Headers", "*")
// 		// Do stuff here
// 		log.Printf("options request - setting header")
// 		// Call the next handler, which can be another middleware in the chain, or the final handler.
// 		w.Write([]byte("options recieved"))
// 	})
// }

// Main+functions were modified from: https://medium.com/@anshap1719/getting-started-with-angular-and-go-setting-up-a-boilerplate-project-8c273b81aa6
// the main function start the server
func main() {
	//initalizing an HTTP request multiplexer- this can check to see if any of the incoming url match
	//those we load it with and then run the appropriate functions
	router := mux.NewRouter()

	router.Use(loggingMiddleware)
	//router.Use(optionsMiddleware)

	router.HandleFunc("/home-page", api.GoHome)
	router.HandleFunc("/calendar", api.DisplayCalendar)
	router.HandleFunc("/login", api.LoginFunc).Methods("POST")
	router.HandleFunc("/login", api.Options).Methods("OPTIONS")
	router.HandleFunc("/register", api.RegisterFunc).Methods("POST")

	//trying to add in a handler for all cases where URL does NOT match one of the above linked to the mux
	// defaultRouter := router.PathPrefix("").Subrouter()
	//router.HandleFunc("/", api.Options).Methods("OPTIONS")

	//set port (backend)
	const port = 8888
	//server will run on local host (your pc address)
	const server = "localhost"

	// the angular client loaded at http://localhost:4200 will make requests to the go server at http://localhost:3000
	// since these addresses are different we need to set a CORS policy to allow requests from our client
	// c := cors.New(cors.Options{
	// 	//tell computer that it can accept requests the origin of the angular app
	// 	AllowedOrigins: []string{"http://localhost:4200"},
	// })
	router.Use(mux.CORSMethodMiddleware(router))

	//handler is assigned the "function" that will call getPageData if it passess the CORS and matches the path
	//handler := c.Handler(mux)

	//log.Printf shows date and time- could also just use Printf, but log better practice
	log.Printf("starting server on http://%s:%d", server, port)
	//start the web server
	//listen for requests sent to the server
	err := http.ListenAndServe(server+":"+strconv.Itoa(port), router)

	//err := http.ListenAndServe(host, handler)
	//if something does not work, (exit status 1) ie. if someone tries to use the same port
	log.Fatal(err)
}
