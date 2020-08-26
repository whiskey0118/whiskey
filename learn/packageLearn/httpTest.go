//http package learn
package main

import (
	"fmt"
)

func main() {
	//http.HandleFunc("/", nil)
	//
	//err := http.ListenAndServe("0.0.0.0:8080", nil)
	//if err != nil {
	//	log.Fatal(err)
	//}

	uMax := ^uint(0)          //无符号整数
	max := int(^uint(0) >> 1) //(有符号整数)
	fmt.Println("uMax:", uMax)
	fmt.Println("max:", max)
	fmt.Printf("%#x\n", uint(0))
	fmt.Printf("%#x\n", ^uint(0))

	a := "a"
	fmt.Printf("%v\n", a[0])
	fmt.Printf("%#x\n", a[0])
	fmt.Printf("%#x", ^a[0])

}

type Test interface {
	start()
}

type haha struct {
	name string
}

func (c *haha) start() {
	fmt.Println("haha")
}
