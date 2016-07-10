package register

import "github.com/AlKoFDC/eva.support/message"

type alwaysTrue struct{}

var _ Conditioner = (*alwaysTrue)(nil)

func (c alwaysTrue) IsTrue(Caller, message.M) bool {
	return true
}

type alwaysTrue2 struct{}

var _ Conditioner = (*alwaysTrue2)(nil)

func (c alwaysTrue2) IsTrue(Caller, message.M) bool {
	return true
}

type alwaysFalse struct{}

var _ Conditioner = (*alwaysFalse)(nil)

func (c alwaysFalse) IsTrue(Caller, message.M) bool {
	return false
}

type alwaysFalse2 struct{}

var _ Conditioner = (*alwaysFalse2)(nil)

func (c alwaysFalse2) IsTrue(Caller, message.M) bool {
	return false
}
