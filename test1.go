package main

import "fmt"

type Mes struct {
	Name    string
	InChan  chan string
	OutChan chan string
}

func main() {
	mes := Mes{
		Name:    "mes",
		InChan:  make(chan string),
		OutChan: nil,
	}
	go func() {
		for i := 0; i < 10; i++ {
			fmt.Printf("go son ...%d\n", i)
		}
		mes.InChan <- "sdfsd"
	}()

	//for i:= range mes.InChan{
	//	fmt.Printf("%s",i)
	//}
	res := <-mes.InChan
	fmt.Println(res)

}
