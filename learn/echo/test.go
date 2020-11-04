package main

import (
	"fmt"
	"os"
	"os/signal"
)

func main() {

	c := make(chan os.Signal)
	signal.Notify(c)
	select {
	case <-c:
		fmt.Println("sdfsdf")
	}
}
