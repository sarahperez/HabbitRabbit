// to perform tests:
//cd into file you want to test, and type go test
//if multiple functions in test file, do go test -v instead of go test

//References used for testing:
//marshaling method for setting up http request bodies from https://golang.cafe/blog/golang-json-marshal-example.html
//converting response body to string method from https://www.educative.io/answers/how-to-read-the-response-body-in-golang
//overall structure for testing functions https://golang.cafe/blog/golang-httptest-example.html

package api

import (
	"encoding/json"
	"main/database"
	"main/interfaces"
	"main/migrations"
	//"main/users"

	"bytes"
	"io"
	"log"
	"net/http"
	"testing"

	//"github.com/stretchr/testify/assert"
	"net/http/httptest"
)

// TestReadBody has been started
//func TestReadBody(t *testing.T) {
//req := httptest.NewRequest(http.MethodGet, nil)
//actualBody := readBody(req)

//}

// TestApiResponse hasnt been started
//func TestApiResponse(t *testing.T) {}

// TestLoginFunc does not run
func TestLoginFunc(t *testing.T) {
	//database.InitTestDatabase()
	//migrations.Migrate()

	//register := users.Register("JamesJoe3", "James", "jamesjoe3@email.com", "J@mesPw0rd!")
	//apiResponse(register, w)

	//reqBody, err := json.Marshal(users.Login{Username: "JamesJoe3", Password: "J@mesPw0rd!"})
	//if err != nil {
	//	log.Print("error encountered in marshal")
	//}

	//req := httptest.NewRequest(http.MethodPost, "/login", bytes.NewBuffer(reqBody))
	//req.Header.Set("Contest-type", "application/json")
	//w := httptest.NewRecorder()
	//EditToDo(w, req)

	//res := w.Result()
	//defer res.Body.Close()
	//data, err := io.ReadAll(res.Body)
	//bodyString := string(data)

	//type Log struct {
	//	Username string
	//	Password string
	//}

	//expected, err := json.Marshal(Log{Username: "JamesJoe3", Password: "J@mesPw0rd!"})
	//if err != nil {
	//	log.Print("error encountered in marshal")
	//}

	//expectedString := string(expected)
	//expectedString += "\n"

	//check to see if the expected matches what was returned from the tested functions
	//if bodyString != expectedString {
	//	t.Errorf("error- failed response: %v", string(data))
	//}

	//reqBody, err = json.Marshal(users.Login{Username: "JamesJoe3", Password: "J@mesPw0rd!"})
	//if err != nil {
	//	log.Print("error encountered in marshal")
	//}

	//reqBody, err = json.Marshal(interfaces.User{Username: "JohnMark123", Name: "johanthan", Email: "johnsemail@email.com", Password: "JohN$pw0rd!"})
	//if err != nil {
	//	log.Print("error encountered in marshal")
	//}

	//req = httptest.NewRequest(http.MethodDelete, "/Login", bytes.NewBuffer(reqBody))
	//req.Header.Set("Content-type", "application/json")
	//w = httptest.NewRecorder()
	//EditToDo(w, req)
}

// TestRegisterFunc hasnt been started
func TestRegisterFunc(t *testing.T) {
	database.InitTestDatabase()
	//migrations.Migrate()

	reqBody, err := json.Marshal(interfaces.User{Username: "JohnMark123", Name: "johanthan", Email: "johnsemail@email.com", Password: "JohN$pw0rd!"})
	if err != nil {
		log.Print("error encountered in marshal")
	}

	req := httptest.NewRequest(http.MethodPost, "/Register", bytes.NewBuffer(reqBody))
	req.Header.Set("Content-type", "application/json")
	w := httptest.NewRecorder()
	RegisterFunc(w, req)

	//res := w.Result()
	//defer res.Body.Close()
	//data, err := io.ReadAll(res.Body)
	//bodyString := string(data)

	//type registration struct {
	//	Username string
	//	Name     string
	//	Mail     string
	//	Pass     string
	//}

	//expected, err := json.Marshal(registration{Username: "JohnMark123", Name: "johanthan", Mail: "johnsemail@email.com", Pass: "JohN$pw0rd!"})
	//expectedString := string(expected) + "\n"
	//if bodyString != expectedString {
	//	t.Errorf("error- failed response: %v", bodyString)
	//	//string(data))
	//}
	reqL := httptest.NewRequest(http.MethodPost, "/Login", bytes.NewBuffer(reqBody))
	reqL.Header.Set("Content-type", "application/json")
	wL := httptest.NewRecorder()
	LoginFunc(wL, reqL)

	//res := wL.Result()
	//defer res.Body.Close()
	//data, err := io.ReadAll(res)
	//bodyString := string(data)
	//if bodyString != "all is fine" {
	//	t.Errorf("error- failed response: %v", bodyString)

	//}

	reqBody, err = json.Marshal(interfaces.User{Username: "JohnMark123", Name: "johanthan", Email: "johnsemail@email.com", Password: "JohN$pw0rd!"})
	if err != nil {
		log.Print("error encountered in marshal")
	}

	req = httptest.NewRequest(http.MethodDelete, "/Register", bytes.NewBuffer(reqBody))
	req.Header.Set("Content-type", "application/json")
	w = httptest.NewRecorder()
	RegisterFunc(w, req)

	reqL = httptest.NewRequest(http.MethodDelete, "/Login", bytes.NewBuffer(reqBody))
	reqL.Header.Set("Content-type", "application/json")
	wL = httptest.NewRecorder()
	LoginFunc(wL, reqL)

}

