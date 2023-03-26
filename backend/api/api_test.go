// to perform tests:
//cd into file you want to test, and type go test
//if multiple functions in test file, do go test -v instead of go test

package api

import (
	"encoding/json"
	"main/database"
	"main/interfaces"

	"bytes"
	"io"
	"net/http"
	"testing"

	//"github.com/stretchr/testify/assert"
	"net/http/httptest"
)

// https://golang.cafe/blog/golang-httptest-example.html

// TestReadBody has been started
//func TestReadBody(t *testing.T) {
//req := httptest.NewRequest(http.MethodGet, nil)
//actualBody := readBody(req)

//}

// TestApiResponse hasnt been started
//func TestApiResponse(t *testing.T) {}

// TestLoginFunc does not run
//func TestLoginFunc(t *testing.T) {
//	user := interfaces.User{
//		Username: "testUser",
//		Password: "testPass",
//	}

//	writer := makeRequest("POST", "/auth/login", user, false)

//	assert.Equal(t, http.StatusOK, writer.Code)

//	var response map[string]string
//	json.Unmarshal(writer.Body.Bytes(), &response)
//	_, exists := response["jwt"]
//	assert.Equal(t, true, exists)
//}

// TestRegisterFunc hasnt been started
//func TestRegisterFunc(t *testing.T) {}

// TestGetUserFunc hasnt been started
//func TestGetUserFunc(t *testing.T) {}

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

func TestEditToDo(t *testing.T) {

	database.InitDatabase()

	reqBody, err := json.Marshal(interfaces.TodoReq{User: 1000, Description: "buy apples"})

	req := httptest.NewRequest(http.MethodPost, "/EditToDo", bytes.NewBuffer(reqBody))
	req.Header.Set("Content-type", "application/json")
	w := httptest.NewRecorder()
	//EditToDo(w, req)

	reqBody, err = json.Marshal(interfaces.UserID{User: 1000})

	req = httptest.NewRequest(http.MethodPost, "/ToDoStatus", bytes.NewBuffer(reqBody))
	req.Header.Set("Content-type", "application/json")
	w = httptest.NewRecorder()
	ToDoStatus(w, req)

	res := w.Result()
	defer res.Body.Close()
	data, err := io.ReadAll(res.Body)
	//body, err := json.Marshal(&data)
	bodyString := string(data)

	type resp struct {
		Complete   []string
		Incomplete []string
		Percentage float64
	}

	arg1 := []string{"buy apples", "buy apples"}
	arg2 := []string{}

	expected, err := json.Marshal(resp{Complete: arg2, Incomplete: arg1, Percentage: 0})
	expectedString := string(expected)
	expectedString += "\n"

	if err != nil {
		t.Errorf("expected error to be nil got %v", err)
	}
	if bodyString != expectedString {
		t.Errorf("error- failed response: %v", string(data))
	}
}

// func TestEditToDo(t *testing.T) {

// 	reqBody, err := json.Marshal(interfaces.TodoReq{User: 1000, Description: "buy apples"})

// 	req := httptest.NewRequest(http.MethodPost, "/EditToDo", bytes.NewBuffer(reqBody))
// 	req.Header.Set("Contest-type", "application/json")
// 	w := httptest.NewRecorder()
// 	EditToDo(w, req)

// 	reqBody, err = json.Marshal(interfaces.TodoReq{User: 1000, Description: "buy oranges"})

// 	req = httptest.NewRequest(http.MethodPost, "/EditToDo", bytes.NewBuffer(reqBody))
// 	req.Header.Set("Contest-type", "application/json")
// 	w = httptest.NewRecorder()
// 	EditToDo(w, req)

// 	reqBody, err = json.Marshal(interfaces.TodoReq{User: 1000, Description: "buy bananas"})

// 	req = httptest.NewRequest(http.MethodPost, "/EditToDo", bytes.NewBuffer(reqBody))
// 	req.Header.Set("Contest-type", "application/json")
// 	w = httptest.NewRecorder()
// 	EditToDo(w, req)

// 	reqBody, err = json.Marshal(interfaces.TodoReq{User: 1000, Description: "buy grapes"})

// 	req = httptest.NewRequest(http.MethodPost, "/EditToDo", bytes.NewBuffer(reqBody))
// 	req.Header.Set("Contest-type", "application/json")
// 	w = httptest.NewRecorder()
// 	EditToDo(w, req)

// 	reqBody, err = json.Marshal(interfaces.TodoReq{User: 1000, Description: "buy apples"})

// 	req = httptest.NewRequest(http.MethodPut, "/EditToDo", bytes.NewBuffer(reqBody))
// 	req.Header.Set("Contest-type", "application/json")
// 	w = httptest.NewRecorder()
// 	EditToDo(w, req)

// 	reqBody, err = json.Marshal(interfaces.TodoReq{User: 1000, Description: "buy grapes"})

// 	req = httptest.NewRequest(http.MethodDelete, "/EditToDo", bytes.NewBuffer(reqBody))
// 	req.Header.Set("Contest-type", "application/json")
// 	w = httptest.NewRecorder()
// 	EditToDo(w, req)

// 	reqBody, err = json.Marshal(interfaces.UserID{User: 1000})

// 	req = httptest.NewRequest(http.MethodPut, "/ToDoStatus", bytes.NewBuffer(reqBody))
// 	req.Header.Set("Contest-type", "application/json")
// 	w = httptest.NewRecorder()
// 	ToDoStatus(w, req)

// 	res := w.Result()
// 	defer res.Body.Close()
// 	data, err := io.ReadAll(res.Body)
// 	body, err := json.Marshal(&data)
// 	bodyString := string(body)

// 	type resp struct {
// 		Incomplete []string
// 		Complete   []string
// 		Percentage float64
// 	}

// 	arg1 := []string{"buy apples"}
// 	arg2 := []string{"buy oranges", "buy bananas"}

// 	expected, err := json.Marshal(resp{Incomplete: arg2, Complete: arg1, Percentage: 33.33})
// 	expectedString := string(expected)

// 	if err != nil {
// 		t.Errorf("expected error to be nil got %v", err)
// 	}
// 	if bodyString != expectedString {
// 		t.Errorf("error- failed response: %v", string(data))
// 	}
// }
