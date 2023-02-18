package api

import (
	"testing"
)

//curl.exe -v -X GET http://localhost:3000/home-page

func TestUpperCaseHandler(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/home-page", nil)
	w := httptest.NewRecorder()
	GoHome(w, req)
	res := w.Result()
	defer res.Body.Close()
	data, err := ioutil.ReadAll(res.Body)
	if err != nil {
		t.Errorf("expected error to be nil got %v", err)
	}
	if string(data) != "ABC" {
		t.Errorf("expected: Welcome to the home page, get request recieved. Got: %v", string(data))
	}
}
