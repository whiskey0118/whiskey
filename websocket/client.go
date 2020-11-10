package websocket

import (
	"github.com/gorilla/websocket"
)

type Client struct {
	Device  *Device
	Conn    *websocket.Conn
	Pool    *Pool
	InChan  chan *Message
	OutChan chan *Message
	Token   string
}

func InitClient(conn *websocket.Conn, device *Device, pool *Pool, token string) (client *Client, err error) {
	client = &Client{
		Device:  device,
		Conn:    conn,
		Pool:    pool,
		InChan:  make(chan *Message, 100),
		OutChan: make(chan *Message, 100),
		Token:   token,
	}
	return
}

func (client *Client) ReadMessage() (message *Message, err error) {
	message = <-client.InChan
	return
}

func (client *Client) WriteMessage(message *Message) (err error) {
	client.OutChan <- message
	return
}

func (client *Client) Close() {
	client.Conn.Close()
}
