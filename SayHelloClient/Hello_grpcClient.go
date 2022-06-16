package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"log"
	"mwt.com/grpc/demo/service"
)

func main() {
	//1.创建连接
	conn, err := grpc.Dial(":8088", grpc.WithInsecure())
	if err != nil {
		log.Fatal("客户端链接失败")
	}
	//2.创建Client
	helloServiceClient := service.NewHelloServiceClient(conn)
	//3.调用Client的方法
	DateFromClient, err2 := helloServiceClient.SayHello(context.Background(), &service.Req{
		Msg: "从客户端来的数据",
	})
	if err2 != nil {
		log.Fatal("客户端调用方法失败", err2)
	}
	fmt.Println(DateFromClient)
}
