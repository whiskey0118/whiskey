package main

import (
	"fmt"
	"log"
	"math/rand"
	"time"
)

func main() {
	//log.SetPrefix("<Read message>")
	////设置日志的格式
	//log.SetFlags(log.LstdFlags | log.Ldate | log.Ltime | log.Lshortfile)
	//log.Printf("this is a test")
	//http.HandleFunc("/", testFun)
	//http.ListenAndServe("127.0.0.1:8080", nil)

	TestGoWithChannel()
	da := []byte(string("bb"))
	data(da)
	da = []byte(string("cc"))
	var a rune
	fmt.Printf("%s", da)
}

func data(da []byte) error {
	da = []byte(string("aa"))
	fmt.Printf("func in %s", da)
	return nil
}

func testBoringWithChannel(msg string, c chan string) {
	for i := 0; ; i++ {
		c <- fmt.Sprintf("%s %d", msg, i)
		time.Sleep(time.Duration(rand.Intn(1e3)) * time.Millisecond)
	}
}

func TestGoWithChannel() {
	ch := make(chan string)
	go testBoringWithChannel("boring!", ch)
	for c := range ch {
		log.Printf("You say: %s", c)
	}
}
