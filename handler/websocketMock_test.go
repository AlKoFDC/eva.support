package handler_test

import (
	"encoding/json"
	"fmt"
	. "github.com/AlKoFDC/eva.support/handler"
)

type TestWebSocket struct {
	responses []string
}

var callCount int

func (ws TestWebSocket) ReadJSON(v interface{}) error {
	defer func() { callCount++ }()
	if len(ws.responses) <= callCount {
		return fmt.Errorf("no more messages")
	}
	return json.Unmarshal([]byte(ws.responses[callCount]), v)
}

var messageSent string

func (ws TestWebSocket) WriteJSON(v interface{}) error {
	bytesSent, err := json.Marshal(v)
	messageSent = string(bytesSent)
	return err
}

var closeCalled bool

func (ws TestWebSocket) Close() error {
	closeCalled = true
	return nil
}

func callWithTestWebsocket(funcToTest func(WebSocket), responses []string) {
	websocket := TestWebSocket{responses: responses}
	callCount = 0
	closeCalled = false
	messageSent = ""
	funcToTest(websocket)
}
