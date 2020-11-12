package services

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"time"
	"tools/iris/common"

	"github.com/gorilla/websocket"
	"github.com/kataras/iris/v12"
)

// https://github.com/gorilla/websocket/blob/master/server.go#L230
var WebSocketReadBufSize = 4096

// WebSocketWriteBufSize is a parameter that is used for WebSocket Upgrader
// https://github.com/gorilla/websocket/blob/master/server.go#L230
var WebSocketWriteBufSize = 4096

var (
	HeartbeatDelay  = 25 * time.Second
	DisconnectDelay = 5 * time.Second
	ResponseLimit   = 128 * 1024
)

type WebSocketService struct {
	Ctx iris.Context
}

func GetWebSocketService(ctx iris.Context) *WebSocketService {
	return &WebSocketService{Ctx: ctx}
}

func (h *WebSocketService) Upgrade(rw http.ResponseWriter, req *http.Request, sessID string) {
	conn, err := websocket.Upgrade(rw, req, nil, WebSocketReadBufSize, WebSocketWriteBufSize)
	if _, ok := err.(websocket.HandshakeError); ok {
		http.Error(rw, `Can "Upgrade" only to "WebSocket".`, http.StatusBadRequest)
		return
	} else if err != nil {
		rw.WriteHeader(http.StatusInternalServerError)
		return
	}

	sess := newSession(sessID, DisconnectDelay, HeartbeatDelay)
	go common.HandleTerminalSession(sess)

	receiver := newWsReceiver(conn)
	sess.attachReceiver(receiver)
	readCloseCh := make(chan struct{})
	go func() {
		var d []string
		for {
			err := conn.ReadJSON(&d)
			if err != nil {
				fmt.Println(err)
				close(readCloseCh)
				return
			}
			sess.accept(d...)
		}
	}()

	select {
	case <-readCloseCh:
	case <-receiver.doneNotify():
	}
	sess.close()
	conn.Close()
}

type wsReceiver struct {
	conn    *websocket.Conn
	closeCh chan struct{}
}

func newWsReceiver(conn *websocket.Conn) *wsReceiver {
	return &wsReceiver{
		conn:    conn,
		closeCh: make(chan struct{}),
	}
}

func (w *wsReceiver) sendBulk(messages ...string) {
	if len(messages) > 0 {
		w.sendFrame(fmt.Sprintf("a[%s]", strings.Join(transform(messages, quote), ",")))
	}
}

func (w *wsReceiver) sendFrame(frame string) {
	if err := w.conn.WriteMessage(websocket.TextMessage, []byte(frame)); err != nil {
		w.close()
	}
}

func (w *wsReceiver) close() {
	select {
	case <-w.closeCh: // already closed
	default:
		close(w.closeCh)
	}
}
func (w *wsReceiver) canSend() bool {
	select {
	case <-w.closeCh: // already closed
		return false
	default:
		return true
	}
}
func (w *wsReceiver) doneNotify() <-chan struct{}        { return w.closeCh }
func (w *wsReceiver) interruptedNotify() <-chan struct{} { return nil }

func quote(in string) string {
	quoted, _ := json.Marshal(in)
	return string(quoted)
}

func transform(values []string, transformFn func(string) string) []string {
	ret := make([]string, len(values))
	for i, msg := range values {
		ret[i] = transformFn(msg)
	}
	return ret
}
