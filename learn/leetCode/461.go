package main

import (
	"fmt"
)

//通过

func main() {
	c := hammingDistance(3, 1)
	fmt.Println(c)
}

func hammingDistance(x int, y int) int {
	if x == y {
		return 0
	} else {
		a := fmt.Sprintf("%b", x^y)
		b := string(a)
		j := 0
		for _, i := range b {
			if string(i) == "1" {
				j++
			}
		}
		return j
	}
}
