//package main
//
//import (
//	"fmt"
//	"myself-chatroom/log"
//	"net"
//)
//
//var myMap = make(chan []byte , 2)
//
//type user struct {
//	conn net.Conn
//} //定义用户结构体
//
//var clientMap  =   make(map[string]net.Conn)
//
//func main() {
//	//1.页面监听
//	//2.接受请求
//	//3.处理逻辑(数据的传输与接收)
//	//4.关闭连接
//	var user user
//	listen, err := net.Listen("tcp","localhost:8000")
//	if err != nil {
//		log.Error.Println("连接错误,房间开启失败: ",err)
//		return
// }
//	defer listen.Close()
//	fmt.Println("聊天室开启成功！正在监听8000端口")
//	 //该go程负责从全局的消息channel中取出数据并发送数据
//
//	for{
//    conn , err := listen.Accept()
//		if err != nil {
//			log.Error.Println("用户连接失败了....",err)
//			return
//		}
//		//conn.RemoteAddr()    获取到用户的ip地址
//		clientMap[conn.RemoteAddr().String()] = conn
//
//		fmt.Println(conn.RemoteAddr(),"用户连接成功！！！")
//	    //开协程完成用户请求任务
//		fmt.Println(conn.RemoteAddr().String())
//		user.conn = conn
//		go handle(user.conn)
//		go manager(user.conn)
//	}
//
//
//}
//
// func handle(conn net.Conn){
//	  for  {
//		  by := make([]byte, 255)
//		  _,err := conn.Read(by)
//		  if err!=nil {
//			  return
//		  }
//		  myMap <- by
//		 // fmt.Println(conn.RemoteAddr(),"用户数据读取成功",string(by[0:leng]))
//	  }
// }
//
// func manager(conn net.Conn)  {
//	 for  {
//		 for _ ,conn :=range clientMap {
//			 var user = user{conn}
//			 m := <-myMap
//			 _, err := user.conn.Write(m)
//			 if err != nil {
//				 return
//			 }
//			 //fmt.Println(U.conn.RemoteAddr(),"用户数据写入成功",string(m[0:leng]))
//	 }
//	}
//}

package main

import (
"fmt"
"net"
"time"
)

// 客户端 map
var clientMap = make(map[string]*net.TCPConn) // 存储当前群聊中所有用户连接信息：key: ip+port, val: 用户连接信息

// 监听请求
func listenClient(ipAndPort string) {
	tcpAddr, _ := net.ResolveTCPAddr("tcp", ipAndPort)
	tcpListener, _ := net.ListenTCP("tcp", tcpAddr)
	for { // 循环接收
		clientConn, _ := tcpListener.AcceptTCP()                 // 监听请求连接
		clientMap[clientConn.RemoteAddr().String()] = clientConn // 将连接添加到 map
		go addReceiver(clientConn)
		fmt.Println("用户 : ", clientConn.RemoteAddr().String(), " 已连接.")
	}
}

// 向连接添加接收器
func addReceiver(newConnect *net.TCPConn) {
	for {
		byteMsg := make([]byte, 2048)
		len, err := newConnect.Read(byteMsg) // 从newConnect中读取信息到缓存中
		if err != nil {
			continue
		}
		fmt.Println(string(byteMsg[:len]))
		msgBroadcast(byteMsg[:len], newConnect.RemoteAddr().String())
	}
}

// 广播给所有 client
func msgBroadcast(byteMsg []byte, key string) {
	for k, con := range clientMap {
		if k != key { // 转发消息给当前群聊中，除自身以外的其他用户
			con.Write(byteMsg)
		}
	}
}

// 初始化
func initGroupChatServer() {
	fmt.Println("服务已启动...")
	time.Sleep(1 * time.Second)
	fmt.Println("等待客户端请求连接...")
	go listenClient("127.0.0.1:8000")
	select {}
}

func main() {
	initGroupChatServer()
}