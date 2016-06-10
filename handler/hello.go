package handler

import (
	"github.com/AlKoFDC/eva.support/message"
)

// hello reacts to a message, that greets the bot.
func (h *SlackMessageHandler) hello(msg message.M) {
	// Respond on the same channel.
	msg.Text = "Hello, <@" + msg.User + ">. How can I help you? :blush:"
	h.Send(msg)
}
