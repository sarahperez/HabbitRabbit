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
	"io"
	"main/database"
	"main/interfaces"
	"main/migrations"

	"bytes"
	"log"
	"net/http"
	"testing"

	//"github.com/stretchr/testify/assert"
	"net/http/httptest"
)

// TestLoginFunc is not necessary as login is called in the testRegisterFunc

func TestRegisterFunc(t *testing.T) {
	database.InitTestDatabase()

	reqBody, err := json.Marshal(interfaces.User{Username: "JohnMark123", Name: "johanthan", Email: "johnsemail@email.com", Password: "JohN$pw0rd!"})
	if err != nil {
		log.Print("error encountered in marshal")
	}

	req := httptest.NewRequest(http.MethodPost, "/register", bytes.NewBuffer(reqBody))
	req.Header.Set("Contest-type", "application/json")
	w := httptest.NewRecorder()
	RegisterFunc(w, req)

	reqBody, err = json.Marshal(interfaces.User{Username: "JohnMark123", Name: "johanthan", Email: "johnsemail@email.com", Password: "JohN$pw0rd!"})
	if err != nil {
		log.Print("error encountered in marshal")
	}

	req = httptest.NewRequest(http.MethodPost, "/login", bytes.NewBuffer(reqBody))
	req.Header.Set("Contest-type", "application/json")
	w = httptest.NewRecorder()
	LoginFunc(w, req)

	reqBody, err = json.Marshal(interfaces.UserID{User: 1})
	if err != nil {
		log.Print("error encountered in marshal")
	}

	req = httptest.NewRequest(http.MethodPost, "/DeleteUser", bytes.NewBuffer(reqBody))
	req.Header.Set("Contest-type", "application/json")
	w = httptest.NewRecorder()
	DeleteUser(w, req)
}

// mini test- tests if the to-do item is sucessfully added and retrieved from the database
func TestAddingToList(t *testing.T) {
	database.InitTestDatabase()

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
	reqBody, err = json.Marshal(interfaces.TodoItem{User: 1000, Description: "buy apples", Completed: false})
	if err != nil {
		log.Print("error encountered in marshal")
	}

	req = httptest.NewRequest(http.MethodPost, "/DeleteToDo", bytes.NewBuffer(reqBody))
	req.Header.Set("Content-type", "application/json")
	w = httptest.NewRecorder()
	DeleteToDo(w, req)
}

// more extensive test of the to do list functionality- involving adding, editing, deleting and returning user status
func TestEditToDo(t *testing.T) {
	database.InitTestDatabase()
	migrations.MigrateToDo()

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
	reqBody, err = json.Marshal(interfaces.TodoItem{User: 1000, Description: "buy grapes", Completed: false})
	if err != nil {
		log.Print("error encountered in marshal")
	}

	//create a http request and send it to the function
	req = httptest.NewRequest(http.MethodPost, "/DeleteToDo", bytes.NewBuffer(reqBody))
	req.Header.Set("Content-type", "application/json")
	w = httptest.NewRecorder()
	DeleteToDo(w, req)

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
	reqBody, err = json.Marshal(interfaces.TodoItem{User: 1000, Description: "buy apples", Completed: true})
	if err != nil {
		log.Print("error encountered in marshal")
	}

	req = httptest.NewRequest(http.MethodPost, "/DeleteToDo", bytes.NewBuffer(reqBody))
	req.Header.Set("Content-type", "application/json")
	w = httptest.NewRecorder()
	DeleteToDo(w, req)

	reqBody, err = json.Marshal(interfaces.TodoItem{User: 1000, Description: "buy oranges", Completed: false})
	if err != nil {
		log.Print("error encountered in marshal")
	}

	req = httptest.NewRequest(http.MethodPost, "/DeleteToDo", bytes.NewBuffer(reqBody))
	req.Header.Set("Content-type", "application/json")
	w = httptest.NewRecorder()
	DeleteToDo(w, req)

	reqBody, err = json.Marshal(interfaces.TodoItem{User: 1000, Description: "buy bananas", Completed: false})
	if err != nil {
		log.Print("error encountered in marshal")
	}

	req = httptest.NewRequest(http.MethodPost, "/DeleteToDo", bytes.NewBuffer(reqBody))
	req.Header.Set("Content-type", "application/json")
	w = httptest.NewRecorder()
	DeleteToDo(w, req)
}

