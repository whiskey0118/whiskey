package main

import (
	"fmt"
	"sync"
	"time"
)

//waitGroup 练习
func main() {
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		fmt.Println("1 goroutine sleep")
		time.Sleep(1 * time.Second)
		fmt.Println("1 goroutine stop")
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		fmt.Println("2 goroutine sleep")
		time.Sleep(1 * time.Second)
		fmt.Println("2 goroutine stop")
	}()

	fmt.Println("waiting for all goroutine ")
	time.Sleep(2 * time.Second)
	wg.Wait()
	fmt.Println("All goroutines finished!")

}
