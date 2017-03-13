package main

import (
	"fmt"
	"time"
)

func main() {
	t := time.Now()
	fmt.Println(t)

	fmt.Println(t.Day())

	day := daysIn(t.Month(), t.Year())
	fmt.Println(day)

	fmt.Println(t.Add(-24 * time.Hour).Format("2006-01-02"))

	zone, _ := time.LoadLocation("Asia/Tokyo")
	fmt.Println(time.Now().In(zone))
}

func daysIn(m time.Month, year int) int {
	// This is equivalent to time.daysIn(m, year).
	return time.Date(year, m+1, 0, 0, 0, 0, 0, time.Local).Day()
}
