package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"github.com/gorilla/websocket"
	"log"
	url2 "net/url"
	"os"
	"time"
	websocket2 "whiskey/websocket"
)

func main() {
	addr := flag.String("addr", "127.0.0.1:8080", "address")
	flag.Parse()

	var mess websocket2.Message
	interrup := make(chan os.Signal)
	url := url2.URL{
		Scheme:     "ws",
		Opaque:     "",
		User:       nil,
		Host:       *addr,
		Path:       "/read",
		RawPath:    "",
		ForceQuery: false,
		RawQuery:   "",
		Fragment:   "",
	}
	conn, _, err := websocket.DefaultDialer.Dial(url.String(), nil)
	if err != nil {
		log.Fatal("dial err:", err)
	}
	defer conn.Close()

	done := make(chan struct{})

	//go func() {
	//	defer close(done)
	//	for {
	//		_, _, err := conn.ReadMessage()
	//		if err != nil {
	//			log.Println("read:", err)
	//			return
	//		}
	//		//log.Printf("recv: %s", message)
	//	}
	//}()

	tic := time.NewTicker(2 * time.Second)
	defer tic.Stop()

	for {
		select {
		case <-tic.C:
			mess.Body = "test"
			mess.Type = "client1"
			res, _ := json.Marshal(&mess)
			fmt.Println(string(res))
			conn.WriteMessage(websocket.TextMessage, res)
		//case <-done:
		//	return
		case <-interrup:
			log.Println("interrupt")
			// Cleanly close the connection by sending a close message and then
			// waiting (with timeout) for the server to close the connection.
			err := conn.WriteMessage(websocket.CloseMessage, websocket.FormatCloseMessage(websocket.CloseNormalClosure, ""))
			if err != nil {
				log.Println("write close:", err)
				return
			}
			select {
			case <-done:
			case <-time.After(time.Second):
			}
			return
		}
	}
}
