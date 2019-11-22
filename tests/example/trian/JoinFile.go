package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
)

func main() {
	fii, err := os.OpenFile("./all.jpg", os.O_CREATE|os.O_WRONLY|os.O_APPEND, os.ModePerm)
	if err != nil {
		panic(err)
		return
	}
	defer fii.Close()

	num := 2
	for i := 0; i <= int(num); i++ {
		f, err := os.OpenFile("./somebigfile_"+strconv.Itoa(int(i)), os.O_RDONLY, os.ModePerm)
		if err != nil {
			fmt.Println(err)
			return
		}
		b, err := ioutil.ReadAll(f)
		if err != nil {
			fmt.Println(err)
			return
		}
		fii.Write(b)
		f.Close()
	}
}
