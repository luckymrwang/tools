package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"sync"
	"time"

	"github.com/gorilla/websocket"
)

type WebSocketController struct {
	BaseController
}

type Msg struct {
	Msgtype   string
	From      string
	Content   string
	Timestamp string
}

var (
	wsPool    sync.Map
	broadcast = make(chan string)
)

func init() {
	go wsHeartBeat()
}

func (c *WebSocketController) Get() {
	id := c.Ctx.Input.Param(":id")
	fmt.Println(id)
	c.TplName = fmt.Sprintf("ws%s.html", id)
	c.Render()
}

func (c *WebSocketController) Receive() {
	auth := c.GetString("auth")
	go func(auth string) {
		broadcast <- auth
	}(auth)
}

func (c *WebSocketController) Join() {
	auth := c.GetString("auth")
	fmt.Println("auth:", auth)
	ws, err := websocket.Upgrade(c.Ctx.ResponseWriter, c.Ctx.Request, nil, 1024, 1024)
	if _, ok := err.(websocket.HandshakeError); ok {
		fmt.Println("Websocket handshake Error:", err.Error())
		http.Error(c.Ctx.ResponseWriter, "Not a websocket handshake", 400)
		return
	} else if err != nil {
		fmt.Println("Cannot setup WebSocket connection:", err)
		return
	}
	if auth == "" {
		ws.WriteMessage(websocket.TextMessage, []byte("Auth Failed"))
		ws.Close()
		return
	}

	wsPool.Store(auth, ws)
	go func(auth string, _ws *websocket.Conn) {
		for {
			_, payload, err := _ws.ReadMessage()
			if err != nil {
				fmt.Println("read message fail:", err)
				return
			}
			var msg Msg
			err = json.Unmarshal(payload, &msg)
			if err != nil {
				fmt.Println("json unmarshal failed: ", err)
				wsClose(auth)
				return
			}
			msg.From = "服务端回复"
			msg.Timestamp = time.Now().Format("2006-01-02 15:04:05")
			_bytes, err := json.Marshal(msg)
			if err != nil {
				fmt.Println("json marshal error:", msg, err)
				wsClose(auth)
				return
			}
			err = _ws.WriteMessage(websocket.TextMessage, _bytes)
			if err != nil {
				fmt.Println("WebSocket Notify Fail:", err)
				wsClose(auth)
				return
			}
		}
	}(auth, ws)
}

func send(auth string) error {
	_ws, ok := wsPool.Load(auth)
	if !ok {
		fmt.Println("WebSocket Notify Get Conn Fail:", auth)
		return fmt.Errorf("WebSocket Notify Get Conn Fail:%v", auth)
	}
	msg := map[string]interface{}{
		"aaa": "xxx",
	}
	err := _ws.(*websocket.Conn).WriteMessage(websocket.TextMessage, []byte(JSONEncode(msg)))
	if err != nil {
		fmt.Println("WebSocket Notify Fail:", err)
		return fmt.Errorf("WebSocket Notify Fail:%v", err)
	}

	return nil
}

func wsClose(auth string) {
	if ws, ok := wsPool.Load(auth); ok {
		ws.(*websocket.Conn).Close()
		wsPool.Delete(auth)
	}
}

func wsHeartBeat() {
	t := time.Tick(10 * time.Second)
	for {
		<-t
		wsPool.Range(func(auth, ws interface{}) bool {
			msg := Msg{
				Msgtype:   "text",
				Content:   "Hi",
				From:      "服务端主动推送",
				Timestamp: time.Now().Format("2006-01-02 15:04:05"),
			}
			_bytes, err := json.Marshal(msg)
			if err != nil {
				fmt.Println("json marshal error:", msg, err)
				return true
			}
			if err := ws.(*websocket.Conn).WriteMessage(websocket.TextMessage, _bytes); err != nil {
				fmt.Println("wsHeartBeat Error:", err)
				wsClose(auth.(string))
			}
			return true
		})
	}
}

//JSONEncode Json格式编码
func JSONEncode(value interface{}) string {
	_bytes, err := json.Marshal(value)
	if err != nil {
		fmt.Println("JSONEncode Error:", value, err)
		return ""
	}
	return string(_bytes)
}

func JSONDecode(str string) (p map[string]interface{}) {
	_bytes := []byte(str)
	if len(_bytes) <= 0 {
		return p
	}
	err := json.Unmarshal(_bytes, &p)
	if err != nil {
		fmt.Println(fmt.Sprintf("JSONDecode Error:%v,Str:%v", err.Error(), str))
	}
	return
}
