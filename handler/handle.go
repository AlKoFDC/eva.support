package handler

import (
	"fmt"
	"github.com/AlKoFDC/eva.support/logger"
	"github.com/AlKoFDC/eva.support/message"
)

type WebSocket interface {
	ReadJSON(interface{}) error
	WriteJSON(interface{}) error
	Close() error
}

// SlackMessage is a handler for slack messages from a web socket.
type SlackMessageHandler struct {
	PrintUnknown bool
	Name         string
	ID           string
	WS           WebSocket
	counter      uint64
}

// SlackMessage implements MessageHandler.
var _ MessageHandler = (*SlackMessageHandler)(nil)

// SlackMessage handles incoming slack messages to ws.
// It is designed to be started as a go routine.
func (h *SlackMessageHandler) Handle() {
	for {
		// Read message.
		msg := message.M{}
		err := h.WS.ReadJSON(&msg)
		if err != nil {
			logger.Error.Println(err)
			return
		}
		switch msg.Type {
		case message.TypeHello:
			logger.Standard.Println("Connection successful.")
		case message.TypeError:
			logger.Error.Println("Error during connection:", msg.Error.Code, "-", msg.Error.Message)
		case message.TypeMessage:
			h.handleTypeMessage(msg)
		case message.TypeReconnectURL, message.TypePresenceChange, message.TypeTyping:
		// Do nothing. FIXME yet?
		default:
			if h.PrintUnknown {
				logger.Error.Println(fmt.Sprintf("Unknown message: %#v", msg))
			}
		}
	}
}

// Send sends a message to the underlying websocket.
func (h *SlackMessageHandler) Send(msg message.M) error {
	h.counter++
	msg.Id = h.counter
	return h.WS.WriteJSON(msg)
}

// Close closes the underlying websocket.
func (h *SlackMessageHandler) Close() {
	h.WS.Close()
}
