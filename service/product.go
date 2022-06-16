package service

import (
	"context"
	"fmt"
	"io"
	"log"
	"time"
)

var ProductService = &productService{}

type productService struct {
}

func (p *productService) GetProductStock(c context.Context, request *ProductRequest) (*ProductResponse, error) {
	//实现具体的业务逻辑要求
	pwd := "mwt4111123"
	Idcard := "河南新乡"
	user := &User{Username: "张三", Password: &pwd, Age: 18, IdCard: &Idcard}
	stock := p.GetStockById(request.GetProdId())
	return &ProductResponse{ProdStock: stock, User: user}, nil
}
func (p *productService) UpdateProductStockClientStream(stream ProdService_UpdateProductStockClientStreamServer) error {
	count := 0
	pwd := "mwt4111123"
	Idcard := "河南新乡"
	for {
		//服务端源源不断的接收客户端发送的请求
		recv, err := stream.Recv()
		if err != nil {
			if err == io.EOF {
				return nil
			}
			log.Fatal("服务端接收请求失败,err=%v", err)
		}
		fmt.Println("服务端接收到的流", recv.ProdId, count)
		count++
		if count > 10 {
			productResponse := ProductResponse{ProdStock: recv.ProdId, User: &User{
				Username: "张三", Password: &pwd, Age: 18, IdCard: &Idcard,
			}}
			err := stream.SendAndClose(&productResponse)
			if err != nil {
				log.Fatal("接收完成但返回失败,err=%v", err)
			}
			return nil
		}
	}

	return nil
}
func (p *productService) GetStockById(id int32) int32 {
	return id
}

//服务端流
//客户端发送一次请求,服务端不断响应
func (p *productService) GetProductStockServerStream(request *ProductRequest, stream ProdService_GetProductStockServerStreamServer) error {
	count := 0
	for {
		p2 := &ProductResponse{ProdStock: request.ProdId}
		err := stream.Send(p2)
		if err != nil {
			log.Fatal("服务端发送请求失败", err)
		}
		time.Sleep(time.Second)
		count++
		if count > 10 {
			break
		}
	}
	return nil

}
func (p *productService) mustEmbedUnimplementedProdServiceServer() {}

//
