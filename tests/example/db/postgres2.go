package main

import (
	"database/sql"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"time"

	_ "github.com/lib/pq"
)

func main() {
	args := os.Args
	if len(args) < 2 {
		panic("请输入文件夹")
	}

	connStr := "postgres://postgres:123456@106.75.27.238/tpc?sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	defer db.Close()
	if err != nil {
		log.Fatal(err)
	}
	err = db.Ping()
	if err != nil {
		panic(err)
	}

	files, _ := ioutil.ReadDir(os.Args[1])
	for _, f := range files {
		sql := getFileContent(os.Args[1] + f.Name())
		if sql == "" {
			fmt.Println("err:sql is empty" + "\t" + f.Name())
			continue
		}

		st := time.Now()
		rows, err := db.Query(sql)
		ed := time.Now()
		if err != nil {
			fmt.Println(fmt.Sprintf("err:%s\t%s\t\"%s\"", err.Error(), f.Name(), sql))
			continue
		}
		rows.Close()

		fmt.Println(fmt.Sprintf("%v\t%s\t\"%s\"", ed.Sub(st).Seconds(), f.Name(), sql))
	}
}

func getFileContent(f string) string {
	file, err := os.Open(f)
	defer file.Close()
	if err != nil {
		return ""
	}
	dat, err := ioutil.ReadFile(f)
	if err != nil {
		return ""
	}
	return string(dat)
}
