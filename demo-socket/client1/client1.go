package main

import (
	"fmt"
	"net"
)

const (
	HOST   = "localhost"
	PORT   = "8000"
)
func main() {
	conn , _ := net.Dial("tcp",HOST+":"+PORT)
	for  {
		Data := make([]byte, 255)
		conn.Read(Data)
		fmt.Println(string(Data))
	}

}
