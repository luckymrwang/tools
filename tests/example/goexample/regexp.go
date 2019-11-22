package main

import (
	"fmt"
	"regexp"
)

func main() {
	word1 := "a2_2*d."
	if Check(word1) {
		fmt.Println(word1, "true")
	} else {
		fmt.Println(word1, "false")
	}

	word2 := "a03c2_-."
	if Check(word2) {
		fmt.Println(word2, "true")
	} else {
		fmt.Println(word2, "false")
	}

	for key, val := range word2 {
		fmt.Println(key, val)
	}
}

func Check(word string) (b bool) {
	if ok, _ := regexp.MatchString(`^[\w_-\.]{1,}$`, word); !ok {
		return false
	}
	return true
}
