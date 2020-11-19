package websocket

import "time"

type Message struct {
	Time time.Time
	Type string `json:"type"`
	Body string `json:"body"`
}
