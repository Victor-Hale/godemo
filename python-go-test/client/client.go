package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"python-go-test/pb"
)

//const address  = "139.196.106.241:50051"
const address  = "127.0.0.1:50051"

func main()  {
	conn,err := grpc.Dial(address,grpc.WithInsecure())
	if err != nil {
		fmt.Println(err)
	}
	defer conn.Close()

	c := pb.NewGreeterClient(conn)
	name := "world"
	res,err := c.SayHello(context.Background(),&pb.HelloRequest{Name: name})

	if err != nil{
		fmt.Println(err)
	}
	fmt.Println(res.Msg)
}