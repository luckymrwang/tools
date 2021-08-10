package channel

import (
	"context"
	"encoding/json"
	"fmt"
	"reflect"
	"testing"
	"time"
)

var rootMap = map[string]string{
	"a": "b",
	"c": "d",
}

type MMM struct {
	ABC string `json:"abc"`
	D   string `json:"d"`
}

func TestMap(t *testing.T) {
	m1 := getM()
	m2 := getM()
	m2["e"] = "f"
	if reflect.DeepEqual(m1, m2) {
		fmt.Println("equal")
	} else {
		fmt.Println("not equal")
	}
	m := MMM{
		ABC: "xxxx",
		D:   "12",
	}
	mstr, _ := json.Marshal(m)
	mq := make(map[string]string)
	json.Unmarshal([]byte(mstr), &mq)
	fmt.Println(mq)
}

func getM() map[string]string {
	m := make(map[string]string)
	m = rootMap
	return m
}

var (
	// singleton
	bee *beehive
)

func TestChx(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	bee = &beehive{
		ctx:      ctx,
		cancel:   cancel,
		channels: make(map[string]chan string),
	}
	channel := make(chan string, 10)
	bee.channels["hi"] = channel
	goch()

	time.Sleep(10 * time.Second)
	bee.channels["hi"] <- "ooookkk1"
	bee.channels["hi"] <- "ooookkk2"
	bee.channels["hi"] <- "ooookkk3"
	bee.channels["hi"] <- "ooookkk4"
	bee.channels["hi"] <- "ooookkk5"
	bee.channels["hi"] <- "ooookkk6"
	bee.channels["hi"] <- "ooookkk7"
	bee.channels["hi"] <- "ooookkk8"
	bee.channels["hi"] <- "ooookkk9"
	bee.channels["hi"] <- "ooookkk10"
	bee.channels["hi"] <- "ooookkk11"
	bee.channels["hi"] <- "ooookkk12"
	bee.channels["hi"] <- "ooookkk13"
	time.Sleep(10 * time.Minute)
}

func goch() {
	go func() {
		for {
			select {
			case <-Done():
				fmt.Println("MetaManager mainloop stop")
				return
			default:
				fmt.Println("xxxx")
			}
			if msg, err := bee.Receive("hi"); err == nil {
				fmt.Printf("get a message %+v", msg)
				time.Sleep(5 * time.Second)
			} else {
				fmt.Printf("get a message %+v: %v", msg, err)
			}
		}
	}()
}

type beehive struct {
	cancel   context.CancelFunc
	ctx      context.Context
	channels map[string]chan string
}

func Done() <-chan struct{} {
	return bee.ctx.Done()
}

func (ctx *beehive) Receive(module string) (string, error) {
	if channel := ctx.getChannel(module); channel != nil {
		content := <-channel
		return content, nil
	}

	fmt.Printf("Failed to get channel for module:%s when receive message", module)
	return "", fmt.Errorf("failed to get channel for module(%s)", module)
}

func (ctx *beehive) getChannel(module string) chan string {
	if _, exist := ctx.channels[module]; exist {
		return ctx.channels[module]
	}
	fmt.Printf("Failed to get channel, type:%s", module)
	return nil
}
