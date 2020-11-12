package controllers

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

type Dog struct {
	Age int `json:",omitempty"`
}

func TestHandleExecShell(t *testing.T) {
	age := 0
	d := Dog{
		Age: age,
	}

	b, _ := json.Marshal(d)
	fmt.Println(string(b))
}

func TestGet(t *testing.T) {
	var f float64 = 1.23
	fmt.Printf("%0.f", f)
}

func TestResponseRecorder(t *testing.T) {
	handler := func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, "<html><body>Hello World!</body></html>")
	}

	req := httptest.NewRequest("GET", "http://example.com/foo", nil)
	w := httptest.NewRecorder()
	handler(w, req)

	resp := w.Result()
	body, _ := ioutil.ReadAll(resp.Body)

	fmt.Println(resp.StatusCode)
	fmt.Println(resp.Header.Get("Content-Type"))
	fmt.Println(string(body))

	// Output:
	// 200
	// text/html; charset=utf-8
	// <html><body>Hello World!</body></html>
}
