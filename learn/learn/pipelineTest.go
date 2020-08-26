package main

import (
	"fmt"
	"time"
)

func main() {
	naturals := make(chan int)
	squares := make(chan int)

	//counter
	go func() {
		for x := 0; x < 5; x++ {
			naturals <- x
		}
		close(naturals)
	}()

	//squarer
	go func() {
		for x := range naturals {
			squares <- x * x
		}
		close(squares)
	}()

	for {
		time.Sleep(1 * time.Second)
		fmt.Println(<-squares)
	}

}
