package main

import "fmt"

type String1 string

func (d String1) Len() int {
	return len(d)
}
func main() {
	lru := New(int64(0), nil)
	lru.Add("key1",String1("1234"))
	v,_:=lru.Get("key1")
	fmt.Println(v)
}
