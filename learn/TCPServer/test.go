package main

import "fmt"

func main() {
	var a string
	var c int
	b, err := fmt.Scanln(&a, &c)
	if err != nil {
		fmt.Println("err", err)
	}
	fmt.Println("b:", b, "a:", a, "c:", c)
}