// TestGoHome runs and passes
func TestGoHome(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/home-page", nil)
	w := httptest.NewRecorder()
	GoHome(w, req)
	res := w.Result()
	defer res.Body.Close()
	data, err := io.ReadAll(res.Body)
	if err != nil {
		t.Errorf("expected error to be nil got %v", err)
	}
	if string(data) != "Welcome to the home page, get request recieved" {
		t.Errorf("expected: Welcome to the home page, get request recieved. Got: %v", string(data))
	}
}

// TestDisplayCalender runs and passes
func TestDisplayCalender(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/calender-page", nil)
	w := httptest.NewRecorder()
	DisplayCalendar(w, req)
	res := w.Result()
	defer res.Body.Close()
	data, err := io.ReadAll(res.Body)
	if err != nil {
		t.Errorf("expected error to be nil got %v", err)
	}
	if string(data) != "Welcome to the calendar page, get request recieved" {
		t.Errorf("expected: Welcome to the calender page, get request recieved. Got: %v", string(data))
	}
}

// mini test- tests if the to-do item is sucessfully added and retrieved from the database
func TestAddingToList(t *testing.T) {

	//start the database used for testing
	database.InitTestDatabase()
	//migrations.Migrate()

	//create a request body
	reqBody, err := json.Marshal(interfaces.TodoReq{User: 1000, Description: "buy apples"})
	if err != nil {
		log.Print("error encountered in marshal")
	}

	//create a http request and send it to the function
	req := httptest.NewRequest(http.MethodPost, "/EditToDo", bytes.NewBuffer(reqBody))
	req.Header.Set("Content-type", "application/json")
	w := httptest.NewRecorder()
	EditToDo(w, req)

	//create a request body
	reqBody, err = json.Marshal(interfaces.UserID{User: 1000})
	if err != nil {
		log.Print("error encountered in marshal")
	}

	//create a http request and send it to the function
	req = httptest.NewRequest(http.MethodPost, "/ToDoStatus", bytes.NewBuffer(reqBody))
	req.Header.Set("Content-type", "application/json")
	w = httptest.NewRecorder()
	ToDoStatus(w, req)

	//format the response body as a string
	res := w.Result()
	defer res.Body.Close()
	data, err := io.ReadAll(res.Body)
	bodyString := string(data)

	//set up the expected response
	type resp struct {
		Complete   []string
		Incomplete []string
		Percentage float64
	}

	arg1 := []string{}
	arg2 := []string{"buy apples"}

	expected, err := json.Marshal(resp{Complete: arg1, Incomplete: arg2, Percentage: 0})
	if err != nil {
		log.Print("error encountered in marshal")
	}

	expectedString := string(expected)
	expectedString += "\n"

	//check to see if the expected matches what was returned from the tested functions
	if bodyString != expectedString {
		t.Errorf("error- failed response: %v", string(data))
	}

	//-----------------------the following is just clearing the database, deleting all added rows--------------------------
	reqBody, err = json.Marshal(interfaces.TodoReq{User: 1000, Description: "buy apples"})
	if err != nil {
		log.Print("error encountered in marshal")
	}

	req = httptest.NewRequest(http.MethodDelete, "/EditToDo", bytes.NewBuffer(reqBody))
	req.Header.Set("Content-type", "application/json")
	w = httptest.NewRecorder()
	EditToDo(w, req)
}

