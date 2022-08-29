//package main
//
//import (
//	"fmt"
//	"sync"
//	"time"
//)
//var m sync.Mutex
//var set = make(map[int]bool,0)
//
//func printonce(num int ){
//	m.Lock()
//	defer m.Unlock()
//	if
//	  _, exist := set[num];
//	 !exist {
//		fmt.Println(num)
//	}
//	set[num] = true
//}
//
//
//func main() {
//	for i := 0; i < 10; i++ {
//		go printonce(100)
//		go fmt.Println("1")
//		go fmt.Println("2")
//	}
//	time.Sleep(6*time.Second)
//}
package main

import (
	"flag"
	"fmt"
	"sync"
)

var (
	mutex   sync.Mutex //锁
	balance int    //余额
	protecting uint  // 是否加锁
	sign = make(chan struct{}, 10) //通道，用于等待所有goroutine
)

// 存钱
func deposit(value int) {
	defer func() {
		sign <- struct{}{}
	}()

	if protecting == 1 {
		mutex.Lock()
		defer mutex.Unlock()
	}

	fmt.Printf("余额: %d\n", balance)
	balance += value
	fmt.Printf("存 %d 后的余额: %d\n", value, balance)
	fmt.Println()

}

// 取钱
func withdraw(value int) {
	defer func() {
		sign <- struct{}{}
	}()

	if protecting == 1 {
		mutex.Lock()
		defer mutex.Unlock()
	}

	fmt.Printf("余额: %d\n", balance)
	balance -= value
	fmt.Printf("取 %d 后的余额: %d\n", value, balance)
	fmt.Println()

}

func main() {

	for i:=0; i < 5; i++ {
		go withdraw(500) // 取500
		go deposit(500)  // 存500
	}

	for i := 0; i < 10; i++ {
		<-sign
	}
	fmt.Printf("当前余额: %d\n", balance)
}

func init() {
	balance = 1000 // 初始账户余额为1000
	flag.UintVar(&protecting, "protecting", 0, "是否加锁，0表示不加锁，1表示加锁") //做一个初始化与赋值protecting默认为0
}
