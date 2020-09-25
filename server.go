package main

import (
	"github.com/gorilla/websocket"
	"log"
	"net/http"
)

var upgrader = websocket.Upgrader{}

func root(w http.ResponseWriter, r *http.Request) {
	c, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Fatal("upgrader err:", err)
	}
	defer c.Close()

	for {
		mt, p, err := c.ReadMessage()
		if err != nil {
			log.Println("read message err: ", err)
			break
		}
		log.Printf("recive: %s", p)
		err = c.WriteMessage(mt, p)
		if err != nil {
			log.Println("write msg err: ", err)
			break
		}
	}
}

func main() {
	http.HandleFunc("/test", root)
	http.ListenAndServe("127.0.0.1:8090", nil)
}
