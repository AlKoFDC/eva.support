package handler

import (
	"fmt"
	"github.com/AlKoFDC/eva.support/logger"
	"github.com/AlKoFDC/eva.support/message"
)

// hello reacts to a message, that greets the bot.
func (h *SlackMessageHandler) hello(msg message.M) {
	logger.Standard.Println(fmt.Sprintf("I received a greeting from %s! *blush*", msg.User))
}
