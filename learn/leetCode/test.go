package main

import "fmt"

func main() {
	a := 4
	b := 3
	a, b = b, a
	fmt.Println("a", a)
	fmt.Println("b", b)
	//
	//c := []int{1,2,3}
	//fmt.Println(c[2])
	//fmt.Println(len(c))
}
