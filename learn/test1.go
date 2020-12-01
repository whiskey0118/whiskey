package main

import (
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

type Clientest struct {
	Mes string
	Ch  chan string
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
	cli := Clientest{
		Mes: "",
		Ch:  make(chan string),
	}

	var a string
	a = string(body)
	//fmt.Printf("a:",a)
	cli.Ch <- a
	go test1(&cli)
}

func test1(cli *Clientest) {
	select {
	case mes := <-cli.Ch:
		fmt.Printf("mes:", mes)
	}
}

func main() {
	http.HandleFunc("/test", test)
	http.ListenAndServe(":8888", nil)
}
