package main

import (
	"fmt"

	"github.com/json-iterator/go"
)

func main() {
	m := map[string]interface{}{
		"3": 3,
		"1": 1,
		"2": 2,
	}
	json := jsoniter.ConfigCompatibleWithStandardLibrary
	b, err := json.Marshal(m)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println(string(b))

	val := jsoniter.Get(b, "3").ValueType()
	fmt.Println(val)
}
