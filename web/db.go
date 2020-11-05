package web

import (
	"log"
	"net/http"
	"time"
	"whiskey/websocket"
)

func PostDevice(r http.ResponseWriter, q *http.Request) {
	d := websocket.Device{
		Uid:          "test",
		Version:      "s301",
		Area:         "china",
		Status:       "enable",
		ExpiringDate: time.Now(),
	}
	_, err := d.InsertUser()
	if err != nil {
		log.Fatal("insert err :", err)
		return
	}
	r.Write([]byte("insert successful"))
}
