package handler

import (
	"fmt"
	"github.com/AlKoFDC/eva.support/logger"
	"github.com/AlKoFDC/eva.support/message"
)

type handler struct {
	name         string
	id           string
	printUnknown bool
}

// handleTypeMessage handles the message and returns an answer, if needed.
// The second return parameter is true, if tha answer needs to be sent.
func (h *handler) handleTypeMessage(msg message.M) (message.M, bool) {
	msgForMe := msg.IsForMe(h.name, h.id)
	if !msgForMe {
		return msg, false
	}
	switch {
	case msg.IsHelloMessage():
		return h.hello(msg), true
	case msg.Subtype == message.SubtypeChanged:
		// Do nothing. FIXME yet?
		return msg, false
	default:
		if h.printUnknown {
			logger.Error.Println(fmt.Sprintf("Unknown subtype: %#v", msg))
		}
		return msg, false
	}
}

// hello reacts to a message, that greets the bot.
func (h *AsynchSlackMessageHandler) hello(msg message.M) {
	// Respond on the same channel.
	msg.Text = "Hello, <@" + msg.User + ">. How can I help you? :blush:"
	h.send <- msg
}
