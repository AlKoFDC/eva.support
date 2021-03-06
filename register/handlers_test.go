package register

import "github.com/AlKoFDC/eva.support/message"

type noop struct{}

var _ Handler = (*noop)(nil)

func (h noop) Handle(Receiver, message.M) {
	return
}

// callCounter counts the amount of times it was called.
type callCounter struct {
	callCount int
}

var _ Handler = (*callCounter)(nil)

func (h *callCounter) Handle(Receiver, message.M) {
	h.callCount++
	return
}
