package register

import (
	"github.com/AlKoFDC/eva.support/message"
	"testing"
)

type statefulCondition struct {
	count int
}

var _ Conditioner = (*statefulCondition)(nil)

func (c *statefulCondition) IsTrue(Caller, message.M) bool {
	c.count++
	return c.count == 3
}

func TestStatefulCondition(t *testing.T) {
	b := Bot{}
	c := statefulCondition{}
	h := callCounter{}
	b.Case(&c, &h)
	msgChan := b.Start(&mockCallerReceiver{})
	defer close(msgChan)
	for count := 0; count < 5; count++ {
		msgChan <- message.M{}
	}
	if called := h.callCount; called != 1 {
		t.Errorf("Expected handler to be called once, but it was called %d times.", called)
	}
}
