package main

import (
	"database/sql"
	"fmt"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db, err := sql.Open("mysql", "root:123456@tcp(106.75.27.238:3307)/sample?charset=utf8")
	checkErr(err)
	defer db.Close()

	err = db.Ping()
	checkErr(err)

	for i := 0; i < 10; i++ {
		var name string

		t := time.Now()
		err = db.QueryRow("select name from movie_copy where id = ?", 285944).Scan(&name)
		if err != nil {
			checkErr(err)
		}
		fmt.Println(name, "time:", time.Now().Sub(t))

		time.Sleep(2 * time.Second)
	}
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
