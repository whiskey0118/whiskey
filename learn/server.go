package main

import (
	"net/http"
)

func WsHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("hello"))

}

func main() {
	http.HandleFunc("/ws", WsHandler)
	http.ListenAndServe("localhost:8080", nil)
}
