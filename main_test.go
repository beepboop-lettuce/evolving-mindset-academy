package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHandler(t *testing.T) {
	//form new http request. this request will be passed to handler
	//first arg= method, second arg= route, third arg= request body
	req, err := http.NewRequest("GET", "", nil)

	// if error in forming request, fail and stop test
	if err != nil {
		t.Fatal(err)
	}

	//use Go's httptest library to creat http recorder. will act as target of http request
	// sort of like a mini browser, which will accept result of http request we make
	recorder := httptest.NewRecorder()

	//creat http handler from handler func. "handler" = handler func defined in main.go file (the one being tested)
	hf := http.HandlerFunc(handler)

	//serve http request to recorder. This line actually executes hander being tested
	hf.ServeHTTP(recorder, req)

	//check that the status code = expected
	if status := recorder.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v\n", status, http.StatusOK)
	}

	//check that the response body = expected
	expected := `Hello World!`
	actual := recorder.Body.String()
	if actual != expected {
		t.Errorf("handler returned unexpected body: got %v want %v\n", actual, expected)
	}

}
