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
	//忘记处理这个错误，导致在浏览器访问进来的连接不是websocket类型，就会导致整个进程挂掉
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Printf("Remote addr:%v, err: %v", r.RemoteAddr, err)
		return
	}
	client, err := websocket2.InitClient(conn, nil, nil, "")
	if err != nil {
		log.Fatal("create client fail: ", err)
	}

	go client.WriteMessage()

	go client.ReadMessage()
}
