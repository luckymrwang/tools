//客户端代码，演示了发送一次数据就断开连接的

package main

import (
	"fmt"

	"net"

	"os"

	"time"
)

func main() {

	server := "127.0.0.1:9988"

	for i := 0; i < 100; i++ {

		tcpAddr, err := net.ResolveTCPAddr("tcp4", server)

		if err != nil {

			fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())

			os.Exit(1)

		}

		conn, err := net.DialTCP("tcp", nil, tcpAddr)

		if err != nil {

			fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())

			os.Exit(1)

		}

		words := "{\"Id\":1,\"Name\":\"golang\",\"Message\":\"message\"}"

		conn.Write([]byte(words))

		conn.Close()

	}

	for {

		time.Sleep(1 * 1e9)

	}

}
