package main

import (
	"encoding/json"
	"fmt"
)

type Desc struct {
	Lang    string `json:"lang"`
	Content string `json:"content"`
	M       int    `json:"m,omitempty"`
}
type DescSlice struct {
	Desc []Desc `json:"body"`
}

func main() {
	app1 := `{"lang":"ch", "content":"1233456"}`
	var info1 Desc
	err := json.Unmarshal([]byte(app1), &info1)
	if err != nil {
		fmt.Printf("error is %v\n", err)
	} else {
		fmt.Printf("%v\n", info1)
	}

	app2 := `[{"lang":"ch01","content":"1233456"},{"lang":"ch02","content":"1233456"}]`
	var info2 []Desc
	err = json.Unmarshal([]byte(app2), &info2)
	if err != nil {
		fmt.Printf("error is %v\n", err)
	} else {
		fmt.Printf("%v\n", info2)
	}

	app3 := `{"body":[{"lang":"ch01","content":"1233456"},{"lang":"ch02","content":"1233456"}]}`
	info3 := DescSlice{}
	err = json.Unmarshal([]byte(app3), &info3)
	if err != nil {
		fmt.Println("error is %v\n", err)
	} else {
		fmt.Printf("%v\n", info3)
	}

	des := Desc{
		Lang:    "kaa",
		Content: "kkk",
	}
	data, _ := json.Marshal(des)
	fmt.Println(string(data), des.M)
}
