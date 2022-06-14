package main

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"grpc/demo/Client/auth"
	"grpc/demo/service"
	"io/ioutil"
	"log"
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
	conn, err := grpc.Dial(":8099", grpc.WithTransportCredentials(creds), grpc.WithPerRPCCredentials(Token))
	if err != nil {
		log.Fatal(err)
	}

	// 退出时关闭链接
	defer conn.Close()

	// 2. 调用Product.pb.go中的NewProdServiceClient方法
	//productServiceClient := service.NewProdServiceClient(conn)
	productServiceClient := service.NewProdServiceClient(conn)
	// 3. 直接像调用本地方法一样调用GetProductStock方法
	req := &service.ProductRequest{ProdId: 488}
	resp, err := productServiceClient.GetProductStock(context.Background(), req)
	if err != nil {
		log.Fatal("调用gRPC方法错误: ", err)
	}

	fmt.Println("调用gRPC方法成功，库存ProdStock为 :", resp.ProdStock)
}