func TestDeleteToDo(t *testing.T) {
	database.InitTestDatabase()
	migrations.MigrateToDo()

	reqBody, err := json.Marshal(interfaces.TodoReq{User: 1000, Description: "buy apples"})
	if err != nil {
		log.Print("error encountered in marshal")
	}

	//create a http request and send it to the function
	req := httptest.NewRequest(http.MethodPost, "/EditToDo", bytes.NewBuffer(reqBody))
	req.Header.Set("Contest-type", "application/json")
	w := httptest.NewRecorder()
	EditToDo(w, req)

	reqBody, err = json.Marshal(interfaces.TodoReq{User: 1000, Description: "buy grapes"})
	if err != nil {
		log.Print("error encountered in marshal")
	}

	//create a http request and send it to the function
	req = httptest.NewRequest(http.MethodPost, "/EditToDo", bytes.NewBuffer(reqBody))
	req.Header.Set("Contest-type", "application/json")
	w = httptest.NewRecorder()
	EditToDo(w, req)

	reqBody, err = json.Marshal(interfaces.TodoReq{User: 1000, Description: "buy grapes"})
	if err != nil {
		log.Print("error encountered in marshal")
	}

	//create a http request and send it to the function
	req = httptest.NewRequest(http.MethodPut, "/EditToDo", bytes.NewBuffer(reqBody))
	req.Header.Set("Contest-type", "application/json")
	w = httptest.NewRecorder()
	EditToDo(w, req)

	reqBody, err = json.Marshal(interfaces.TodoReq{User: 1000, Description: "buy kiwi"})
	if err != nil {
		log.Print("error encountered in marshal")
	}

	//create a http request and send it to the function
	req = httptest.NewRequest(http.MethodPost, "/EditToDo", bytes.NewBuffer(reqBody))
	req.Header.Set("Contest-type", "application/json")
	w = httptest.NewRecorder()
	EditToDo(w, req)

	reqBody, err = json.Marshal(interfaces.TodoItem{User: 1000, Description: "buy apples", Completed: false})
	if err != nil {
		log.Print("error encountered in marshal")
	}

	//create a http request and send it to the function
	req = httptest.NewRequest(http.MethodPost, "/DeleteToDo", bytes.NewBuffer(reqBody))
	req.Header.Set("Contest-type", "application/json")
	w = httptest.NewRecorder()
	DeleteToDo(w, req)

	reqBody, err = json.Marshal(interfaces.TodoItem{User: 1000, Description: "buy grapes", Completed: true})
	if err != nil {
		log.Print("error encountered in marshal")
	}

	//create a http request and send it to the function
	req = httptest.NewRequest(http.MethodPost, "/DeleteToDo", bytes.NewBuffer(reqBody))
	req.Header.Set("Contest-type", "application/json")
	w = httptest.NewRecorder()
	DeleteToDo(w, req)

	reqBody, err = json.Marshal(interfaces.UserID{User: 1000})
	if err != nil {
		log.Print("error encountered in marshal")
	}

	//create a http request and send it to the function
	req = httptest.NewRequest(http.MethodPost, "/ToDoStatus", bytes.NewBuffer(reqBody))
	req.Header.Set("Contest-type", "application/json")
	w = httptest.NewRecorder()
	ToDoStatus(w, req)

	res := w.Result()
	defer res.Body.Close()
	data, err := io.ReadAll(res.Body)
	bodyString := string(data)

	type resp struct {
		Complete   []string
		Incomplete []string
		Percentage float64
	}

	arg1 := []string{}
	arg2 := []string{"buy kiwi"}

	expected, err := json.Marshal(resp{Complete: arg1, Incomplete: arg2, Percentage: 0})
	expectedString := string(expected) + "\n"

	if bodyString != expectedString {
		t.Errorf("error- failed response: %v", string(data))
	}

	reqBody, err = json.Marshal(interfaces.TodoReq{User: 1000, Description: "buy kiwi"})
	if err != nil {
		log.Print("error encountered in marshal")
	}

	//create a http request and send it to the function
	req = httptest.NewRequest(http.MethodPost, "/DeleteToDo", bytes.NewBuffer(reqBody))
	req.Header.Set("Contest-type", "application/json")
	w = httptest.NewRecorder()
	DeleteToDo(w, req)
}

