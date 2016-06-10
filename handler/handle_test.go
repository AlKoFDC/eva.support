package handler_test

import (
	. "github.com/AlKoFDC/eva.support/handler"
	"testing"
)

const (
	testID       = "test"
	helloMessage = `{"type": "hello"}`
)

func TestShouldCallWebsocketClose(t *testing.T) {
	callWithTestWebsocket(func(ws WebSocket) {
		handler := SlackMessageHandler{
			ID: testID,
			WS: ws,
		}
		handler.Close()
		if !closeCalled {
			t.Errorf("Expected close to be called, but it was not.")
		}
	},
		[]string{helloMessage},
	)
}
