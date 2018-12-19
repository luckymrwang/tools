package main

import (
	"fmt"
	"github.com/hpcloud/tail"
)

func main() {
	t, err := tail.TailFile("/Users/sino/Downloads/1_1539779451004.csv", tail.Config{MaxLineSize: 3})
	if err != nil {
		return
	}
	for line := range t.Lines {
		fmt.Println(line.Text)
	}
}
