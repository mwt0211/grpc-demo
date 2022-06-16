package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"log"
	"mwt.com/grpc/demo/service"
	"net"
)

//1.取出server
type helloService struct {
	service.UnimplementedHelloServiceServer
}

//2.实现挂载方法
func (h *helloService) SayHello(ctx context.Context, req *service.Req) (res *service.Res, err error) {
	//实现proto的方法
	msg := req.GetMsg()
	fmt.Println("request中的msg=", msg)
	return &service.Res{Msg: msg}, nil
}
func main() {

	//3.注册服务
	listen, err := net.Listen("tcp", ":8089")
	if err != nil {
		log.Fatal("注册监听失败", err)
	}
	s := grpc.NewServer()
	service.RegisterHelloServiceServer(s, &helloService{})
	//4.建立监听
	s.Serve(listen)

}
