package main

import "fmt"

func main() {
	a := 10
	b := 20
	a = a ^ b
	b = a ^ b
	a = a ^ b
	fmt.Println("a:", a, "b:", b)
}

func test(nums []int) int {
	a := 0
	for i := range nums {
		a ^= i
	}
	return a
}