func TestEditCal(t *testing.T) {
	database.InitTestDatabase()
	migrations.MigrateCal()

	//create a request body
	calBody, err := json.Marshal(interfaces.CalendarItem{User: 100, EventID: 1560, StartStr: "2023-04-23T11:00:00", EndStr: "2023-04-23T12:00:00", Title: "SWE Meeting"})
	if err != nil {
		log.Print("error encountered in marshal")
	}

	//create a http request and send it to the function
	cal := httptest.NewRequest(http.MethodPost, "/EditCal", bytes.NewBuffer(calBody))
	cal.Header.Set("Contest-type", "application/json")
	w := httptest.NewRecorder()
	EditCal(w, cal)

	//create a request body
	calBody, err = json.Marshal(interfaces.CalendarItem{User: 100, EventID: 150, StartStr: "2023-04-19T10:00:00", EndStr: "2023-04-19T11:30:00", Title: "Film SWE Sprint 3"})
	if err != nil {
		log.Print("error encountered in marshal")
	}

	//create a http request and send it to the function
	cal = httptest.NewRequest(http.MethodPost, "/EditCal", bytes.NewBuffer(calBody))
	cal.Header.Set("Contest-type", "application/json")
	w = httptest.NewRecorder()
	EditCal(w, cal)

	//create a request body
	calBody, err = json.Marshal(interfaces.CalendarItem{User: 100, EventID: 17, StartStr: "2023-04-20T10:40:00", EndStr: "2023-04-20T11:30:00", Title: "Physics 2 Quiz"})
	if err != nil {
		log.Print("error encountered in marshal")
	}

	//create a http request and send it to the function
	cal = httptest.NewRequest(http.MethodPost, "/EditCal", bytes.NewBuffer(calBody))
	cal.Header.Set("Contest-type", "application/json")
	w = httptest.NewRecorder()
	EditCal(w, cal)

	//create a request body
	calBody, err = json.Marshal(interfaces.CalendarItem{User: 100, EventID: 56, StartStr: "2023-04-24T9:35:00", EndStr: "2023-04-24T10:25:00", Title: "CAP3032 Meeting"})
	if err != nil {
		log.Print("error encountered in marshal")
	}

	//create a http request and send it to the function
	cal = httptest.NewRequest(http.MethodPost, "/EditCal", bytes.NewBuffer(calBody))
	cal.Header.Set("Contest-type", "application/json")
	w = httptest.NewRecorder()
	EditCal(w, cal)

	//create a request body
	calBody, err = json.Marshal(interfaces.CalendarItem{User: 100, EventID: 56, StartStr: "2023-04-24T9:35:00", EndStr: "2023-04-24T10:25:00", Title: "CAP3032 Meeting"})
	if err != nil {
		log.Print("error encountered in marshal")
	}

	//create a http request and send it to the function
	cal = httptest.NewRequest(http.MethodPost, "/CalDelete", bytes.NewBuffer(calBody))
	cal.Header.Set("Contest-type", "application/json")
	w = httptest.NewRecorder()
	DeleteCal(w, cal)

	//create a request body
	calBody, err = json.Marshal(interfaces.CalendarItem{User: 100})
	if err != nil {
		log.Print("error encountered in marshal")
	}

	//create a http request and send it to the function
	cal = httptest.NewRequest(http.MethodPost, "/CalStatus", bytes.NewBuffer(calBody))
	cal.Header.Set("Contest-type", "application/json")
	w = httptest.NewRecorder()
	CalStatus(w, cal)

	//set up the expected response
	res := w.Result()
	defer res.Body.Close()
	data, err := io.ReadAll(res.Body)
	bodyString := string(data)

	expectedString := "{\"items\":[{\"EventID\":1560,\"StartStr\":\"2023-04-23T11:00:00\",\"EndStr\":\"2023-04-23T12:00:00\",\"Title\":\"SWE Meeting\"},{\"EventID\":150,\"StartStr\":\"2023-04-19T10:00:00\",\"EndStr\":\"2023-04-19T11:30:00\",\"Title\":\"Film SWE Sprint 3\"},{\"EventID\":17,\"StartStr\":\"2023-04-20T10:40:00\",\"EndStr\":\"2023-04-20T11:30:00\",\"Title\":\"Physics 2 Quiz\"}]}\n"

	//check to see if the expected matches what was returned from the tested functions
	if bodyString != expectedString {
		t.Errorf("error- failed response: %v", string(data))
	}

	//-----------------------the following is just clearing the database, deleting all added rows--------------------------
	calBody, err = json.Marshal(interfaces.CalendarItem{User: 100, EventID: 1560, StartStr: "2023-04-23T11:00:00", EndStr: "2023-04-23T12:00:00", Title: "SWE Meeting"})
	if err != nil {
		log.Print("error encountered in marshal")
	}

	//create a http request and send it to the function
	cal = httptest.NewRequest(http.MethodPost, "/CalDelete", bytes.NewBuffer(calBody))
	cal.Header.Set("Contest-type", "application/json")
	w = httptest.NewRecorder()
	DeleteCal(w, cal)

	calBody, err = json.Marshal(interfaces.CalendarItem{User: 100, EventID: 150, StartStr: "2023-04-19T10:00:00", EndStr: "2023-04-19T11:30:00", Title: "Film SWE Sprint 3"})
	if err != nil {
		log.Print("error encountered in marshal")
	}

	//create a http request and send it to the function
	cal = httptest.NewRequest(http.MethodPost, "/CalDelete", bytes.NewBuffer(calBody))
	cal.Header.Set("Contest-type", "application/json")
	w = httptest.NewRecorder()
	DeleteCal(w, cal)

	calBody, err = json.Marshal(interfaces.CalendarItem{User: 100, EventID: 17, StartStr: "2023-04-20T10:40:00", EndStr: "2023-04-20T11:30:00", Title: "Physics 2 Quiz"})
	if err != nil {
		log.Print("error encountered in marshal")
	}

	//create a http request and send it to the function
	cal = httptest.NewRequest(http.MethodPost, "/CalDelete", bytes.NewBuffer(calBody))
	cal.Header.Set("Contest-type", "application/json")
	w = httptest.NewRecorder()
	DeleteCal(w, cal)
}

