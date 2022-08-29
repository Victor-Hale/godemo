
package main

import "fmt"

func testa(){
	fmt.Println("aaaaaaa")
}

func testb(x int){
	defer func() {
           err:= recover()
		   if err!=nil {
			   fmt.Println(err)
		   }
	}()
	//设置recover,
	//defer func(){
	//	//recover()可以打印panic的错误信息
	//	//fmt.Println(recover())
	//	if err := recover(); err != nil{ //产生了panic异常
	//		fmt.Println(err)
	//	}
	//}()
	var a [10]int
	a[x] = 111


}

func testc(){
	fmt.Println("ccccccc")
}
func main(){
	testa()
	testb(10)
	testc()
}
