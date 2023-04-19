package api

// this file contains code from https://github.com/Duomly/go-bank-backend/tree/Golang-course-Lesson-6

import (
	"encoding/json"
	"io"
	"log"
	"math"
	"net/http"

	"main/database"
	"main/helpers"
	"main/interfaces"
	"main/users"
)

// code from https://github.com/Duomly/go-bank-backend/tree/Golang-course-Lesson-6 below
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

// function to log a user in
func LoginFunc(w http.ResponseWriter, request *http.Request) {
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

		//set the header and unload the data
		w.Header().Set("Access-Control-Allow-Origin", "*")
		//log.Print("inside loginfunc")
		body := readBody(request)
		var formattedBody Login
		err := json.Unmarshal(body, &formattedBody)
		helpers.HandleErr(err)

		//call the login helper
		login := users.Login(formattedBody.Username, formattedBody.Password)
		// Refactor login to use apiResponse function
		apiResponse(login, w)
	}
}

// function to register a user
func RegisterFunc(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodOptions:
		w.Header().Set("Access-Control-Allow-Origin", "*")  //this is denoting the origin from which the preflight request may come, right now the star is indicating it can come from anywhere, but this can be changed for better security in the future
		w.Header().Set("Access-Control-Allow-Headers", "*") //is will allow the sent request following the preflight to have any type of header (indicated by the star)
		return
	case http.MethodPost:

		//set the header and unload the data
		w.Header().Set("Access-Control-Allow-Origin", "*")
		body := readBody(r)
		var formattedBody Register
		err := json.Unmarshal(body, &formattedBody)
		helpers.HandleErr(err)

		//log.Print("inside register func")
		//log.Print(formattedBody.Username, formattedBody.Name, formattedBody.Email, formattedBody.Password)

		//call the register helper
		register := users.Register(formattedBody.Username, formattedBody.Name, formattedBody.Email, formattedBody.Password)

		if register["message"] == "user added" {
			log.Print("user added")
			register = users.Login(formattedBody.Username, formattedBody.Password)
		}
		// Refactor register to use apiResponse function
		apiResponse(register, w)
	}
}

//--------------------------------------------our added functions----------------------------------------------------------

// practice struct to catch user data
type userinfo struct {
	Username string `json:"Username"`
	Password string `json:"Password"`
}

// function to delete a user from the database
func DeleteUser(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodOptions:
		w.Header().Set("Access-Control-Allow-Origin", "*")  //this is denoting the origin from which the preflight request may come, right now the star is indicating it can come from anywhere, but this can be changed for better security in the future
		w.Header().Set("Access-Control-Allow-Headers", "*") //is will allow the sent request following the preflight to have any type of header (indicated by the star)
		return
	case http.MethodPost:

		//set the header and unload the data
		w.Header().Set("Access-Control-Allow-Origin", "*")
		body := readBody(r)
		var formattedBody interfaces.UserID
		err := json.Unmarshal(body, &formattedBody)
		helpers.HandleErr(err)

		switch r.Method {
		case http.MethodPost:
			//locate the target in the database and delete it
			var item interfaces.CalendarItem
			if err := database.DB.Table("users").Where("ID", formattedBody.User).Delete(&item).Error; err != nil {
				json.NewEncoder(w).Encode("user could not be found, not deleted")
			} else {
				json.NewEncoder(w).Encode("user deleted")
			}
			return
		}
	}
}

// Copyright (c) 2020 Mohamad Fadhil
// code derived from https://github.com/sdil/learning/blob/master/go/todolist-mysql-go/todolist.go
func EditToDo(w http.ResponseWriter, request *http.Request) {

	if request.Method == http.MethodOptions {
		w.Header().Set("Access-Control-Allow-Origin", "*")  //this is denoting the origin from which the preflight request may come, right now the star is indicating it can come from anywhere, but this can be changed for better security in the future
		w.Header().Set("Access-Control-Allow-Headers", "*") //is will allow the sent request following the preflight to have any type of header (indicated by the star)
		return
	}

	w.Header().Set("Access-Control-Allow-Origin", "*")

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
	case http.MethodPut:
		//if the request is a put we need to edit the status of the task
		var task interfaces.TodoItem
		database.DB.Table("todo_items").Where("User = ? AND description = ?", formattedBody.User, formattedBody.Description).First(&task)
		if err := database.DB.Table("todo_items").Where("ID = ?", task.ID).Update("Completed", true).Error; err != nil {
			//if they cannot find the task in the database no action can be taken
			json.NewEncoder(w).Encode("task could not be found, so it could not be completed/deleted")
		} else {
			//send result message
			json.NewEncoder(w).Encode("Task completion status now updated to completed")
		}
		return
	}
}

