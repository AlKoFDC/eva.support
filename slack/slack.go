package slack

import (
	"encoding/json"
	"fmt"
	"github.com/AlKoFDC/eva.support/handler"
	"github.com/gorilla/websocket"
	"io/ioutil"
	"net/http"
)

// These two structures represent the response of the Slack API rtm.start.
// Only some fields are included. The rest are ignored by json.Unmarshal.

type responseRtmStart struct {
	Ok    bool         `json:"ok"`
	Error string       `json:"error"`
	Url   string       `json:"url"`
	Self  responseSelf `json:"self"`
}

type responseSelf struct {
	Id string `json:"id"`
}

// slackStart does a rtm.start, and returns a websocket URL and user ID. The
// websocket URL can be used to initiate an RTM session.
func slackStart(token string) (wsurl, id string, err error) {
	url := fmt.Sprintf("https://slack.com/api/rtm.start?token=%s", token)
	resp, err := http.Get(url)
	if err != nil {
		return
	}
	if resp.StatusCode != 200 {
		err = fmt.Errorf("API request failed with code %d", resp.StatusCode)
		return
	}
	body, err := ioutil.ReadAll(resp.Body)
	resp.Body.Close()
	if err != nil {
		return
	}
	var respObj responseRtmStart
	err = json.Unmarshal(body, &respObj)
	if err != nil {
		return
	}

	if !respObj.Ok {
		err = fmt.Errorf("Slack error: %s", respObj.Error)
		return
	}

	wsurl = respObj.Url
	id = respObj.Self.Id
	return
}

// Connect starts a websocket-based Real Time API session and returns a Slack Message Handler.
func Connect(token string) (handler.SlackMessageHandler, error) {
	emptyResponse := handler.SlackMessageHandler{}
	wsConnection, id, err := getWSandID(token)
	if err != nil {
		return emptyResponse, err
	}
	return handler.SlackMessageHandler{WS: wsConnection, ID: id}, nil
}

// ConnectAsynch starts a websocker-based Real Time API session and returns
// an asynchronous Slack Message Handler.
func ConnectAsynch(token string) (handler.AsynchSlackMessageHandler, error) {
	emptyResponse := handler.AsynchSlackMessageHandler{}
	wsConnection, id, err := getWSandID(token)
	if err != nil {
		return emptyResponse, err
	}
	return handler.AsynchSlackMessageHandler{WS: wsConnection, ID: id}, nil
}

// ConnectRegister starts a websocket-based Real Time API session and returns
// a Slack Message Handler, that uses the register package.
func ConnectRegister(token string) (handler.RegisterSlackHandler, error) {
	emptyResponse := handler.RegisterSlackHandler{}
	wsConnection, id, err := getWSandID(token)
	if err != nil {
		return emptyResponse, err
	}
	return handler.RegisterSlackHandler{WS: wsConnection, ID: id}, nil
}

// getWSandID connects to slack with the token and returns the websocket connection and an ID,
// that identifies the user.
func getWSandID(token string) (*websocket.Conn, string, error) {
	var empty *websocket.Conn = nil
	response := ""
	wsurl, id, err := slackStart(token)
	if err != nil {
		return empty, response, err
	}

	wsConnection, _, err := websocket.DefaultDialer.Dial(wsurl, nil)
	if err != nil {
		return empty, response, err
	}
	return wsConnection, id, nil
}
