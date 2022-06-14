package main

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
	"grpc/demo/service"
	"io/ioutil"
	"log"
	"net"
)

func main() {
	//添加证书
	/*	//单向认证
		//creds, err2 := credentials.NewServerTLSFromFile(公钥, 私钥)
		creds, err2 := credentials.NewServerTLSFromFile("cert/server.pem", "cert/server.key")
		if err2!=nil{
			log.Fatal("添加证书失败", err2)
			fmt.Printf("err:%s", err2)
		}*/

	// 证书认证-双向认证
	// 从证书相关文件中读取和解析信息，得到证书公钥、密钥对
	cert, err2 := tls.LoadX509KeyPair("cert/server.pem", "cert/server.key")
	if err2 != nil {
		log.Fatal("添加证书失败", err2)
		fmt.Printf("err:%s", err2)
	}
	// 创建一个新的、空的 CertPool
	certPool := x509.NewCertPool()
	ca, err := ioutil.ReadFile("cert/ca.crt")
	if err != nil {
		log.Fatal("ca证书读取错误", err)
	}
	// 尝试解析所传入的 PEM 编码的证书。如果解析成功会将其加到 CertPool 中，便于后面的使用
	certPool.AppendCertsFromPEM(ca)
	// 构建基于 TLS 的 TransportCredentials 选项
	creds := credentials.NewTLS(&tls.Config{
		// 设置证书链，允许包含一个或多个
		Certificates: []tls.Certificate{cert},
		// 要求必须校验客户端的证书。可以根据实际情况选用以下参数
		ClientAuth: tls.RequireAndVerifyClientCert,
		// 设置根证书的集合，校验方式使用 ClientAuth 中设定的模式
		ClientCAs: certPool,
	})
	//实现token认证,需要合法的用户名和密码
	//需要拦截器
	var authInterceptor grpc.UnaryServerInterceptor
	//type UnaryServerInterceptor func(ctx context.Context, req interface{}, info *UnaryServerInfo, handler UnaryHandler) (resp interface{}, err error)
	authInterceptor = func(
		ctx context.Context,
		req interface{},
		info *grpc.UnaryServerInfo,
		handler grpc.UnaryHandler) (resp interface{}, err error) {
		//实现拦截普通方法请求,验证Token
		err = Auth(ctx)
		if err != nil {
			return
		}
		//无err处理请求
		return handler(ctx, req)

	}
	server := grpc.NewServer(grpc.Creds(creds), grpc.UnaryInterceptor(authInterceptor))
	//注册
	service.RegisterProdServiceServer(server, service.ProductService)
	listen, err := net.Listen("tcp", ":8099")
	if err != nil {
		log.Fatal("启动监听失败", err)
		fmt.Printf("err:%s", err)
	}
	err = server.Serve(listen)
	if err != nil {

		log.Fatal("启动服务失败", err)
		fmt.Printf("err:%s", err)

	}
	fmt.Println("启动product_grpcServer端成功")
}

func Auth(ctx context.Context) error {
	//拿到传输的用户名和密码

	//func FromIncomingContext(ctx context.Context) (MD, bool)
	//type MD map[string][]string
	md, ok := metadata.FromIncomingContext(ctx)

	if !ok {
		return fmt.Errorf("missing credentials")
	}
	var user string
	var password string

	if val, ok := md["user"]; ok {
		user = val[0]
	}
	if val, ok := md["password"]; ok {
		password = val[0]
	}

	if !(user == "admin" && password == "admin") {
		return status.Errorf(codes.Unauthenticated, "token不合法")
	}
	return nil
}
