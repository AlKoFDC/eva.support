package handler

import (
	"fmt"
	"github.com/AlKoFDC/eva.support/logger"
	"github.com/AlKoFDC/eva.support/message"
	"strings"
)

// handleTypeMessage handles a message of the type message.
func (h *SlackMessageHandler) handleTypeMessage(msg message.M) {
	msgForMe := h.isForMe(msg)
	switch {
	case msgForMe && msg.IsHelloMessage():
		h.hello(msg)
	case msg.Subtype == message.SubtypeChanged:
	// Do nothing. FIXME yet?
	default:
		if h.PrintUnknown {
			logger.Error.Println(fmt.Sprintf("Message: %#v", msg))
		}
	}
}

// isForMe returns true if the message was sent @ the ID set in the SlackMessageHandler
// or for any variation of the set bot name.
func (h *SlackMessageHandler) isForMe(msg message.M) bool {
	return (h.Name != "" && strings.Contains(strings.ToLower(msg.Text), "@"+h.Name)) ||
		(h.ID != "" && strings.Contains(msg.Text, "@"+h.ID))
}
