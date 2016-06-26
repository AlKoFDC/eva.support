package handler

import "github.com/AlKoFDC/eva.support/message"

func handleHelloWorld(msg message.M) message.M {
	msg.Text = "Hello, world!"
	return msg
}
