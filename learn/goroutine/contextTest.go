package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	go watch(ctx, "one")
	go watch(ctx, "two")
	time.Sleep(5 * time.Second)
	cancel()
	fmt.Println("cancel 通知")
	time.Sleep(5 * time.Second)
}

func watch(ctx context.Context, name string) {
	for {
		select {
		case <-ctx.Done():
			fmt.Printf("%s stop \n", name)
			return
		default:
			fmt.Println(name, " running...")
			time.Sleep(1 * time.Second)
		}
	}
}
