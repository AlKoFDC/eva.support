package handler

import (
	"github.com/AlKoFDC/eva.support/message"
)

// hello reacts to a message, that greets the bot.
func (h *handler) hello(msg message.M) message.M {
	// Respond on the same channel.
	msg.Text = "Hello, <@" + msg.User + ">. How can I help you? :blush:"
	return msg
}
