package register

import (
	"github.com/AlKoFDC/eva.support/message"
)

type Receiver interface {
	Send(message.M) error
}

type Handler interface {
	Handle(Receiver, message.M)
}
