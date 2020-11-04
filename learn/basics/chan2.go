package main

import (
	"fmt"
	"time"
)

func main() {
	start := time.Now()
	c := make(chan interface{})
	ch1 := make(chan int)
	ch2 := make(chan int)
	go func() {
		time.Sleep(4 * time.Second)
		close(c)
	}()

	go func() {
		//time.Sleep(5*time.Second)
		ch1 <- 3
	}()

	go func() {
		//time.Sleep(5*time.Second)
		ch2 <- 5
	}()

	fmt.Println("read blocking")
	select {
	case <-c:
		fmt.Printf("unblocked %v later.\n", time.Since(start))
	case <-ch1:
		fmt.Println("ch1 case...")
	case <-ch2:
		fmt.Println("ch2 case ...")
	default:
		fmt.Println("default go ...")
	}
}
