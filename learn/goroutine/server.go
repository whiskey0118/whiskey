package main

import "net/http"

func main() {
	http.HandleFunc("/", post)
	http.ListenAndServe("127.0.0.1:8081", nil)
}

func post(r http.ResponseWriter, re *http.Request) {
	if re.Header.Get("one") == "me" {
		r.Write([]byte("yes"))
	} else {
		r.Write([]byte("no"))
	}

}
