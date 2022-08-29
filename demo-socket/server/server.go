package main

import (
	"fmt"
	"net"
)

func main() {
	listen, _ := net.Listen("tcp","localhost:8000")
	defer listen.Close()
	for  {
		conn , _ := listen.Accept()
		fmt.Println(conn.RemoteAddr(),"用户连接成功！！！")

		go func() {
			var Date = []byte("123")
			fmt.Println("准备写入")
			conn.Write(Date)
			fmt.Println("写入完毕")
		}()
	}

}
