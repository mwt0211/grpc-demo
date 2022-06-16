package service

import (
	"context"
	"fmt"
	"io"
	"log"
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
		fmt.Println("服务端接收到的流", recv.ProdId)
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
func (p *productService) mustEmbedUnimplementedProdServiceServer() {}
