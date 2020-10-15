package main

import "net/http"

func main() {
	http.HandleFunc("/", post)
}

func post(r http.ResponseWriter, re *http.Request) {
	r.Write([]byte("haha"))

}
