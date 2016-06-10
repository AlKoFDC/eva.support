package slack

import (
	"encoding/json"
	"fmt"
	"github.com/AlKoFDC/eva.support/handler"
	"github.com/gorilla/websocket"
	"io/ioutil"
	"log"
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

// Starts a websocket-based Real Time API session and return a slack message handler.
func Connect(token string) handler.SlackMessageHandler {
	wsurl, id, err := slackStart(token)
	if err != nil {
		log.Fatal(err)
	}

	wsConnection, _, err := websocket.DefaultDialer.Dial(wsurl, nil)
	if err != nil {
		log.Fatal(err)
	}

	return handler.SlackMessageHandler{WS: wsConnection, ID: id}
}
