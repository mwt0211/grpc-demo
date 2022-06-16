package main

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"io"
	"io/ioutil"
	"log"
	"mwt.com/grpc/demo/Client/auth"
	"mwt.com/grpc/demo/service"
	"time"
)

func main() {
	// 1. 新建连接，端口是服务端开放的8002端口
	// 没有证书会报错
	/*	//单向认证添加证书
		creds, err2 := credentials.NewClientTLSFromFile("cert/server.pem", "*.mwt.com")
		if err2 != nil {
			log.Fatal("添加证书失败", err2)
			fmt.Printf("err:%s", err2)
		}*/
	// 证书认证-双向认证
	// 从证书相关文件中读取和解析信息，得到证书公钥、密钥对
	cert, _ := tls.LoadX509KeyPair("cert/client.pem", "cert/client.key")
	// 创建一个新的、空的 CertPool
	certPool := x509.NewCertPool()
	ca, _ := ioutil.ReadFile("cert/ca.crt")
	// 尝试解析所传入的 PEM 编码的证书。如果解析成功会将其加到 CertPool 中，便于后面的使用
	certPool.AppendCertsFromPEM(ca)
	// 构建基于 TLS 的 TransportCredentials 选项
	creds := credentials.NewTLS(&tls.Config{
		// 设置证书链，允许包含一个或多个
		Certificates: []tls.Certificate{cert},
		// 要求必须校验客户端的证书。可以根据实际情况选用以下参数
		ServerName: "*.mwt.com",
		RootCAs:    certPool,
	})

	//无认证状态的调用
	//conn, err := grpc.Dial(":8099", grpc.WithTransportCredentials(insecure.NewCredentials()))
	//添加证书之后的调用
	//grpc.WithPerRPCCredentials()完成token的认证
	Token := &auth.Authentication{
		Password: "admin",
		User:     "admin",
	}
	conn, err := grpc.Dial(":8098", grpc.WithTransportCredentials(creds), grpc.WithPerRPCCredentials(Token))
	if err != nil {
		log.Fatal(err)
	}

	// 退出时关闭链接
	defer conn.Close()

	// 2. 调用Product.pb.go中的NewProdServiceClient方法
	//productServiceClient := service.NewProdServiceClient(conn)
	productServiceClient := service.NewProdServiceClient(conn)
	// 3. 直接像调用本地方法一样调用GetProductStock方法
	//req := &service.ProductRequest{ProdId: 511}
	/*	resp, err := productServiceClient.GetProductStock(context.Background(), req)
		if err != nil {
			log.Fatal("调用gRPC方法错误: ", err)
		}

		fmt.Println("调用gRPC方法成功，库存ProdStock为 :", resp.ProdStock, resp.User)*/
	/*stream, err := productServiceClient.UpdateProductStockClientStream(context.Background())
	if err != nil {
		log.Fatal("获取流失败: ", err)
	}
	//客户端源源不断的发送数据
	resp := make(chan struct{}, 1)
	go ProdRequest(stream, resp)
	select {
	case <-resp:
		recv, err := stream.CloseAndRecv()
		if err != nil {
			log.Fatal("接收并关闭流失败", err)
		}
		stock := recv.ProdStock
		fmt.Println("客户端收到响应", stock, recv.User)
	}*/
	//客户端发送一次请求,服务端不断发送请求
	req := &service.ProductRequest{ProdId: 166}
	stream, err := productServiceClient.GetProductStockServerStream(context.Background(), req)
	if err != nil {
		log.Fatal("客户端发送一次请求,服务端不断接收请求失败", err)
	}
	//
	for {
		recv, err := stream.Recv()
		if err != nil {
			if err == io.EOF {
				fmt.Println("客户端接收服务端的请求完毕")
				stream.CloseSend()
				break
			}
			log.Fatal("客户端接受请求失败", err)
		}
		fmt.Println("客户端接收服务端发送的流", recv.ProdStock)
	}

}

func ProdRequest(stream service.ProdService_UpdateProductStockClientStreamClient, resp chan struct{}) {
	//源源不断的发送数据
	count := 0
	for {

		req := &service.ProductRequest{ProdId: 166}
		err := stream.Send(req)
		if err != nil {
			log.Fatal("发送流数据失败", err)
		}
		//隔五秒发送一次
		time.Sleep(time.Second * 2)
		count++
		fmt.Printf("第%d次发送\n", count)
		if count > 10 {
			//数据发送完成
			resp <- struct{}{}
			break
		}
	}
}
