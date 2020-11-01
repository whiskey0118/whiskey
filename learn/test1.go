package main

import (
	"fmt"
)

func main() {
	a := "a"
	b := "b"
	fmt.Println(&a)
	fmt.Println(&b)
	fmt.Println(a < b)

}
