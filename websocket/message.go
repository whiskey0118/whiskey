package websocket

import "time"

type Message struct {
	Time time.Time
	Type int    `json:"type"`
	Body string `json:"body"`
}
