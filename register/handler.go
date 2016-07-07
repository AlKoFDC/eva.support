package register

import "github.com/AlKoFDC/eva.support/message"

type Handler interface {
	Handle(message.M)
}
