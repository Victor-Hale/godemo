package main

import (
	"fmt"
)

func work (choo chan bool){
	fmt.Println("正在完成工作")
     choo <- true
}
func main() {

	// 构建一个通道
	ch := make(chan int)
	// 开启一个并发匿名函数
	go func() {
		fmt.Println("start goroutine")
		// 通过通道通知main的goroutine
		ch <- 1
		fmt.Println("exit goroutine")

	}()
//	time.Sleep(2*time.Second)
	fmt.Println("wait goroutine")
//	time.Sleep(2*time.Second)
	// 等待匿名goroutine
	<- ch
//	v, _ := <- ch
	fmt.Println("all done")
//	fmt.Println(v)

     choo := make(chan bool)
	 go work(choo)
	 <-choo
	 fmt.Println("任务完成")
}