// function to return the to do status of a user
func ToDoStatus(w http.ResponseWriter, request *http.Request) {

	if request.Method == http.MethodOptions {
		w.Header().Set("Access-Control-Allow-Origin", "*")  //this is denoting the origin from which the preflight request may come, right now the star is indicating it can come from anywhere, but this can be changed for better security in the future
		w.Header().Set("Access-Control-Allow-Headers", "*") //is will allow the sent request following the preflight to have any type of header (indicated by the star)
		return
	}

	w.Header().Set("Access-Control-Allow-Origin", "*")

	switch request.Method {
	case http.MethodPost:

		//"unload" the input data from the request
		body := readBody(request)
		log.Print(string(body))
		var formattedBody interfaces.UserID
		err := json.Unmarshal(body, &formattedBody)
		helpers.HandleErr(err)

		//log.Print("curr ID ", formattedBody.User)

		//fetch all the completed and incomplete tasks
		completed := database.GetCompletedItems(formattedBody.User)
		incomplete := database.GetIncompleteItems(formattedBody.User)

		//calculate the percent completed
		perComplete := 100.0 * float64(len(completed)) / (float64(len(completed)) + float64(len(incomplete)))

		var response = map[string]interface{}{"Incomplete": incomplete, "Complete": completed, "Percentage": math.Round(perComplete)}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(response)
		return
	}
}

// function to delete an item from the to do list
func DeleteToDo(w http.ResponseWriter, request *http.Request) {

	if request.Method == http.MethodOptions {
		w.Header().Set("Access-Control-Allow-Origin", "*")  //this is denoting the origin from which the preflight request may come, right now the star is indicating it can come from anywhere, but this can be changed for better security in the future
		w.Header().Set("Access-Control-Allow-Headers", "*") //is will allow the sent request following the preflight to have any type of header (indicated by the star)
		return
	}

	//set the header and unload the data
	w.Header().Set("Access-Control-Allow-Origin", "*")
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
		var task interfaces.TodoItem
		var item interfaces.TodoItem
		//locate the target in the database and delete it
		database.DB.Table("todo_items").Where("User = ? AND description = ?", formattedBody.User, formattedBody.Description).First(&task)
		if err := database.DB.Delete(&item, task.ID).Error; err != nil {
			json.NewEncoder(w).Encode("Task could not be found, so it could not be deleted")
		} else {
			json.NewEncoder(w).Encode("Task deleted")
		}
		return
	}
}

// function to add to the calendar
func EditCal(w http.ResponseWriter, request *http.Request) {
	if request.Method == http.MethodOptions {
		w.Header().Set("Access-Control-Allow-Origin", "*")  //this is denoting the origin from which the preflight request may come, right now the star is indicating it can come from anywhere, but this can be changed for better security in the future
		w.Header().Set("Access-Control-Allow-Headers", "*") //is will allow the sent request following the preflight to have any type of header (indicated by the star)
		return
	}

	//set the header and unload the data
	w.Header().Set("Access-Control-Allow-Origin", "*")
	body := readBody(request)
	var formattedBody interfaces.CalendarItem
	err := json.Unmarshal(body, &formattedBody)
	helpers.HandleErr(err)

	switch request.Method {
	case http.MethodPost:
		//add the event into the database
		database.DB.Create(&formattedBody)
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode("Item added")
		return
	}
}

