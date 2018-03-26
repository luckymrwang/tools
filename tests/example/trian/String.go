package main

import (
	"encoding/json"
	"fmt"
	"strings"
)

type App struct {
	AppType string `json:"app_type"`
	Add     string `json:"add"`
}

func main() {
	a := `{"app_type":"0", "add":""}`
	var app App
	err := json.Unmarshal([]byte(a), &app)
	if err != nil {
		fmt.Println("err", err)
	}
	fmt.Println("app", app)

	if app.AppType == "0" {
		fmt.Println("OK")
	} else {
		fmt.Println("NO")
	}

	var sil []string = []string{"a", "b", "c"}
	fmt.Println("'" + strings.Join(sil, "','") + "'")

	var ids []string
	for _, v := range ids {
		fmt.Println(v)
	}

	var s map[string]bool
	for k, _ := range s {
		fmt.Println(k)
	}
}
