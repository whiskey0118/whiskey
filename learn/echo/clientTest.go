package main

import (
	"fmt"
	"github.com/gorilla/websocket"
	"log"
	"time"
)

func main() {
	connect, _, err := websocket.DefaultDialer.Dial("ws://localhost:8080/echo", nil)
	if err != nil {
		fmt.Println("conn err:", err)
	}

	ticker := time.NewTicker(time.Second)
	defer ticker.Stop()

	t := <-ticker.C
	for {
		err = connect.WriteMessage(websocket.TextMessage, []byte(t.String()))
		if err != nil {
			log.Println("write:", err)
			return
		}
	}

	_, mes, err := connect.ReadMessage()
	if err != nil {
		fmt.Println(err)
	}

	fmt.Printf("%v", mes)

}
