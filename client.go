package main

import (
	"fmt"
	"github.com/gorilla/websocket"
	"log"
)

func main() {
	cli := websocket.Dialer{
		NetDial:           nil,
		NetDialContext:    nil,
		Proxy:             nil,
		TLSClientConfig:   nil,
		HandshakeTimeout:  0,
		ReadBufferSize:    0,
		WriteBufferSize:   0,
		WriteBufferPool:   nil,
		Subprotocols:      nil,
		EnableCompression: false,
		Jar:               nil,
	}

	_, rep, err := cli.Dial("ws://127.0.0.1:8090/test", nil)
	if err != nil {
		log.Println("err:", err)
	}
	fmt.Printf("%s", rep.Body)
}
