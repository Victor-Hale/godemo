package main

import (
	"fmt"
	"time"
)

func main() {
	var a[10] int
	for i:=0;i<10;i++ {
		go func(i int){
			for  {
			fmt.Println("helllo"+"horeirh %d",i)
			}
		}(i)//末尾的括号表明匿名函数被调用，并将返回的函数指针赋给变量a
	}
	time.Sleep(time.Minute)
	fmt.Println(a)
}
