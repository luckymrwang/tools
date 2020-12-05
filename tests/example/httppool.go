package main

import (
	"io"

	"io/ioutil"

	"net/http"

	"time"
)

func worker(client *http.Client) {
	resp, _ := client.Get("http://www.qq.com")
	io.Copy(ioutil.Discard, resp.Body)
	defer resp.Body.Close()
}

func worker2(client *http.Client) {
	resp, _ := client.Get("http://baidu.com")
	io.Copy(ioutil.Discard, resp.Body)
	defer resp.Body.Close()
}

func main() {
	tr := &http.Transport{
		MaxIdleConns:        1,
		MaxIdleConnsPerHost: 1,
	}
	client := &http.Client{
		Transport: tr,
	}
	worker(client)
	time.Sleep(time.Second * 1)
	worker(client)
	time.Sleep(time.Second * 1)
	worker2(client)
	time.Sleep(time.Second * 1)
	worker(client)
	time.Sleep(time.Second * 1)
}
