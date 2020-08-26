package main

import "fmt"

func main() {
	fmt.Println(mySqrt(6))

}
func mySqrt(x int) int {
	var (
		left  int
		mid   int
		right int
	)
	left = 0
	right = x - 1
	if x == 1 {
		return 1
	}
	for left < x {
		right = mid
		mid = (left + right) / 2

	}

	return mid

}
