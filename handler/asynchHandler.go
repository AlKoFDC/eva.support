package handler

import (
	"fmt"
	"github.com/AlKoFDC/eva.support/logger"
	"github.com/AlKoFDC/eva.support/message"
	"runtime/debug"
)

type AsynchSlackMessageHandler struct {
	ID string
	WS WebSocket

	Name         string
	PrintUnknown bool

	receive chan message.M
	send    chan message.M

	counter uint64
}

// AsynchSlackMessageHandler implements AsynchMessageHandler.
var _ AsynchMessageHandler = (*AsynchSlackMessageHandler)(nil)

// Start starts the asynchronous handling of messages and also cleans it up
// on receiving a signal on the finish channel or on channel close.
func (h *AsynchSlackMessageHandler) Start(finish chan struct{}) {
	h.send, h.receive = make(chan message.M, 10), make(chan message.M, 10)
	go h.receiveMessage()
	go h.sendMessage()
	go h.handle()

	defer h.close()
	<-finish
}

// receiveMessage synchronizes the receiving of messages.
func (h *AsynchSlackMessageHandler) receiveMessage() {
	defer func() {
		// If the read is called too often on a closed websocket, it
		// will panic. That's when we return.
		if err := recover(); err != nil {
			logger.Error.Println(fmt.Sprintf("Exited message receiver with a panic.\n%s\n%s", err.(error), debug.Stack()))
		}
	}()
	// Close receive to return the handle function.
	defer close(h.receive)
	for {
		// Read message.
		msg := message.M{}
		err := h.WS.ReadJSON(&msg)
		if err != nil {
			logger.Error.Println(err)
			continue
		}
		h.receive <- msg
	}
}

// sendMessage synchronizes the sending of messages.
func (h *AsynchSlackMessageHandler) sendMessage() {
	for {
		msg, ok := <-h.send
		if !ok {
			// Return on closed send channel.
			return
		}
		h.counter++
		msg.Id = h.counter
		err := h.WS.WriteJSON(msg)
		if err != nil {
			logger.Error.Println(fmt.Sprintf("Error while sending message: %s\nMessage: %#v", err, msg))
		}
	}
}

// This function provides asynchronous message handling.
func (h *AsynchSlackMessageHandler) handle() {
	// Close send on returning to return the send function.
	defer close(h.send)
	mh := handler{
		name: h.Name, id: h.ID, printUnknown: h.PrintUnknown,
	}
	for {
		msg, ok := <-h.receive
		if !ok {
			// Return on closed receive channel.
			return
		}
		switch msg.Type {
		case message.TypeHelloWorld:
			h.send <- handleHelloWorld(msg)
		case message.TypeHello:
			logger.Standard.Println("Connection successful.")
		case message.TypeError:
			logger.Error.Println("Error during connection:", msg.Error.Code, "-", msg.Error.Message)
		case message.TypeMessage:
			answer, send := mh.handleTypeMessage(msg)
			if send {
				h.send <- answer
			}
		case message.TypeReconnectURL, message.TypePresenceChange, message.TypeTyping:
		// Do nothing. FIXME yet?
		default:
			if h.PrintUnknown {
				logger.Error.Println(fmt.Sprintf("Unknown message: %#v", msg))
			}
		}
	}
}

// Close frees underlying structures.
func (h *AsynchSlackMessageHandler) close() {
	if h.WS != nil {
		h.WS.Close()
		h.WS = nil
	}
}
