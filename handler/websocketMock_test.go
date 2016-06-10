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
var closeCalled bool

func (ws TestWebSocket) ReadJSON(v interface{}) error {
	defer func() { callCount++ }()
	if len(ws.responses) <= callCount {
		return fmt.Errorf("no more messages")
	}
	return json.Unmarshal([]byte(ws.responses[callCount]), v)
}

func (ws TestWebSocket) Close() error {
	closeCalled = true
	return nil
}

func callWithTestWebsocket(funcToTest func(WebSocket), responses []string) {
	websocket := TestWebSocket{responses: responses}
	callCount = 0
	closeCalled = false
	funcToTest(websocket)
}
