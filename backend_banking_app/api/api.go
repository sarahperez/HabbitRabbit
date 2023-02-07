package api

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"

	"example.com/backend_banking_app/helpers"
	"example.com/backend_banking_app/users"

	"github.com/gorilla/mux"
)

type Login struct {
	Username string
	Password string
}

type ErrResponse struct {
	Message string
}

func login(w http.ResponseWriter, r *http.Request) {
	//read body
	body, err := io.ReadAll(r.Body)
	//verify that everything is working correctly
	helpers.HandleErr(err)

	//handle login
	var formattedBody Login
	err = json.Unmarshal(body, &formattedBody)
	helpers.HandleErr(err)
	login := users.Login(formattedBody.Username, formattedBody.Password)

	//prepare response
	if login["message"] == "all is fine" {
		resp := login
		json.NewEncoder(w).Encode(resp)
	} else {
		//Handle error
		resp := ErrResponse{Message: "Wrong username of password"}
		json.NewEncoder(w).Encode(resp)
	}
}

func StartApi() {
	//creates router
	router := mux.NewRouter()
	//http listener and 8888 port
	router.HandleFunc("/login", login).Methods("POST")
	fmt.Println("App is working on port :8888")
	log.Fatal(http.ListenAndServe(":8888", router))
}