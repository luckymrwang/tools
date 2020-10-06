package controllers

import (
	"encoding/json"
	"fmt"
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