// more extensive test of the to do list functionality- involving adding, editing, deleting and returning user status
func TestEditToDo(t *testing.T) {
	database.InitTestDatabase()
	migrations.Migrate()

	//create a request body
	reqBody, err := json.Marshal(interfaces.TodoReq{User: 1000, Description: "buy apples"})
	if err != nil {
		log.Print("error encountered in marshal")
	}

	//create a http request and send it to the function
	req := httptest.NewRequest(http.MethodPost, "/EditToDo", bytes.NewBuffer(reqBody))
	req.Header.Set("Contest-type", "application/json")
	w := httptest.NewRecorder()
	EditToDo(w, req)

	//create a request body
	reqBody, err = json.Marshal(interfaces.TodoReq{User: 1000, Description: "buy oranges"})
	if err != nil {
		log.Print("error encountered in marshal")
	}

	//create a http request and send it to the function
	req = httptest.NewRequest(http.MethodPost, "/EditToDo", bytes.NewBuffer(reqBody))
	req.Header.Set("Contest-type", "application/json")
	w = httptest.NewRecorder()
	EditToDo(w, req)

	//create a request body
	reqBody, err = json.Marshal(interfaces.TodoReq{User: 1000, Description: "buy bananas"})
	if err != nil {
		log.Print("error encountered in marshal")
	}

	//create a http request and send it to the function
	req = httptest.NewRequest(http.MethodPost, "/EditToDo", bytes.NewBuffer(reqBody))
	req.Header.Set("Contest-type", "application/json")
	w = httptest.NewRecorder()
	EditToDo(w, req)

	//create a request body
	reqBody, err = json.Marshal(interfaces.TodoReq{User: 1000, Description: "buy grapes"})
	if err != nil {
		log.Print("error encountered in marshal")
	}

	//create a http request and send it to the function
	req = httptest.NewRequest(http.MethodPost, "/EditToDo", bytes.NewBuffer(reqBody))
	req.Header.Set("Contest-type", "application/json")
	w = httptest.NewRecorder()
	EditToDo(w, req)

	//create a request body
	reqBody, err = json.Marshal(interfaces.TodoReq{User: 1000, Description: "buy apples"})
	if err != nil {
		log.Print("error encountered in marshal")
	}

	//create a http request and send it to the function
	req = httptest.NewRequest(http.MethodPut, "/EditToDo", bytes.NewBuffer(reqBody))
	req.Header.Set("Contest-type", "application/json")
	w = httptest.NewRecorder()
	EditToDo(w, req)

	//create a request body
	reqBody, err = json.Marshal(interfaces.TodoReq{User: 1000, Description: "buy grapes"})
	if err != nil {
		log.Print("error encountered in marshal")
	}

	//create a http request and send it to the function
	req = httptest.NewRequest(http.MethodDelete, "/EditToDo", bytes.NewBuffer(reqBody))
	req.Header.Set("Contest-type", "application/json")
	w = httptest.NewRecorder()
	EditToDo(w, req)

	//create a request body
	reqBody, err = json.Marshal(interfaces.UserID{User: 1000})
	if err != nil {
		log.Print("error encountered in marshal")
	}

	//create a http request and send it to the function
	req = httptest.NewRequest(http.MethodPost, "/ToDoStatus", bytes.NewBuffer(reqBody))
	req.Header.Set("Contest-type", "application/json")
	w = httptest.NewRecorder()
	ToDoStatus(w, req)

	//set up the expected response
	res := w.Result()
	defer res.Body.Close()
	data, err := io.ReadAll(res.Body)
	bodyString := string(data)

	type resp struct {
		Complete   []string
		Incomplete []string
		Percentage float64
	}

	arg1 := []string{"buy apples"}
	arg2 := []string{"buy oranges", "buy bananas"}

	expected, err := json.Marshal(resp{Complete: arg1, Incomplete: arg2, Percentage: 33})
	expectedString := string(expected) + "\n"

	//check to see if the expected matches what was returned from the tested functions
	if bodyString != expectedString {
		t.Errorf("error- failed response: %v", string(data))
	}

	//-----------------------the following is just clearing the database, deleting all added rows--------------------------
	reqBody, err = json.Marshal(interfaces.TodoReq{User: 1000, Description: "buy apples"})
	if err != nil {
		log.Print("error encountered in marshal")
	}

	req = httptest.NewRequest(http.MethodDelete, "/EditToDo", bytes.NewBuffer(reqBody))
	req.Header.Set("Content-type", "application/json")
	w = httptest.NewRecorder()
	EditToDo(w, req)

	reqBody, err = json.Marshal(interfaces.TodoReq{User: 1000, Description: "buy bananas"})
	if err != nil {
		log.Print("error encountered in marshal")
	}

	req = httptest.NewRequest(http.MethodDelete, "/EditToDo", bytes.NewBuffer(reqBody))
	req.Header.Set("Content-type", "application/json")
	w = httptest.NewRecorder()
	EditToDo(w, req)

	reqBody, err = json.Marshal(interfaces.TodoReq{User: 1000, Description: "buy oranges"})
	if err != nil {
		log.Print("error encountered in marshal")
	}

	req = httptest.NewRequest(http.MethodDelete, "/EditToDo", bytes.NewBuffer(reqBody))
	req.Header.Set("Content-type", "application/json")
	w = httptest.NewRecorder()
	EditToDo(w, req)
}
