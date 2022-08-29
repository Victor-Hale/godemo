package main

import "fmt"

func add() func(int) int {
 sum :=0
 return func(i int) int {
	 sum += i
	 return sum
 }
}

func main() {
	a:=add()
	for i :=0;i<=5;i++ {
		fmt.Println(a(i))
	}
}
