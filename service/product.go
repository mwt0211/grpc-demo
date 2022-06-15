package service

import "context"

var ProductService = &productService{}

type productService struct {
}

func (p *productService) GetProductStock(c context.Context, request *ProductRequest) (*ProductResponse, error) {
	//实现具体的业务逻辑要求
	pwd := "mwt4111123"
	user := &User{Username: "张三", Password: &pwd}
	stock := p.GetStockById(request.GetProdId())
	return &ProductResponse{ProdStock: stock, User: user}, nil
}

func (p *productService) GetStockById(id int32) int32 {
	return id
}
func (p *productService) mustEmbedUnimplementedProdServiceServer() {}
