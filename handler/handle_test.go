package handler_test

import (
	. "github.com/AlKoFDC/eva.support/handler"
	"github.com/AlKoFDC/eva.support/message"
	"testing"
)

const (
	testID       = "test"
	helloMessage = `{"type": "hello"}`
)

func TestShouldCallWebsocketClose(t *testing.T) {
	callWithTestWebsocket(
		func(ws WebSocket) {
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

func TestShouldWriteToWebsocket(t *testing.T) {
	callWithTestWebsocket(
		func(ws WebSocket) {
			handler := SlackMessageHandler{
				ID: testID,
				WS: ws,
			}
			handler.Send(message.M{
				Type:    message.TypeMessage,
				Text:    "message1",
				Channel: "channel1",
			})
			const message1 = `{"id":1,"type":"message","subtype":"","user":"","channel":"channel1","text":"message1","error":{"code":0,"msg":""}}`
			if messageSent != message1 {
				t.Errorf("Expected sent message to be %s, but it is: %s", message1, messageSent)
			}
			handler.Send(message.M{
				Type:    message.TypeMessage,
				Text:    "message2",
				Channel: "channel2",
			})
			const message2 = `{"id":2,"type":"message","subtype":"","user":"","channel":"channel2","text":"message2","error":{"code":0,"msg":""}}`
			if messageSent != message2 {
				t.Errorf("Expected sent message to be %s, but it is: %s", message2, messageSent)
			}
		},
		nil,
	)
}
