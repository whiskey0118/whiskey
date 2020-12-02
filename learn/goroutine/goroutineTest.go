package main

import (
	"fmt"
	"time"
)

type Test struct {
}

func main() {
	ch := make(chan Test)
	go func() {
		mes := Test{}
		ch <- mes
		time.Sleep(2 * time.Second)
		close(ch)
	}()

	for {
		select {
		case mes := <-ch:
			fmt.Println("rev mes: ", mes)
			fmt.Printf("mes type: %p", mes)
			time.Sleep(1 * time.Second)
		default:
			fmt.Println("default")
			time.Sleep(1 * time.Second)
		}
	}

	//time.Sleep(10*time.Second)
}
