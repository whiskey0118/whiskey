package websocket

import "github.com/gorilla/websocket"

type Client struct {
	Device *Device
	Conn   *websocket.Conn
	Pool   *Pool
}