func TestDeleteCal(t *testing.T) {
	calBody, err := json.Marshal(interfaces.CalendarItem{User: 100, EventID: 1560, StartStr: "2023-04-23T11:00:00", EndStr: "2023-04-23T12:00:00", Title: "SWE Meeting"})
	if err != nil {
		log.Print("error encountered in marshal")
	}

	//create a http request and send it to the function
	cal := httptest.NewRequest(http.MethodPost, "/EditCal", bytes.NewBuffer(calBody))
	cal.Header.Set("Contest-type", "application/json")
	w := httptest.NewRecorder()
	EditCal(w, cal)

	//create a request body
	calBody, err = json.Marshal(interfaces.CalendarItem{User: 100, EventID: 150, StartStr: "2023-04-19T10:00:00", EndStr: "2023-04-19T11:30:00", Title: "Film SWE Sprint 3"})
	if err != nil {
		log.Print("error encountered in marshal")
	}

	//create a http request and send it to the function
	cal = httptest.NewRequest(http.MethodPost, "/EditCal", bytes.NewBuffer(calBody))
	cal.Header.Set("Contest-type", "application/json")
	w = httptest.NewRecorder()
	EditCal(w, cal)

	//create a request body
	calBody, err = json.Marshal(interfaces.CalendarItem{User: 100, EventID: 17, StartStr: "2023-04-20T10:40:00", EndStr: "2023-04-20T11:30:00", Title: "Physics 2 Quiz"})
	if err != nil {
		log.Print("error encountered in marshal")
	}

	//create a http request and send it to the function
	cal = httptest.NewRequest(http.MethodPost, "/EditCal", bytes.NewBuffer(calBody))
	cal.Header.Set("Contest-type", "application/json")
	w = httptest.NewRecorder()
	EditCal(w, cal)

	//create a request body
	calBody, err = json.Marshal(interfaces.CalendarItem{User: 100, EventID: 56, StartStr: "2023-04-24T9:35:00", EndStr: "2023-04-24T10:25:00", Title: "CAP3032 Meeting"})
	if err != nil {
		log.Print("error encountered in marshal")
	}

	//create a http request and send it to the function
	cal = httptest.NewRequest(http.MethodPost, "/EditCal", bytes.NewBuffer(calBody))
	cal.Header.Set("Contest-type", "application/json")
	w = httptest.NewRecorder()
	EditCal(w, cal)

	calBody, err = json.Marshal(interfaces.CalendarItem{User: 100, EventID: 56, StartStr: "2023-04-24T9:35:00", EndStr: "2023-04-24T10:25:00", Title: "CAP3032 Meeting"})
	if err != nil {
		log.Print("error encountered in marshal")
	}

	//create a http request and send it to the function
	cal = httptest.NewRequest(http.MethodPost, "/CalDelete", bytes.NewBuffer(calBody))
	cal.Header.Set("Contest-type", "application/json")
	w = httptest.NewRecorder()
	DeleteCal(w, cal)

	calBody, err = json.Marshal(interfaces.CalendarItem{User: 100, EventID: 17, StartStr: "2023-04-20T10:40:00", EndStr: "2023-04-20T11:30:00", Title: "Physics 2 Quiz"})
	if err != nil {
		log.Print("error encountered in marshal")
	}

	//create a http request and send it to the function
	cal = httptest.NewRequest(http.MethodPost, "/CalDelete", bytes.NewBuffer(calBody))
	cal.Header.Set("Contest-type", "application/json")
	w = httptest.NewRecorder()
	DeleteCal(w, cal)

	calBody, err = json.Marshal(interfaces.CalendarItem{User: 100, EventID: 150, StartStr: "2023-04-19T10:00:00", EndStr: "2023-04-19T11:30:00", Title: "Film SWE Sprint 3"})
	if err != nil {
		log.Print("error encountered in marshal")
	}

	//create a http request and send it to the function
	cal = httptest.NewRequest(http.MethodPost, "/CalDelete", bytes.NewBuffer(calBody))
	cal.Header.Set("Contest-type", "application/json")
	w = httptest.NewRecorder()
	DeleteCal(w, cal)

	calBody, err = json.Marshal(interfaces.CalendarItem{User: 100})
	if err != nil {
		log.Print("error encountered in marshal")
	}

	//create a http request and send it to the function
	cal = httptest.NewRequest(http.MethodPost, "/CalStatus", bytes.NewBuffer(calBody))
	cal.Header.Set("Contest-type", "application/json")
	w = httptest.NewRecorder()
	CalStatus(w, cal)

	//set up the expected response
	res := w.Result()
	defer res.Body.Close()
	data, err := io.ReadAll(res.Body)
	bodyString := string(data)

	expectedString := "{\"items\":[{\"EventID\":1560,\"StartStr\":\"2023-04-23T11:00:00\",\"EndStr\":\"2023-04-23T12:00:00\",\"Title\":\"SWE Meeting\"}]}\n"

	//check to see if the expected matches what was returned from the tested functions
	if bodyString != expectedString {
		t.Errorf("error- failed response: %v", string(data))
	}

	calBody, err = json.Marshal(interfaces.CalendarItem{User: 100, EventID: 1560, StartStr: "2023-04-23T11:00:00", EndStr: "2023-04-23T12:00:00", Title: "SWE Meeting"})
	if err != nil {
		log.Print("error encountered in marshal")
	}

	//create a http request and send it to the function
	cal = httptest.NewRequest(http.MethodPost, "/CalDelete", bytes.NewBuffer(calBody))
	cal.Header.Set("Contest-type", "application/json")
	w = httptest.NewRecorder()
	DeleteCal(w, cal)

}

