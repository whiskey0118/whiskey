//分布式id递增

package main

import "fmt"

func main() {
	a := New(2, 3)
	b := New(2, 3)
	fmt.Println("id is :", a.Id())
	fmt.Println("id is :", a.Id())
	fmt.Println("id is :", b.Id())
	fmt.Println("id is :", b.Id())
}

//创建一个AutoInc自增数的结构体
type AutoInc struct {
	start   int
	step    int
	queue   chan int
	running bool
}

func New(start, step int) (ai *AutoInc) {
	ai = &AutoInc{
		start:   start,
		step:    step,
		running: true,
		queue:   make(chan int, 1),
	}

	go ai.process()
	return
}

//AutoInc结构体的方法，当ai.running是true的时候，循环增加ai.start的数字
//并将i的值传递给ai.queue通道中
func (ai *AutoInc) process() {
	defer func() { recover() }()
	for i := ai.start; ai.running; i = i + ai.step {
		ai.queue <- i
	}
}

//ai.queue通道输出，函数并返回通道的值
func (ai *AutoInc) Id() int {
	return <-ai.queue
}

//将ai.running值设置为false，那么结构体方法process()将不会自增，并关闭通道
//当前没有调用close方法，就是说这个ai.process()会一直在运行状态，通道一直存在
func (ai *AutoInc) Close() {
	ai.running = false
	close(ai.queue)
}
