package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	var subAppIds []string
	appIds := `["2005001","2005002","2005003"]`
	err := json.Unmarshal([]byte(appIds), &subAppIds)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(subAppIds)
	}

	var zones map[string]string
	data := `{"1":"2005001","2":"2005002","3":"2005003"}`
	err = json.Unmarshal([]byte(data), &zones)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(zones)
	}
}
