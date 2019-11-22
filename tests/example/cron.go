package main

import (
	"fmt"
	"time"

	"gopkg.in/robfig/cron.v2"
)

var Cron *cron.Cron

func main() {
	Cron = cron.New()
	Cron.Start()
	schedule, err := cron.Parse("3 * * * * *")
	if err != nil {
		fmt.Println(err)
	}
	Cron.Schedule(schedule, cron.FuncJob(func() {
		fmt.Println(time.Now())
	}))

	time.Sleep(500 * time.Second)
}
