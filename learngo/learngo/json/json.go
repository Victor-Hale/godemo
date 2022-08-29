package main

import (
	"encoding/json"
	"fmt"
)

type itam struct {
  Id string
  Writer string
}

type books struct {
	Name string `json:"name,omitempty"`
	Price string `json:"price,omitempty"`
	Num   string `json:"num,omitempty"`
	Item  []itam
}


func main() {
	book := books{
		//tag标签
		Name: "长津湖",
		Price: "15元",
		Num: "25个",
		Item:[]itam{
			{
				Id:"1",
				Writer: "稳住好",
			},
			{
				Id:"2",
				Writer: "曾俊",
			},
		},
	}
	booker, err := json.Marshal(book)
	if err !=nil{
		panic(err)
	}
	fmt.Println(string(booker))
}

