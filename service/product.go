package service

import "context"

var ProductService = &productService{}

type productService struct {
}

func (p *productService) GetProductStock(c context.Context, request *ProductRequest) (*ProductResponse, error) {
	//实现具体的业务逻辑要求

	stock := p.GetStockById(request.GetProdId())
	return &ProductResponse{ProdStock: stock}, nil
}

func (p *productService) GetStockById(id int32) int32 {
	return id
}
func (p *productService) mustEmbedUnimplementedProdServiceServer() {}
