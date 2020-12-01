//8.7
package main

import (
	"log"
	"time"
)

func Chann(ch chan int, stopCh chan bool) {
	var i int
	i = 10
	for j := 0; j < 10; j++ {
		ch <- i
		time.Sleep(time.Second)
	}
	stopCh <- true
}

type Messages1 struct {
	//Time time.Time `json:"time"`
	Type string `json:"type"`
	Body string `json:"body"`
}

func main() {

	mes := &Messages1{
		Type: "test",
		Body: "haha",
	}

	log.Printf("%s,%s", mes.Body, mes.Type)
}