// function to return the calendar items for a user
func CalStatus(w http.ResponseWriter, request *http.Request) {

	if request.Method == http.MethodOptions {
		w.Header().Set("Access-Control-Allow-Origin", "*")  //this is denoting the origin from which the preflight request may come, right now the star is indicating it can come from anywhere, but this can be changed for better security in the future
		w.Header().Set("Access-Control-Allow-Headers", "*") //is will allow the sent request following the preflight to have any type of header (indicated by the star)
		return
	}

	switch request.Method {
	case http.MethodPost:

		//set the header and unload the data
		w.Header().Set("Access-Control-Allow-Origin", "*")
		body := readBody(request)
		var formattedBody interfaces.UserID
		err := json.Unmarshal(body, &formattedBody)
		helpers.HandleErr(err)

		items := database.GetCalItems(formattedBody.User)
		var object interfaces.CalObj
		var objList []interfaces.CalObj

		//reduce the stored calendar infromation into what is nessisary for an event object in the front end
		for i := 0; i < len(items); i++ {
			object.EventID = items[i].EventID
			object.StartStr = items[i].StartStr
			object.EndStr = items[i].EndStr
			object.Title = items[i].Title
			objList = append(objList, object)
		}

		var response = map[string]interface{}{"items": objList}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(response)
		return
	}
}

// function to delete calendar object
func DeleteCal(w http.ResponseWriter, request *http.Request) {
	if request.Method == http.MethodOptions {
		w.Header().Set("Access-Control-Allow-Origin", "*")  //this is denoting the origin from which the preflight request may come, right now the star is indicating it can come from anywhere, but this can be changed for better security in the future
		w.Header().Set("Access-Control-Allow-Headers", "*") //is will allow the sent request following the preflight to have any type of header (indicated by the star)
		return
	}

	//set header and unmarshal data
	w.Header().Set("Access-Control-Allow-Origin", "*")
	body := readBody(request)
	var formattedBody interfaces.CalendarItem
	err := json.Unmarshal(body, &formattedBody)
	helpers.HandleErr(err)

	switch request.Method {
	case http.MethodPost:
		var item interfaces.CalendarItem
		//find the target in the database and delete it
		if err := database.DB.Table("calendar_items").Where("event_id = ?", formattedBody.EventID).Delete(&item).Error; err != nil {
			json.NewEncoder(w).Encode("task could not be found, so it could not be deleted")
		} else {
			json.NewEncoder(w).Encode("task deleted")
		}
		return
	}
}

// function to request to be friends
func RequestFriend(w http.ResponseWriter, request *http.Request) {

	if request.Method == http.MethodOptions {
		w.Header().Set("Access-Control-Allow-Origin", "*")  //this is denoting the origin from which the preflight request may come, right now the star is indicating it can come from anywhere, but this can be changed for better security in the future
		w.Header().Set("Access-Control-Allow-Headers", "*") //is will allow the sent request following the preflight to have any type of header (indicated by the star)
		return
	}

	switch request.Method {
	case http.MethodPost:
		w.Header().Set("Access-Control-Allow-Origin", "*")
		//"unload" the input data from the request- should be a user ID
		body := readBody(request)
		var formattedBody interfaces.FriendRequest
		err := json.Unmarshal(body, &formattedBody)
		helpers.HandleErr(err)

		//validate requester and reciever
		user := &interfaces.User{}
		if err := database.DB.Where("username = ? ", formattedBody.Requester).First(&user).Error; err != nil {
			json.NewEncoder(w).Encode("Requester not found")
			return
		}

		user = &interfaces.User{}
		if err := database.DB.Where("username = ? ", formattedBody.Reciever).First(&user).Error; err != nil {
			user := &interfaces.User{}
			if err := database.DB.Where("username = ? ", formattedBody.Reciever).First(&user).Error; err != nil {
				json.NewEncoder(w).Encode("Reciever and Requester not found")
			}
			json.NewEncoder(w).Encode("Reciever not found")
			return
		} else {
			//requester and reciever are valid...
			stat := &interfaces.FriendStatus{}
			if err := database.DB.Where("requester = ? AND reciever = ?", formattedBody.Requester, formattedBody.Reciever).First(&stat).Error; err != nil {
				//create an object to add to the database if the request does not exist yet
				var object interfaces.FriendStatus
				object.Requester = formattedBody.Requester
				object.Reciever = formattedBody.Reciever
				object.Status = "Requested"
				//add to database
				database.DB.Create(&object)
				json.NewEncoder(w).Encode("request sent")
			} else if stat.Status == "Requested" {
				json.NewEncoder(w).Encode("this connection was already requested, status unchanged")
			} else if stat.Status == "Accepted" {
				json.NewEncoder(w).Encode("already friends")
			} else {
				json.NewEncoder(w).Encode("the reciever has blocked this requester")
			}
		}
		return
	}
}

