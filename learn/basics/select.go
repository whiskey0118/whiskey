package main

import "fmt"

func main() {
	//var ch1,ch2 chan int
	//a:=1
	//b:=1
	//select {
	//case a = <-ch1:
	//	fmt.Println("cha1",a)
	//case ch2 <- b:
	//	fmt.Println("cha2",b)
	//default:
	//	fmt.Println("xixi")
	//}

	c1 := make(chan int)

	var i1 int
	select {
	case i1 = <-c1:
		fmt.Printf("received ", i1, " from c1\n")

	//default:
	//	fmt.Printf("no communication\n")
	}
}
