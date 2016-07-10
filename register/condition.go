package register

import (
	"github.com/AlKoFDC/eva.support/message"
)

type Caller interface {
	Ident() (id string, name string)
}

type Conditioner interface {
	IsTrue(Caller, message.M) bool
}
