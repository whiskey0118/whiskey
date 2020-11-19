package web

import (
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
	if err != nil {
		log.Fatal()
	}
	client, err := websocket2.InitClient(conn, nil, nil, "")
	if err != nil {
		log.Fatal("create client fail: ", err)
	}

	go func() {
		mes, _ := client.ReadMessage()
		log.Printf("%s", mes)
		defer client.Close()
	}()

}
