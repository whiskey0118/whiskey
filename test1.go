package main

import (
	"encoding/json"
	"fmt"
)

type Messages struct {
	Name string
	Body string
}

//func main() {
//	var mes Message
//	mes.name = "test"
//	res,_ := json.Marshal(&mes)
//	fmt.Println(string(res))
//	fmt.Printf("%s",string(res))
//}

type Product struct {
	Name      string
	ProductID int64
}

func main() {
	var p Product
	p.Name = "Xiao mi 6"
	p.ProductID = 1
	data, _ := json.Marshal(&p)
	fmt.Println(string(data))

	var mes Messages
	mes.Name = "test"
	res, _ := json.Marshal(&mes)
	fmt.Println(string(res))
}
