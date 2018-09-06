package main

import (
	"bufio"
	"database/sql"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
	"time"

	_ "github.com/lib/pq"
)

func main() {
	args := os.Args
	if len(args) < 2 {
		panic("请输入文件名")
	}
	file, err := os.Open(os.Args[1])
	defer file.Close()
	if err != nil {
		return
	}

	// Start reading from the file with a reader.
	reader := bufio.NewReader(file)
	var line string
	var str string
	rets := make([]string, 0)
	for {
		line, err = reader.ReadString('\n')
		if err != nil {
			break
		}
		if line == "\n" {
			continue
		}
		str += strings.Replace(line, "\n", " ", -1)
		if strings.Contains(line, ";") {
			rets = append(rets, strings.Trim(str, " "))
			str = ""
		}
	}

	if err != io.EOF {
		fmt.Printf(" > Failed!: %v\n", err)
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

	for _, sql := range rets {
		st := time.Now()
		rows, err := db.Query(sql)
		ed := time.Now()
		if err != nil {
			fmt.Println("err", err)
			continue
		}
		rows.Close()

		fmt.Println(fmt.Sprintf("%v\t%s", ed.Sub(st).Seconds(), sql))
		time.Sleep(1 * time.Second)
	}
}
