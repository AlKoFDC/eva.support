package message

import (
	"strings"
)

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

// isForMe returns true if the message was sent @ the ID set in the SlackMessageHandler
// or for any variation of the set bot name.
func (msg M) IsForMe(name, id string) bool {
	return (name != "" && strings.Contains(strings.ToLower(msg.Text), "@"+name)) ||
		(id != "" && strings.Contains(msg.Text, "@"+id))
}
