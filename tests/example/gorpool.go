package main

import (
	"fmt"
	"github.com/yale8848/gorpool"
	"time"
)

func main() {

	// workerNum is worker number of goroutine pool ,one worker have one goroutine ,
	// jobNum is job number of job pool
	p := gorpool.NewPool(5, 10).
		Start()
	defer p.StopAll()
	for i := 0; i < 100; i++ {
		count := i
		p.AddJob(func() {
			time.Sleep(10 * time.Millisecond)
			fmt.Printf("%d\r\n", count)
		})

	}
	time.Sleep(2 * time.Second)
}
