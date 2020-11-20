package main

import "log"

type Message struct {
	Name string
	In   chan string
}

func main() {

}

func test() {
	var mes Message
	select {
	case <-mes.In:
		log.Println("message in")

	}
}
