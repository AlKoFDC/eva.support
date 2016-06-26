package handler

import (
	"github.com/AlKoFDC/eva.support/message"
	"testing"
	"time"
)

const (
	testID = "test"
)

func TestMessageHandling(t *testing.T) {
	callWithMockedAsynchHandler(func(send, receive chan message.M) {
		send <- message.M{
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
	})
}

func TestShouldHandleHelloMessage(t *testing.T) {
	callWithMockedAsynchHandler(func(send, receive chan message.M) {
		const testUser = "testUser"
		send <- message.M{
			User: testUser,
			Type: message.TypeMessage,
			Text: "Hello @" + testID,
		}

		const (
			maxResponseTime = 1 * time.Second
			expectedAnswer  = "Hello, <@testUser>. How can I help you? :blush:"
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
	})
}

func callWithMockedAsynchHandler(do func(send, receive chan message.M)) {
	send, receive := make(chan message.M), make(chan message.M)
	// This functions sends, the handler receives and vice versa.
	// That's why the channels are switched around.
	h := AsynchSlackMessageHandler{WS: nil, ID: testID, send: receive, receive: send}
	go h.handle()
	do(send, receive)
	defer h.close()
}
