package main

import "fmt"

type User struct {
	Name string
	Age  string
}

func (U User) Changeanme(newName string){
	U.Name  =newName
}
func (U*User) Changname( NewName string){
	U.Name = NewName
}

func main() {
	  u := User{
	  	Name: "jirry",
	  }
	  up:=&User{
	  	Name: "BOb",
	  }
	  //u.Changeanme("TOM")
	  //up.Changname("BB")
	  u.Changname("TOM")
	  up.Changeanme("BB")
	  fmt.Printf("%v",u)
	  fmt.Printf("%v",up)
}
