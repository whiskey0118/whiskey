package main

import (
	"flag"
	"fmt"
)

func main() {
	var a string
	flag.StringVar(&a, "config", "haha", "hahaa")
	flag.Parse()
	fmt.Println(a)
}
