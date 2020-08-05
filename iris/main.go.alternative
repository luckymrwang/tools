package main

import (
	"fmt" //只是一个可选的助手
	"io"
	"time" //展示延迟

	"github.com/kataras/iris"
)

func main() {
	app := iris.New()
	app.Get("/", func(ctx iris.Context) {
		ctx.ContentType("text/html")
		ctx.Header("Transfer-Encoding", "chunked")
		i := 0
		ints := []int{1, 2, 3, 5, 7, 9, 11, 13, 15, 17, 23, 29}
		//以块的形式发送响应，并在每个块之间等待半秒钟
		ctx.StreamWriter(func(w io.Writer) bool {
			fmt.Fprintf(w, "Message number %d<br>", ints[i])
			time.Sleep(500 * time.Millisecond) // simulate delay.
			if i == len(ints)-1 {
				return false //关闭并刷新
			}
			i++
			return true //继续写入数据
		})
	})

	type messageNumber struct {
		Number int `json:"number"`
	}

	app.Get("/alternative", func(ctx iris.Context) {
		ctx.ContentType("application/json")
		ctx.Header("Transfer-Encoding", "chunked")
		i := 0
		ints := []int{1, 2, 3, 5, 7, 9, 11, 13, 15, 17, 23, 29}
		//以块的形式发送响应，并在每个块之间等待半秒钟。
		for {
			ctx.JSON(messageNumber{Number: ints[i]})
			ctx.WriteString("\n")
			time.Sleep(500 * time.Millisecond) // simulate delay.
			if i == len(ints)-1 {
				break
			}
			i++
			ctx.ResponseWriter().Flush()
		}
	})
	app.Run(iris.Addr(":8080"))
}
