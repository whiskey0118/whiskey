package main

import (
	"fmt"
)

func main() {
	done := make(chan int)

	go func() {
		done <- 2
		fmt.Println("go down")
		a := <-done
		fmt.Println(a)
		fmt.Println("xixixoxox")
		done <- 22
		fmt.Println("xixi")
	}()

	select {
	case <-done:
		fmt.Println("haha")
	}

}
