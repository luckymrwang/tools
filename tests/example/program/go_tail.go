package main

import (
	"fmt"
	"io"

	"github.com/papertrail/go-tail/follower"
)

func main() {
	t, err := follower.New("/Users/sino/Downloads/1_1539779451004.csv", follower.Config{
		Whence: io.SeekEnd,
		Offset: 0,
	})

	if err != nil {
		return
	}

	for line := range t.Lines() {
		fmt.Println(line)
	}
}
