package web

import (
	"github.com/gorilla/websocket"
	"net/http"
)

var upgrader websocket.Upgrader

func InitConn(w http.ResponseWriter, r *http.Request) (conn *websocket.Conn, err error) {
	conn, err = upgrader.Upgrade(w, r, nil)
	return
}

func ReadMessage() {
	//conn := InitConn()
	//client,err := websocket2.InitClient()
}
