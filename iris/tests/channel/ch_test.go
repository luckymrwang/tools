package channel

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"testing"
	"time"
)

var errSignal chan struct{}

type MyHandler map[string]interface{}

func (self MyHandler) list(w http.ResponseWriter, req *http.Request) {
	for item, price := range self {
		fmt.Fprintf(w, "%s: %s\n", item, price)
	}
}
func (self MyHandler) price(w http.ResponseWriter, req *http.Request) {
	item := req.URL.Query().Get("item")
	price, ok := self[item]
	if !ok {
		w.WriteHeader(http.StatusNotFound) // 404
		fmt.Fprintf(w, "no such item: %q\n", item)
		return
	}
	fmt.Fprintf(w, "%s\n", price)
}

func TestCh(t *testing.T) {
	if err := start(setup()); err != nil {
		fmt.Println("exit")
		os.Exit(1)
	}

	handler := MyHandler{"shoes": 50, "socks": 5}
	http.HandleFunc("/list", handler.list)
	http.HandleFunc("/price", handler.price)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func (self MyHandler) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	switch req.URL.Path {
	case "/list":
		for item, price := range self {
			fmt.Fprintf(w, "%s: %s\n", item, price)
		}
	case "/price":
		item := req.URL.Query().Get("item")
		price, ok := self[item]
		if !ok {
			w.WriteHeader(http.StatusNotFound) // 404
			fmt.Fprintf(w, "no such item: %q\n", item)
			return
		}
		fmt.Fprintf(w, "%s\n", price)
	default:
		w.WriteHeader(http.StatusNotFound) // 404
		fmt.Fprintf(w, "no such page: %s\n", req.URL)
	}
}

func TestHttp(t *testing.T) {
	handler := MyHandler{"shoes": 50, "socks": 5}
	log.Fatal(http.ListenAndServe("localhost:8000", handler))
}

func start(stop <-chan struct{}) error {
	fmt.Println("start....")
	select {
	case <-stop:
		fmt.Println("stop......")
		return nil
	case <-errSignal:
		fmt.Println("errSignal...")
		// Error starting a controller
		return nil
	}
}

func setup() (stopCh <-chan struct{}) {
	stop := make(chan struct{})
	go func() {
		time.Sleep(10 * time.Second)
		fmt.Println("time over")
		close(stop)
	}()

	return stop
}
