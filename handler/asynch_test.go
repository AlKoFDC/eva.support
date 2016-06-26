package handler_test

import (
	. "github.com/AlKoFDC/eva.support/handler"
	"runtime"
	"testing"
	"time"
)

func TestShouldCallWebsocketCloseAsynch(t *testing.T) {
	callWithTestWebsocket(
		func(ws WebSocket) {
			handler := AsynchSlackMessageHandler{
				ID: "",
				WS: ws,
			}
			finish := make(chan struct{})
			go handler.Start(finish)
			close(finish)
			// Give the go routine some time to react.
			success := make(chan bool)
			go func() {
				for {
					runtime.Gosched()
					if closeCalled {
						success <- true
					}
				}
			}()
			const maxWaitTime = 1 * time.Second
			select {
			case <-success:
			case <-time.After(maxWaitTime):
				t.Errorf("Expected close to be called after %s, but it was not.", maxWaitTime)
			}
		},
	)
}

func TestShouldReceiveAndSendMessagesThroughWebsocket(t *testing.T) {
	t.Skip("Send a message to the bot.")
	callWithTestWebsocket(
		func(ws WebSocket) {
			handler := AsynchSlackMessageHandler{
				ID: testID,
				WS: ws,
			}
			finish := make(chan struct{})
			go handler.Start(finish)
			// Give the go routine some time to react.
			success := make(chan bool)
			go func() {
				for {
					runtime.Gosched()
					if messageSent != "" {
						success <- true
					}
				}
			}()
			const (
				maxWaitTime     = 1 * time.Second
				expectedMessage = ""
			)
			select {
			case <-success:
				if messageSent != expectedMessage {
					t.Errorf("Expected message '%s', but got '%s'.", expectedMessage, messageSent)
				}
			case <-time.After(maxWaitTime):
				t.Errorf("Expected a message after %s, but got none.", maxWaitTime)
			}
			close(finish)
		},
		helloWorldTestMessage,
	)
}
