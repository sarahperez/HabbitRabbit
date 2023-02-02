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

// specifies that the HTTP response writer will be the sink for JSON data
// sink meaning what the JSOn data is converted to?
func helloWorld(w http.ResponseWriter, r *http.Request) {
	//parsing data into a struct so the compiler can have a spot in memory for everything
	var data = struct {
		Title string `json:"title"`
	}{
		Title: "Golang + Angular Starter Kit",
	}
	//StructToJson converts data, which is complex to JSON
	jsonBytes, err := utils.StructToJson(data)
	//if something was converted, print it out
	if err != nil {
		fmt.Print(err)
	}

	//this sets the header, clients will be aware to accept json
	//"Content-Type" makes it so the server can inform the client that JSON data is being sent
	w.Header().Set("Content-Type", "application/json")
	//I believe this function places JSON data onto the server
	w.Write(jsonBytes)
	return
}
