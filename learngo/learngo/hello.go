package main

import (
	"fmt"
	"io/ioutil"
)

 var (
 	q int = 1
    t int =3
  )


func chansghi()  {
	var ame1 string = "qwe"
    jieguo,error:=ioutil.ReadFile(ame1)
    if ame1 == "qwe"{
    fmt.Println(jieguo)
    } else{
       fmt.Println(error)
    }


}

func test()(arr[5] int){
	arr1 :=[5] int {7,85,6,4,5}
	return arr1
}


func slice(arr[]int)(value []int){
	   slice :=arr[2:4]
	return slice

}

func mapqwe()(map[string]string){
	m :=map[string]string{
		"name" : "文助豪",
		"age"  : "21" ,
	}//key [vulue]
	return m

}
func main() {
	//arr :=[] int {0,2,3,4,5}
	//fmt.Println("hello world")
	//chansghi()
	//fmt.Println(q,t)
	//var bingo=test()
	//for i,v :=range bingo{
	//	fmt.Println(i,v)
	//}
	//fmt.Println(slice(arr))
	v :=mapqwe()
	name := v["name"]
	fmt.Println(name)
}
