package websocket

import "time"

type AccessMessage struct {
	Time   time.Time
	From   interface{}
	To     interface{}
	Status string
	Reason interface{}
}