/*

func TestRequestFriend(t *testing.T) {
	database.InitTestDatabase()
	migrations.MigrateFriends()

	//---------------registering users to test requesting friends---------------//
	reqBody, err := json.Marshal(interfaces.User{Username: "JohnMark123", Name: "johanthan", Email: "johnsemail@email.com", Password: "JohN$pw0rd!"})
	if err != nil {
		log.Print("error encountered in marshal")
	}
	req := httptest.NewRequest(http.MethodPost, "/Register", bytes.NewBuffer(reqBody))
	req.Header.Set("Contest-type", "application/json")
	w := httptest.NewRecorder()
	RegisterFunc(w, req)

	reqBody, err = json.Marshal(interfaces.User{Username: "Bnel2802!", Name: "Brooke", Email: "brookenelson1@icloud.com", Password: "Din0$aur123"})
	if err != nil {
		log.Print("error encountered in marshal")
	}
	req = httptest.NewRequest(http.MethodPost, "/Register", bytes.NewBuffer(reqBody))
	req.Header.Set("Contest-type", "application/json")
	w = httptest.NewRecorder()
	RegisterFunc(w, req)

	reqBody, err = json.Marshal(interfaces.User{Username: "C9Feld27!", Name: "Caroline", Email: "carolinefeld@email.com", Password: "icedCh@i17"})
	if err != nil {
		log.Print("error encountered in marshal")
	}
	req = httptest.NewRequest(http.MethodPost, "/Register", bytes.NewBuffer(reqBody))
	req.Header.Set("Contest-type", "application/json")
	w = httptest.NewRecorder()
	RegisterFunc(w, req)

	//blocking
	blockBody, err := json.Marshal(interfaces.FriendRequest{Requester: "Bnel2802!", Reciever: "C9Feld27!"})
	if err != nil {
		log.Print("error encountered in marshal")
	}

	req = httptest.NewRequest(http.MethodPost, "/BlockFriend", bytes.NewBuffer(blockBody))
	req.Header.Set("Contest-type", "application/json")
	w = httptest.NewRecorder()
	BlockFriend(w, req)

	//blockBody, err = json.Marshal(interfaces.FriendRequest{Requester: "Bnel2802!", Reciever: "NotUs3r!nDB"})
	//if err != nil {
	//	log.Print("error encountered in marshal")
	//}

	//req = httptest.NewRequest(http.MethodPost, "/BlockFriend", bytes.NewBuffer(blockBody))
	//req.Header.Set("Contest-type", "application/json")
	//w = httptest.NewRecorder()
	//BlockFriend(w, req)

	//blockBody, err = json.Marshal(interfaces.FriendRequest{Requester: "F@keUser123", Reciever: "NotUs3r!nDB"})
	//if err != nil {
	//	log.Print("error encountered in marshal")
	//}

	//req = httptest.NewRequest(http.MethodPost, "/BlockFriend", bytes.NewBuffer(blockBody))
	//req.Header.Set("Contest-type", "application/json")
	//w = httptest.NewRecorder()
	//BlockFriend(w, req)

	//friend requests
	friendBody, err := json.Marshal(interfaces.FriendRequest{Requester: "Bnel2802!", Reciever: "C9Feld27!"})
	if err != nil {
		log.Print("error encountered in marshal")
	}

	req = httptest.NewRequest(http.MethodPost, "/RequestFriend", bytes.NewBuffer(friendBody))
	req.Header.Set("Contest-type", "application/json")
	w = httptest.NewRecorder()
	RequestFriend(w, req)

	friendBody, err = json.Marshal(interfaces.FriendRequest{Requester: "C9Feld27!", Reciever: "JohnMark123"})
	if err != nil {
		log.Print("error encountered in marshal")
	}

	req = httptest.NewRequest(http.MethodPost, "/RequestFriend", bytes.NewBuffer(friendBody))
	req.Header.Set("Contest-type", "application/json")
	w = httptest.NewRecorder()
	RequestFriend(w, req)

	//friendBody, err = json.Marshal(interfaces.FriendRequest{Requester: "Bnel2802!", Reciever: "F@keus3r"})
	//if err != nil {
	//	log.Print("error encountered in marshal")
	//}

	//req = httptest.NewRequest(http.MethodPost, "/RequestFriend", bytes.NewBuffer(friendBody))
	//req.Header.Set("Contest-type", "application/json")
	//w = httptest.NewRecorder()
	//RequestFriend(w, req)

	//friendBody, err = json.Marshal(interfaces.FriendRequest{Requester: "F@keUs3r2", Reciever: "F@keUs3r3"})
	//if err != nil {
	//	log.Print("error encountered in marshal")
	//}

	//req = httptest.NewRequest(http.MethodPost, "/RequestFriend", bytes.NewBuffer(friendBody))
	//req.Header.Set("Contest-type", "application/json")
	//w = httptest.NewRecorder()
	//RequestFriend(w, req)

	friendBody, err = json.Marshal(interfaces.FriendRequest{Requester: "C9Feld27!", Reciever: "JohnMark123"})
	if err != nil {
		log.Print("error encountered in marshal")
	}

	req = httptest.NewRequest(http.MethodPost, "/RequestFriend", bytes.NewBuffer(friendBody))
	req.Header.Set("Contest-type", "application/json")
	w = httptest.NewRecorder()
	RequestFriend(w, req)

	req = httptest.NewRequest(http.MethodPost, "/FriencStatus", bytes.NewBuffer(reqBody))
	req.Header.Set("Contest-type", "application/json")
	w = httptest.NewRecorder()
	CalStatus(w, req)

	//set up the expected response
	res := w.Result()
	defer res.Body.Close()
	data, err := io.ReadAll(res.Body)
	bodyString := string(data)

	expectedString := "lalala"

	if bodyString != expectedString {
		t.Errorf("error- failed response: %v", string(data))
	}

	//________________________________________________________________CLEARING  USERS FROM TABLE (dont have a function for clearing but will work if we create one)__________________________________________________________________________________/////
	reqBody, err = json.Marshal(interfaces.User{Username: "JohnMark123", Name: "johanthan", Email: "johnsemail@email.com", Password: "JohN$pw0rd!"})
	if err != nil {
		log.Print("error encountered in marshal")
	}

	//create a http request and send it to the function
	req = httptest.NewRequest(http.MethodDelete, "/Register", bytes.NewBuffer(reqBody))
	req.Header.Set("Content-type", "application/json")
	w = httptest.NewRecorder()
	RegisterFunc(w, req)

	reqBody, err = json.Marshal(interfaces.User{Username: "Bnel2802!", Name: "Brooke", Email: "brookenelson1@icloud.com", Password: "Din0$aur123"})
	if err != nil {
		log.Print("error encountered in marshal")
	}

	//create a http request and send it to the function
	req = httptest.NewRequest(http.MethodDelete, "/Register", bytes.NewBuffer(reqBody))
	req.Header.Set("Content-type", "application/json")
	w = httptest.NewRecorder()
	RegisterFunc(w, req)

	reqBody, err = json.Marshal(interfaces.User{Username: "C9Feld27!", Name: "Caroline", Email: "carolinefeld@email.com", Password: "icedCh@i17"})
	if err != nil {
		log.Print("error encountered in marshal")
	}

	//create a http request and send it to the function
	req = httptest.NewRequest(http.MethodDelete, "/Register", bytes.NewBuffer(reqBody))
	req.Header.Set("Content-type", "application/json")
	w = httptest.NewRecorder()
	RegisterFunc(w, req)
}
/*

//func TestAcceptFriend(t *testing.T) {}

//func TestFriendStat(t *testing.T) {
//	database.InitTestDatabase()
//	migrations.MigrateFriends()
//}
*/
