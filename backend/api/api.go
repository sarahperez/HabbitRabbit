package api

// code derived from https://github.com/Duomly/go-bank-backend

import (
	"encoding/json"
	"io"
	"log"
	"net/http"

	"main/database"
	"main/helpers"
	"main/interfaces"
	"main/users"

	"github.com/gorilla/mux"
)

type Login struct {
	Username string
	Password string
}

type Register struct {
	Username string
	Name     string
	Email    string
	Password string
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

func LoginFunc(w http.ResponseWriter, request *http.Request) {
	//
	switch request.Method {
	case http.MethodOptions:
		//CORS Preflight request sent as a OPTIONS method before the actual request is sent- to check if "CORS protocol is being understood"
		//this is a kind of way to attempt to protect the server from bad requests coming from bad addresses
		//good resources and readings- https://developer.mozilla.org/en-US/docs/Glossary/Preflight_request https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/Access-Control-Allow-Headers
		w.Header().Set("Access-Control-Allow-Origin", "*")  //this is denoting the origin from which the preflight request may come, right now the star is indicating it can come from anywhere, but this can be changed for better security in the future
		w.Header().Set("Access-Control-Allow-Headers", "*") //is will allow the sent request following the preflight to have any type of header (indicated by the star)
		//w.Header().Set("Access-Control-Allow-Methods", "POST") //this is saying that the request following the preflight request should be a POST method
		return
	case http.MethodPost:
		//once the Preflight request is handled, then we can handle the post request that follows normally
		w.Header().Set("Access-Control-Allow-Origin", "*")
		log.Print("inside loginfunc")

		// Refactor login to use readBody
		body := readBody(request)
		var formattedBody Login
		err := json.Unmarshal(body, &formattedBody)
		helpers.HandleErr(err)

		login := users.Login(formattedBody.Username, formattedBody.Password)
		// Refactor login to use apiResponse function
		apiResponse(login, w)
	}
}

func RegisterFunc(w http.ResponseWriter, r *http.Request) {
	body := readBody(r)

	var formattedBody Register
	err := json.Unmarshal(body, &formattedBody)
	helpers.HandleErr(err)

	log.Print("inside register func")
	log.Print(formattedBody.Username, formattedBody.Name, formattedBody.Email, formattedBody.Password)
	register := users.Register(formattedBody.Username, formattedBody.Name, formattedBody.Email, formattedBody.Password)
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

// Copyright (c) 2020 Mohamad Fadhil
// code derived from https://github.com/sdil/learning/blob/master/go/todolist-mysql-go/todolist.go
func EditToDo(w http.ResponseWriter, request *http.Request) {

	//"unload" the input data from the request- should be a user ID and a task description
	body := readBody(request)
	var formattedBody interfaces.TodoReq
	err := json.Unmarshal(body, &formattedBody)
	helpers.HandleErr(err)

	//start making an "entry" object that will be used to add/update an entry in the database
	var entry interfaces.TodoItem
	entry.User = formattedBody.User
	entry.Description = formattedBody.Description

	switch request.Method {
	case http.MethodPost:
		//if the request is a post- then it means an entry is being added
		entry.Completed = false
		database.DB.Create(&entry)
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode("Item added")
		return
	case http.MethodDelete:
		var task interfaces.TodoItem
		database.DB.Table("todo_items").Where("User = ? AND description = ?", formattedBody.User, formattedBody.Description).First(&task)
		if err := database.DB.Table("todo_items").Where("ID = ?", task.ID).Update("Completed", true).Error; err != nil {
			json.NewEncoder(w).Encode("task could not be found, so it could not be completed/deleted")
		} else {
			json.NewEncoder(w).Encode("Task completion status now updated to completed")
		}
		return
	}
}

func ToDoStatus(w http.ResponseWriter, request *http.Request) {
	switch request.Method {
	case http.MethodGet:
		//"unload" the input data from the request- should be a user ID
		body := readBody(request)
		var formattedBody interfaces.UserID
		err := json.Unmarshal(body, &formattedBody)
		helpers.HandleErr(err)

		log.Print("curr ID ", formattedBody.ID)

		completed := database.GetCompletedItems(formattedBody.ID)
		incomplete := database.GetIncompleteItems(formattedBody.ID)

		var response = map[string]interface{}{"incomplete": incomplete, "complete": completed}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(response)
		return
	}
}

//-----------------------------------these might get deleted? -----------------------------------

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

func defaultFunc() {

}