// function to accept a friend request
func AcceptFriend(w http.ResponseWriter, request *http.Request) {

	if request.Method == http.MethodOptions {
		w.Header().Set("Access-Control-Allow-Origin", "*")  //this is denoting the origin from which the preflight request may come, right now the star is indicating it can come from anywhere, but this can be changed for better security in the future
		w.Header().Set("Access-Control-Allow-Headers", "*") //is will allow the sent request following the preflight to have any type of header (indicated by the star)
		return
	}

	switch request.Method {
	case http.MethodPost:
		w.Header().Set("Access-Control-Allow-Origin", "*")
		//"unload" the input data from the request- should be a user ID
		body := readBody(request)
		var formattedBody interfaces.FriendRequest
		err := json.Unmarshal(body, &formattedBody)
		helpers.HandleErr(err)

		//check that requester and reciever are valid
		user := &interfaces.User{}
		if err := database.DB.Where("username = ? ", formattedBody.Requester).First(&user).Error; err != nil {
			json.NewEncoder(w).Encode("Requester not found")
			return
		}

		user = &interfaces.User{}
		if err := database.DB.Where("username = ? ", formattedBody.Reciever).First(&user).Error; err != nil {
			user := &interfaces.User{}
			if err := database.DB.Where("username = ? ", formattedBody.Reciever).First(&user).Error; err != nil {
				json.NewEncoder(w).Encode("Reciever and Requester not found")
			}
			json.NewEncoder(w).Encode("Reciever not found")
			return
		} else {
			//user and requester are valid...
			stat := &interfaces.FriendStatus{}
			if err := database.DB.Where("requester = ? AND reciever = ?", formattedBody.Requester, formattedBody.Reciever).First(&stat).Error; err != nil {
				//no request exists
				json.NewEncoder(w).Encode("no request to accept")
			} else if stat.Status == "Requested" {
				if err := database.DB.Table("friend_statuses").Where("ID = ?", stat.ID).Update("status", "Accepted").Error; err != nil {
					//error encountered
					json.NewEncoder(w).Encode("request not found")
				} else {
					//status updated
					json.NewEncoder(w).Encode("status now updated")
				}
			} else if stat.Status == "Accepted" {
				//users already friends- no action needed
				json.NewEncoder(w).Encode("already accepted")
			} else {
				//user is blocked- cannot accept friend request
				json.NewEncoder(w).Encode("requester has been blocked")
			}
		}
		return
	}
}

