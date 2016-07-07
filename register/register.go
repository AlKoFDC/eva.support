package register

import (
	"errors"
	"github.com/AlKoFDC/eva.support/message"
)

type Bot struct {
	conditionOrder    []Conditioner
	conditionHandlers map[Conditioner]Handler
	defaultHandler    Handler

	finish chan struct{}
}

// Case adds a condition and a handler. If c is true, h is called.
// The conditions are checked in the order they are added.
func (b *Bot) Case(c Conditioner, h Handler) error {
	if b.conditionHandlers == nil {
		b.conditionHandlers = make(map[Conditioner]Handler)
	}
	if b.conditionOrder == nil {
		b.conditionOrder = make([]Conditioner, 0)
	}
	if c == nil || h == nil {
		return errors.New("neither the condition nor the handler can be nil")
	}
	if _, ok := b.conditionHandlers[c]; ok {
		return errors.New("there is already a handler defined for this condition, remove it first")
	}
	b.conditionOrder = append(b.conditionOrder, c)
	b.conditionHandlers[c] = h
	return nil
}

// Remove removes the condition from the bot.
func (b *Bot) Remove(c Conditioner) {
	for index, registered := range b.conditionOrder {
		if registered == c {
			b.conditionOrder = append(b.conditionOrder[:index], b.conditionOrder[index+1:]...)
		}
	}
	delete(b.conditionHandlers, c)
}

// Default adds h as the default handler, that is called when no condition is true.
// If h is nil, the default is removed.
func (b *Bot) Default(h Handler) {
	// Add h as default.
	b.defaultHandler = h
}

// Start starts the handling of the messages and returns the channel, on which
// it expects to receive messages.
// If the channel is closed, the handling stops.
func (b *Bot) Start() chan message.M {
	mChan := make(chan message.M)
	go b.handleMessages(mChan)
	return mChan
}

// handleMessages handles the messages by checking if the condition is true
func (b *Bot) handleMessages(messages chan message.M) {
mainLoop:
	for msg := range messages {
		// Check for conditions.
		for _, condition := range b.conditionOrder {
			if condition.IsTrue(msg) {
				b.conditionHandlers[condition].Handle(msg)
				continue mainLoop
			}
		}
		// If no condition matches, use the default handler.
		if b.defaultHandler != nil {
			b.defaultHandler.Handle(msg)
		}
	}
}
