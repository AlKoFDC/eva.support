package handler

import (
	"github.com/AlKoFDC/eva.support/logger"
	"github.com/AlKoFDC/eva.support/message"
	"github.com/AlKoFDC/eva.support/register"
)

type RegisterSlackHandler struct {
	ID string
	WS WebSocket

	Name         string
	PrintUnknown bool

	bot *register.Bot

	counter uint64
}

// Start starts the handling of messages and also cleans it up
// on receiving a signal on the finish channel or on channel close.
func (sh *RegisterSlackHandler) Start(finish chan struct{}) {
	if sh.bot == nil {
		sh.bot = &register.Bot{}
	}
	msgChan := sh.bot.Start(sh)
	go sh.sendMessagesToChannel(msgChan)
	defer sh.close()
	<-finish
}

func (sh *RegisterSlackHandler) Send(msg message.M) error {
	sh.counter++
	msg.Id = sh.counter
	return sh.WS.WriteJSON(msg)
}

// sendMessagesToChannel sends the messages, that the websocket receives to
// msgChan.
// When an error is received on the websocket, it closes msgChan and returns.
func (sh *RegisterSlackHandler) sendMessagesToChannel(msgChan chan message.M) {
	defer close(msgChan)
	for {
		msg := message.M{}
		err := sh.WS.ReadJSON(&msg)
		if err != nil {
			logger.Error.Println(err, msg)
			return
		}
		msgChan <- msg
	}
}

// close frees underlying structures.
func (sh *RegisterSlackHandler) close() {
	sh.WS.Close()
}

// AddCase adds a condition and a handler. If c is true, h is called.
// The conditions are checked in the order they are added.
func (sh *RegisterSlackHandler) AddCase(c register.Conditioner, h register.Handler) error {
	if sh.bot == nil {
		sh.bot = &register.Bot{}
	}
	return sh.bot.Case(c, h)
}

// RemoveCase removes the condition.
func (sh *RegisterSlackHandler) RemoveCase(c register.Conditioner) {
	if sh.bot == nil {
		sh.bot = &register.Bot{}
	}
	sh.bot.Remove(c)
}

// DefaultCase adds h as the default handler, that is called when no condition is true.
// If h is nil, the default is removed.
func (sh *RegisterSlackHandler) DefaultCase(h register.Handler) {
	if sh.bot == nil {
		sh.bot = &register.Bot{}
	}
	sh.bot.Default(h)
}

var _ register.Caller = (*RegisterSlackHandler)(nil)

// Ident returns the identifications of the handler to satisfy the register
// interface Caller.
func (sh *RegisterSlackHandler) Ident() (string, string) {
	return sh.ID, sh.Name
}
