package main

import (
	"net"
	"net/rpc"
)

type HolleServer struct {

}

func (s *HolleServer)Hello(request string , reply *string)error{
	*reply = "holle ," +request
	//为什么这个地方使用指针，因为我们监听了这个服务有没有被调用成功返回error。
	//我们使用指针可以直接通过地址去拿到我们的返回值。
	return nil
}
func main() {
	//1.经典三步走 , 实例化一个server
	listener ,_ := net.Listen("tcp",":1234")
	//2.经典三步走 ， 注册处理逻辑headler
	_ = rpc.RegisterName("HelloService", &HolleServer{})
	//3.经典三步走 ， 启动服务
     conn ,_ := listener.Accept()   //当一个新的连接创建的时候进行套接字传递
     rpc.ServeConn(conn)

}
