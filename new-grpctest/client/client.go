package main

import (
	"context"
	"encoding/json"
	"fmt"
	"google.golang.org/grpc"
	"net/http"
	Fss "new-grpctest/pb"
)
type JsonResult  struct{
	Code int `json:"code"`
	Msg  string `json:"msg"`
	Date string `json:"date"`
}
const address  = "localhost:50051"
var imgdir string
var red string
func handler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")             //允许访问所有域
	w.Header().Add("Access-Control-Allow-Headers", "Content-Type") //header的类型
	w.Header().Set("content-type", "application/json")             //返回数据格式是json

	imgdir =r.FormValue("password")
    fmt.Println(string(imgdir))
	conn,err := grpc.Dial(address,grpc.WithInsecure())
	if err != nil {
		fmt.Println(err)
	}
	defer conn.Close()

	c := Fss.NewGreeterClient(conn)
	//imgdir ="22"
	res, err := c.SayHello(context.Background(), &Fss.HelloRequest{Imgdir:imgdir})

	for k , v :=range res.String() {
		if v == '?' {
			red = res.String()[5:k]
		}
	}
	msg, _ := json.Marshal(JsonResult{Code: 200, Msg: "验证成功",Date: red})
	res.String()
	if err != nil{
		fmt.Println(err)
	}
	fmt.Println(res,"成功")
	w.Write(msg)

}

func main()  {
	    fmt.Println("服务开启......等待连接")
		http.HandleFunc("/hello", handler)
		http.ListenAndServe(":8080", nil)


	//conn,err := grpc.Dial(address,grpc.WithInsecure())
	//if err != nil {
	//	fmt.Println(err)
	//}
	//defer conn.Close()
	//
	//c := Fss.NewGreeterClient(conn)
	////imgdir ="22"
	//res, err := c.SayHello(context.Background(), &Fss.HelloRequest{Imgdir:imgdir})
	//result=res.String()
	//if err != nil{
	//	fmt.Println(err)
	//}
	//fmt.Println(res,"成功")
}