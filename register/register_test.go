package register

import (
	"github.com/AlKoFDC/eva.support/message"
	"testing"
)

func TestShouldNotAddSecondHandlerForSameCondition(t *testing.T) {
	c1 := alwaysTrue{}
	c2 := alwaysTrue{}
	h := noop{}
	b := Bot{}
	if err := b.Case(c1, h); err != nil {
		t.Errorf("Couldn't add first condition: %s", err)
	}
	if err := b.Case(c2, h); err == nil {
		t.Error("Expected to get error when adding the same condition.")
	}
}

func TestShouldAllowToAddConditionAgainAfterItGotDeleted(t *testing.T) {
	c1 := alwaysTrue{}
	h := noop{}
	b := Bot{}
	if err := b.Case(c1, h); err != nil {
		t.Errorf("Couldn't add first condition: %s", err)
	}
	b.Remove(c1)
	if cLen := len(b.conditionOrder); cLen != 0 {
		t.Errorf("Expected no conditions to be registered in the slice, but got %d.", cLen)
	}
	if err := b.Case(c1, h); err != nil {
		t.Errorf("Couldn't add condition the second time after removing: %s", err)
	}
}

func TestShouldNotBeAbleToAddCaseWithEitherNil(t *testing.T) {
	b := Bot{}
	if err := b.Case(nil, noop{}); err == nil {
		t.Errorf("Expected to get an error when adding a nil condition, but didn't.")
	}
	if err := b.Case(alwaysTrue{}, nil); err == nil {
		t.Errorf("Expected to get an error when adding a nil handler, but didn't.")
	}
}

func TestShouldUseFirstConditionThatIsTrue(t *testing.T) {
	b := Bot{}
	c1, c2, c3, c4 := alwaysFalse{}, alwaysFalse2{}, alwaysTrue{}, alwaysFalse{}
	h1, h2, h3, h4 := callCounter{}, callCounter{}, callCounter{}, callCounter{}
	b.Case(c1, &h1)
	b.Case(c2, &h2)
	b.Case(c3, &h3)
	b.Case(c4, &h4)
	msgChan := b.Start()
	msgChan <- message.M{}
	close(msgChan)
	if h1.callCount != 0 || h2.callCount != 0 || h3.callCount != 1 || h4.callCount != 0 {
		t.Errorf("Expected first true function to be called once, but got:\nfalse1: %d calls, false2: %d calls, true1: %d calls, true2: %d calls",
			h1.callCount, h2.callCount, h3.callCount, h4.callCount)
	}
}

func TestShouldUseDefaultWhenNoConditionFits(t *testing.T) {
	b := Bot{}
	c := alwaysFalse{}
	h := callCounter{}
	d := callCounter{}
	b.Default(&d)
	b.Case(c, &h)
	msgChan := b.Start()
	msgChan <- message.M{}
	close(msgChan)
	if h.callCount != 0 || d.callCount != 1 {
		t.Errorf("Expected the default handler to be called once, but got:\nfalse: %d calls, default: %d calls",
			h.callCount, d.callCount)
	}
}
