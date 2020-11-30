package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type AutotaskRequest struct {
	RequestID string     `json:"requestid"`
	Clone     CloneModel `json:"clone"`
	Push      PushModel  `json:"push"`
}

type CloneModel struct {
	//TODO
	//"Method": string `json:"ceph"`
	RequestID   string `json:"requestid"`
	CallbackURL string `json:"callbackurl"`
}

type PushModel struct {
	RequestID   string `json:"requestiD"`
	CallbackURL string `json:"callbackuRL"`
	IP          string `json:"remoteip"`
	Port        int    `json:"remoteport"`
	User        string `json:"user"`
}

type Messages struct {
	//Time time.Time `json:"time"`
	Type string `json:"type"`
	Body string `json:"body"`
}

func test(w http.ResponseWriter, r *http.Request) {
	// r.ParseForm()
	defer fmt.Fprintf(w, "ok\n")

	fmt.Println("method:", r.Method)
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Printf("read body err, %v\n", err)
		return
	}
	println("json:", string(body))

	var a Messages
	if err = json.Unmarshal(body, &a); err != nil {
		fmt.Printf("Unmarshal err, %v\n", err)
		return
	}
	fmt.Printf("%+v", a)

}

func main() {
	http.HandleFunc("/test", test)
	http.ListenAndServe(":8888", nil)
}
