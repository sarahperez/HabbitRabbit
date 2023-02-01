package main

import (
	"fmt"
	"log"
	"main/utils"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

func main() {

	//test print to see if main is running
	fmt.Println("starting server on ", os.Getenv("PORT"))

	//"A router receives and sends data on computer networks" - https://www.cisco.com/c/en/us/solutions/small-business/resource-center/networking/what-is-a-router.html
	//initalize a router
	r := mux.NewRouter()

	//if the URL path matches "/hello-world" call the helloworld function
	r.HandleFunc("/hello-world", helloWorld)

	// Solves Cross Origin Access Issue
	c := cors.New(cors.Options{
		AllowedOrigins: []string{"http://localhost:4200"},
	})
	handler := c.Handler(r)

	srv := &http.Server{
		Handler: handler,
		Addr:    ":" + os.Getenv("PORT"),
	}

	log.Fatal(srv.ListenAndServe())
}

func helloWorld(w http.ResponseWriter, r *http.Request) {
	var data = struct {
		Title string `json:"title"`
	}{
		Title: "Golang + Angular Starter Kit",
	}

	jsonBytes, err := utils.StructToJson(data)
	if err != nil {
		fmt.Print(err)
	}

	//this sets the header
	//"Content-Type" makes it so the server can inform the client that JSON data is being sent
	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonBytes)
	return
}
