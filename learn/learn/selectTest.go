//8.7
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
		time.Sleep(3 * time.Second)
		ch1 <- 3
	}()

	go func() {
		time.Sleep(3 * time.Second)
	}()

	fmt.Println("blocking on read")
	select {
	case <-c:
		fmt.Printf("unblocked %v later \n", time.Since(start))
	case <-ch1:
		fmt.Println("ch1 case...")
	case <-ch2:
		fmt.Println("ch2 case...")
		//default:
		//	fmt.Println("default go...")
	}

}

func test1() {
	a := time.NewTicker(time.Second)
	defer a.Stop()
	//done := make(chan struct{})
	//select {
	//case -> a:
	//
	//}
	t := <-a.C
	fmt.Printf("%T\n", t)
	fmt.Printf("%s", t.String())

}
