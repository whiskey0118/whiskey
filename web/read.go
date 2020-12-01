package web

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/websocket"
	"log"
	"net/http"
	websocket2 "whiskey/websocket"
)

var upgrader websocket.Upgrader

func init() {
	//设置日志头部信息
	log.SetPrefix("<Read message>")
	//设置日志的格式
	log.SetFlags(log.LstdFlags | log.Ldate | log.Ltime | log.Lshortfile)
}

func ReadMessage(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	client, err := websocket2.InitClient(conn, nil, nil, "")
	if err != nil {
		log.Fatal("create client fail: ", err)
	}

	go func() {
		var message websocket2.Message
		for {
			_, b, err := client.Conn.ReadMessage()
			if err != nil {
				fmt.Println("read body err:", err)
				break
			}
			err = json.Unmarshal(b, &message)
			if err != nil {
				log.Printf("client addr:%s message:json unmarshal err %s", r.RemoteAddr, err)
				return
			}
			client.InChan <- &message
		}
	}()

	go func() {
		for {
			client.ReadMessage()
		}
	}()
}
