package main

import (
	"fmt"
	"os"
)

func main() {
	s := make([]string, 0)
	fmt.Println(s)
	for _, v := range s {
		fmt.Println("hello", v)
	}

	fileInfo, err := os.Stat("/Users/sino/Downloads/dhtest/2014售电量.xlsx")
	if err == nil {
		//文件大小
		fileSize := fileInfo.Size()
		fmt.Println(fileSize / 1024)
	} else {
		fmt.Println(err)
	}

}
