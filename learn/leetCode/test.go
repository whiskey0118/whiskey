package main

import "fmt"

func main() {

	a := []int{1, 2, 3}
	fmt.Printf("%d\n", a[len(a)-1])
	for i := 0; i < len(a); i++ {
		fmt.Println(a[i])
	}
}
