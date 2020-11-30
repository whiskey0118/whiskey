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
	time string
	int64
}

func main() {
	var p Messages
	//p.Name = "Xiao mi 6"
	//p.ProductID = 1
	//data, _ := json.Marshal(&p)
	//fmt.Println(string(data))
	//
	//var mes Messages
	//mes.Name = "test"
	//res, _ := json.Marshal(&mes)
	//fmt.Println(string(res))

	t := `{"name":"client1","body":"test"}`
	fmt.Println(t)
	err := json.Unmarshal([]byte(t), &p)
	if err != nil {
		fmt.Printf("err:", err)
	}
	fmt.Println(p)

}
