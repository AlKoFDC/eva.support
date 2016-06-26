package handler

import (
	"github.com/AlKoFDC/eva.support/message"
	"testing"
	"time"
)

func TestMessageHandling(t *testing.T) {
	send, receive := make(chan message.M), make(chan message.M)
	// This functions sends, the handler receives and vice versa.
	// That's why the channels are switched around.
	h := AsynchSlackMessageHandler{WS: nil, ID: "", send: receive, receive: send}
	go h.handle()
	defer h.close()

	send <- message.M{
		Id:   42,
		Type: message.TypeHelloWorld,
	}

	const (
		maxResponseTime = 1 * time.Second
		expectedAnswer  = "Hello, world!"
	)
	select {
	case answer, ok := <-receive:
		if !ok {
			t.Fatal("Handler's send channel was closed before receiving an answer.")
		}
		if answer.Text != expectedAnswer {
			t.Fatalf("Expected text of answer to be '%s', but got %#v.", expectedAnswer, answer)
		}
	case <-time.After(maxResponseTime):
		t.Fatalf("Didn't receive a response in %s.", maxResponseTime)
	}
}
