package websocket

import "time"

type Device struct {
	Uid          string    `json:"uid"`
	Version      string    `json:"version"`
	Area         string    `json:"area"`
	Status       string    `json:"status"`
	ExpiringDate time.Time `json:"expiring_date"`
}
