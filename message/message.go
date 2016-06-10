package message

import "strings"

// These are the messages read off and written into the websocket. Since this
// struct serves as both read and write, we include the "Id" field which is
// required only for writing.
type M struct {
	Id      uint64     `json:"id"`
	Type    string     `json:"type"`
	Subtype string     `json:"subtype"`
	User    string     `json:"user"`
	Channel string     `json:"channel"`
	Text    string     `json:"text"`
	Error   SlackError `json:"error"`
}

// SlackError is a structure to hold the error message of a response.
type SlackError struct {
	Code    uint64 `json:"code"`
	Message string `json:"msg"`
}

func (m M) IsHelloMessage() bool {
	return strings.Contains(strings.ToLower(m.Text), "hello")
}
