package main

import (
	"fmt"
	"time"
)

//func main() {
//		ticker := time.NewTicker(time.Second)
//		ticker.Stop()
//		go func(t *time.Ticker) {
//			for {
//				<-ticker.C
//			}
//		}(ticker)
//		time.Sleep(5*time.Second)
//	//ch := make(chan int)
//	//go func() {
//	//	ch <- 0
//	//	fmt.Println("this is goroutine")
//	//}()
//	//a := <-ch
//	//fmt.Println(a)
//	//fmt.Println("this is main goroutine")
//
//}


func main() {
	//NewTimer 创建一个 Timer，它会在最少过去时间段 d 后到期，向其自身的 C 字段发送当时的时间

	//NewTicker 返回一个新的 Ticker，该 Ticker 包含一个通道字段，并会每隔时间段 d 就向该通道发送当时的时间。它会调
	//整时间间隔或者丢弃 tick 信息以适应反应慢的接收者。如果d <= 0会触发panic。关闭该 Ticker 可
	//以释放相关资源。
	ticker1 := time.NewTicker(1 * time.Second)
	defer ticker1.Stop()
	go func(t *time.Ticker) {
		for {
			<-t.C
			fmt.Println(time.Now().Format("2006-01-02 15:04:05"))
		}
	}(ticker1)

	time.Sleep(5*time.Second)
}