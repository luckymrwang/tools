package main

import (
	"encoding/json"
	"fmt"
)

type Student struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

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

	data = `"[{\"name\":\"hello\",\"age\":12},{\"name\":\"boy\",\"age\":12}]"`
	var stu []*Student
	err = json.Unmarshal([]byte(data), &stu)
	if err != nil {
		fmt.Println("eerrrr:", err)
	} else {
		for _, s := range stu {
			fmt.Println(s.Name, s.Age)
		}
	}

}
