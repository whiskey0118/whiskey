package main

import (
	"fmt"
)

func main() {

	a := "11001"
	for i, j := range a {
		fmt.Println(i, string(j))
	}
	fmt.Println([]byte(a))
}