// function to block a friend request
func BlockFriend(w http.ResponseWriter, request *http.Request) {

	if request.Method == http.MethodOptions {
		w.Header().Set("Access-Control-Allow-Origin", "*")  //this is denoting the origin from which the preflight request may come, right now the star is indicating it can come from anywhere, but this can be changed for better security in the future
		w.Header().Set("Access-Control-Allow-Headers", "*") //is will allow the sent request following the preflight to have any type of header (indicated by the star)
		return
	}

	switch request.Method {
	case http.MethodPost:

		//set header and unload data
		w.Header().Set("Access-Control-Allow-Origin", "*")
		body := readBody(request)
		var formattedBody interfaces.FriendRequest
		err := json.Unmarshal(body, &formattedBody)
		helpers.HandleErr(err)

		//check that the user is a valid user in the database
		user := &interfaces.User{}
		if err := database.DB.Where("username = ? ", formattedBody.Requester).First(&user).Error; err != nil {
			json.NewEncoder(w).Encode("Requester not found")
			return
		}

		//validate that a request exists
		user = &interfaces.User{}
		if err := database.DB.Where("username = ? ", formattedBody.Reciever).First(&user).Error; err != nil {
			user := &interfaces.User{}
			if err := database.DB.Where("username = ? ", formattedBody.Reciever).First(&user).Error; err != nil {
				json.NewEncoder(w).Encode("Reciever and Requester not found")
			}
			json.NewEncoder(w).Encode("Reciever not found")
			return
		} else {
			//if the requester and reciever are valid users...
			stat := &interfaces.FriendStatus{}
			if err := database.DB.Where("requester = ? AND reciever = ?", formattedBody.Requester, formattedBody.Reciever).First(&stat).Error; err != nil {
				//no request exists
				json.NewEncoder(w).Encode("no request to block")
			} else if stat.Status == "Requested" || stat.Status == "Accepted" {
				if err := database.DB.Table("friend_statuses").Where("ID = ?", stat.ID).Update("status", "Blocked").Error; err != nil {
					//error encountered in updating the status
					json.NewEncoder(w).Encode("request not found")
				} else {
					//status updated
					json.NewEncoder(w).Encode("status now updated")
				}
			} else {
				//user was already blocked- no action needed
				json.NewEncoder(w).Encode("requester has already been blocked")
			}
		}
		return
	}
}

// function to return the friends, requests and blocked users for a certian user
func FriendStat(w http.ResponseWriter, request *http.Request) {

	if request.Method == http.MethodOptions {
		w.Header().Set("Access-Control-Allow-Origin", "*")  //this is denoting the origin from which the preflight request may come, right now the star is indicating it can come from anywhere, but this can be changed for better security in the future
		w.Header().Set("Access-Control-Allow-Headers", "*") //is will allow the sent request following the preflight to have any type of header (indicated by the star)
		return
	}

	switch request.Method {
	case http.MethodPost:

		//set header and unload data
		w.Header().Set("Access-Control-Allow-Origin", "*")
		body := readBody(request)
		var formattedBody interfaces.UserID
		err := json.Unmarshal(body, &formattedBody)
		helpers.HandleErr(err)

		user := &interfaces.User{}
		if err := database.DB.Where("ID = ?", formattedBody.User).First(&user).Error; err != nil {
			json.NewEncoder(w).Encode("user not in database")
		} else {
			//collect all of the needed fields from the database
			name := user.Username
			friends := database.GetFriends(name)
			requests := database.GetRequests(name)
			blocked := database.GetBlocked(name)

			//prepare response
			var response = map[string]interface{}{"Friends": friends, "Requests from": requests, "Blocked Users": blocked}

			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(response)
			return
		}

		return
	}
}

// function to delete a request from the database- testing only function
func DeleteRequest(w http.ResponseWriter, request *http.Request) {

	if request.Method == http.MethodOptions {
		w.Header().Set("Access-Control-Allow-Origin", "*")  //this is denoting the origin from which the preflight request may come, right now the star is indicating it can come from anywhere, but this can be changed for better security in the future
		w.Header().Set("Access-Control-Allow-Headers", "*") //is will allow the sent request following the preflight to have any type of header (indicated by the star)
		return
	}

	switch request.Method {
	case http.MethodPost:

		//set header and unload the data
		w.Header().Set("Access-Control-Allow-Origin", "*")
		body := readBody(request)
		var formattedBody interfaces.FriendRequest
		err := json.Unmarshal(body, &formattedBody)
		helpers.HandleErr(err)

		stat := &interfaces.FriendStatus{}
		obj := &interfaces.FriendStatus{}
		//check to see if the request exists, and delete it if it does
		if err := database.DB.Where("requester = ? AND reciever = ?", formattedBody.Requester, formattedBody.Reciever).First(&stat).Error; err != nil {
			database.DB.Table("friend_statuses").Where("ID", stat.ID).Delete(&obj)
			json.NewEncoder(w).Encode("request deleted")
		} else {
			json.NewEncoder(w).Encode("could not delete request")
		}
	}
	return
}
