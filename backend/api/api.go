package api

// code from https://github.com/Duomly/go-bank-backend

import (
	"encoding/json"
	"io"
	"log"
	"net/http"

	"main/helpers"
	"main/users"

	"github.com/gorilla/mux"
)

type Login struct {
	Username string
	Password string
}

type Register struct {
	Username string
	Email    string
	Password string
}

type TransactionBody struct {
	UserId uint
	From   uint
	To     uint
	Amount int
}

// Create readBody function
func readBody(r *http.Request) []byte {
	body, err := io.ReadAll(r.Body)
	helpers.HandleErr(err)

	return body
}

// Refactor apiResponse
func apiResponse(call map[string]interface{}, w http.ResponseWriter) {
	if call["message"] == "all is fine" {
		resp := call
		json.NewEncoder(w).Encode(resp)
		// Handle error in else
	} else {
		//not sure why if and else as the same, look into this
		resp := call
		json.NewEncoder(w).Encode(resp)
	}
}

func LoginFunc(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	// Refactor login to use readBody
	log.Print("inside loginfunc")
	body := readBody(r)

	var formattedBody Login
	err := json.Unmarshal(body, &formattedBody)
	helpers.HandleErr(err)

	login := users.Login(formattedBody.Username, formattedBody.Password)
	// Refactor login to use apiResponse function
	apiResponse(login, w)
}

func RegisterFunc(w http.ResponseWriter, r *http.Request) {
	body := readBody(r)

	var formattedBody Register
	err := json.Unmarshal(body, &formattedBody)
	helpers.HandleErr(err)

	register := users.Register(formattedBody.Username, formattedBody.Email, formattedBody.Password)
	// Refactor register to use apiResponse function
	apiResponse(register, w)
}

func GetUserFunc(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userId := vars["id"]
	auth := r.Header.Get("Authorization")

	user := users.GetUser(userId, auth)
	apiResponse(user, w)
}

//--------------------------------------------our added functions----------------------------------------------------------

// practice struct to catch user data
type userinfo struct {
	Username string `json:"Username"`
	Password string `json:"Password"`
}

// this function will be called with the following URL: http://localhost:3000/home-page
// example code from https://golang.ch/which-golang-router-to-use-for-what-cases/ used as a reference
func GoHome(w http.ResponseWriter, request *http.Request) {

	switch request.Method {
	//if the request is a GET
	case http.MethodGet:
		//tell the client that we are sending a json (Header in pr)
		w.Header().Set("Content-Type", "application/json")
		//need to pass infromation in as a string of bytes
		w.Write([]byte("Welcome to the home page, get request recieved"))

	case http.MethodPost:
		//if the request is a POST (incoming data)

		//reference for decoding (structure taken from example)- https://www.alexedwards.net/blog/how-to-properly-parse-a-json-request-body
		//set up a struct object to decode the json file into
		var info userinfo
		//decode the json file
		error := json.NewDecoder(request.Body).Decode(&info)
		if error != nil {
			//if statement to deal with decoder errors
			log.Println("decoding unsucessful", error)
			http.Error(w, error.Error(), http.StatusBadRequest)
			return
		}

		//print the decoded info
		log.Println("decoded string:", info)
		//tell the client that we are sending a json (Header in pr)
		w.Header().Set("Content-Type", "application/json")
		//pass infromation back
		w.Write([]byte(info.Username))
		w.Write([]byte(info.Password))
	}
}

// this function will get called by the following URL: http://localhost:3000/calendar
// example code from https://golang.ch/which-golang-router-to-use-for-what-cases/ used as a reference
func DisplayCalendar(w http.ResponseWriter, request *http.Request) {

	switch request.Method {
	//if the request is a GET
	case http.MethodGet:
		//tell the client that we are sending a json (Header in pr)
		w.Header().Set("Content-Type", "application/json")
		//need to pass infromation in as a string of bytes
		w.Write([]byte("Welcome to the calendar page, get request recieved"))

	case http.MethodPost:
		//if the request is a POST (incoming data)
		//tell the client that we are sending a json (Header in pr)
		w.Header().Set("Content-Type", "application/json")
		//need to pass infromation in as a string of bytes
		w.Write([]byte("Welcome to the calendar page, post request recieved"))
	}
}

func Options(w http.ResponseWriter, request *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "*")
	log.Printf("inside options")
	w.Header().Set("Content-Type", "application/text")
	//need to pass infromation in as a string of bytes
	w.Write([]byte("options recieved"))

}

func defaultFunc() {

}
