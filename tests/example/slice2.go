package main

import (
	"fmt"
	"strconv"
	"time"
)

func main() {
	t := time.Now()
	retSum := map[string]int64{
		"retention_1d":  0,
		"retention_2d":  0,
		"retention_3d":  0,
		"retention_4d":  0,
		"retention_5d":  0,
		"retention_6d":  0,
		"retention_7d":  0,
		"retention_14d": 0,
		"retention_30d": 0,
		"retention_60d": 0,
		"retention_90d": 0,
	}

	for k, _ := range retSum {
		dayStr := k[10 : len(k)-1]
		day, _ := strconv.ParseInt(dayStr, 10, 32)

		fmt.Println(day, t.AddDate(0, 0, -int(day)).Format("2006-01-02"))
	}
}
