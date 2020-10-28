package main

import (
	"fmt"
	"time"
)

func main() {
	b := make(chan int)
	b <- 22
	go func() {
		c := <-b
		fmt.Println("c:", c)
		fmt.Println("sub goroutine is gone")
	}()
	time.Sleep(1 * time.Second)
	fmt.Println("main is gone")
}
