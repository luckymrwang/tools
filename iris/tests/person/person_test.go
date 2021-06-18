package person

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

var personResponse = []Person{
	{
		Name:    "wahaha",
		Address: "shanghai",
		Age:     20,
	},
	{
		Name:    "lebaishi",
		Address: "shanghai",
		Age:     10,
	},
}

var personResponseBytes, _ = json.Marshal(personResponse)

func TestPublishWrongResponseStatus(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write(personResponseBytes)
		if r.Method != "GET" {
			t.Errorf("Expected 'GET' request, got '%s'", r.Method)
		}
		if r.URL.EscapedPath() != "/person" {
			t.Errorf("Expected request to '/person', got '%s'", r.URL.EscapedPath())
		}
		r.ParseForm()
		topic := r.Form.Get("addr")
		if topic != "shanghai" {
			t.Errorf("Expected request to have 'addr=shanghai', got: '%s'", topic)
		}
	}))

	defer ts.Close()
	api := ts.URL
	fmt.Println("url:", api)
	resp, _ := GetInfo(api)

	fmt.Println("reps:", resp)
}

type intee interface {
	Method()
	Json(a string) string
}

func TestIntee(t *testing.T) {
	var tee intee
	fmt.Println(tee)
	fmt.Println("xx")

	encodeString := "CjVmZmNhMDY0NjBmOTk0YjcwYjJiOWUyNDM3ZWM4N2YwOS4zNy4xNjIwNzIzODg2ODM1MDAwMRI1ZmZjYTA2NDYwZjk5NGI3MGIyYjllMjQzN2VjODdmMDkuMzcuMTYyMDcyMzg4NjgzNTAwMDAaVxD///////////8BGPPVr9WVLyD21a/VlS8yAS9IA1ABYiEKA3VybBIaaHR0cDovLzEwLjQ4LjUxLjEzNTozMTk1Ni9iEgoLaHR0cC5tZXRob2QSA0dFVCIkZmFmNGQ1ZjEtZWRlYS00MzQ2LWI5MDgtNGEzNTFjMDIwMzM3Ki4zMmE3ZWQ4ODMxYTA0NzY0YjVmZjYxOWViYzJjYzkyZEAxMC4yMzMuMTIxLjYx"
	decodeBytes, err := base64.StdEncoding.DecodeString(encodeString)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(decodeBytes))

	base64er := base64.RawURLEncoding

	decoder_buf, _ := base64er.DecodeString(encodeString)
	fmt.Println(string(decoder_buf))
}

func TestError(t *testing.T) {
	err := errfmt()
	fmt.Println(err)
}

func errfmt() error {
	err := fmt.Errorf("88")
	return fmt.Errorf("err: %s", err)
}
