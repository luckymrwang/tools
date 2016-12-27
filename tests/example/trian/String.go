package main

import (
	"encoding/json"
	"fmt"
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
}
