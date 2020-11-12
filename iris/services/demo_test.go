package services

import (
	"fmt"
	"log"
	"strconv"
	"testing"
	"unicode/utf8"
	"encoding/json"
)

func TestEcho(t *testing.T) {
	dService := GetDemoService(nil)
	dService.Echo()
}

func TestGet(t *testing.T) {
	s := "xx1x王"
	fmt.Println(utf8.RuneCountInString(s))
	fmt.Println(len(s))

	_, err := strconv.ParseInt("10000000000", 10, 64)
	if err != nil {
		t.Error(err)
		return
	}
}

type AS struct {
	AA string
	BB string
}

func TestMap(t *testing.T) {
	var a []string
	data := map[string]interface{}{
		"a": a,
	}
	buf, err := json.Marshal(&data)
	log.Println(string(buf), err)
	data = map[string]interface{}{
		"b": make([]string, 0),
	}
	buf, err = json.Marshal(&data)
	log.Println(string(buf), err)
}

func TestAS(t *testing.T)  {
	as := new(AS)
	fmt.Println(*as)
	ass := AS{}
	fmt.Println(ass)

	f := float64(30.33)
	fmt.Printf("%.f", f)
}
