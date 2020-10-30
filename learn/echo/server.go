package main

import (
	"flag"
	"github.com/gorilla/websocket"
	"log"
	"net/http"
)

var (
	addr     = flag.String("addr", "127.0.0.1:8080", "web")
	upGrader websocket.Upgrader
)

func echo(w http.ResponseWriter, r *http.Request) {
	c, err := upGrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println("upgrade:", err)
		return
	}
	defer c.Close()
	for {
		mt, message, err := c.ReadMessage()
		if err != nil {
			log.Println("read:", err)
			break
		}
		log.Printf("reac: %s", message)
		err = c.WriteMessage(mt, message)
		if err != nil {
			log.Println("write:", err)
			break
		}
	}
}

func main() {
	flag.Parse()
	log.SetFlags(0)
	http.HandleFunc("/echo", echo)
	http.ListenAndServe(*addr, nil)
}
