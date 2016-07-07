package register

import "github.com/AlKoFDC/eva.support/message"

type Conditioner interface {
	IsTrue(message.M) bool
}
