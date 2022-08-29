package main

import "fmt"

// 定义一个接口


//type People interface {
//	ReturnName() string
//}
//type Student struct {
//	Name string
//}
////定义一个结构体方法
//func  (s Student)ReturnName() string {
//	return s.Name
//}

func sum(a[]int  ,  c chan int){
	sum:=0
	for _,b := range a{
		sum += b
	}
	c <- sum // 把 sum 发送到通道 c
}


func main() {
	//resp,err:=http.Get("https://www.imooc.c om")
	//if err !=nil{
	//	panic(err)
	//}
	//defer resp.Body.Close()
	//bytes,_:=ioutil.ReadAll(resp.Body)
	//fmt.Println(bytes)

	//cbs := Student{Name:"咖啡色的羊驼"}
	//var a People
	//a = cbs
	//name := a.ReturnName()
	//fmt.Println(name) // 输出"咖啡色的羊驼"

   a:=[]int{1,2,3,4,5,6}
   b:=[]int{7,8,9,10,11}
   c:=make(chan int)
   go sum(a,c)
   go sum(b,c)
   x,y := <- c, <-c
   fmt.Println(x,y)

}

