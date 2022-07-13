package service

import (
	"context"
	"fmt"
)

var Pro = &productService{}

type productService struct {
}

func (p *productService) mustEmbedUnimplementedProdServiceServer() {
	//TODO implement me
	fmt.Println("mustEmbedUnimplementedProdServiceServer")
	panic("please implement me")
}

/* GetProductStock 需要一个载体来实现 grpc pb文件内的接口, 也是业务逻辑实现的地方
type ProdServiceServer interface {
	// 服务具体实现方法
	GetProductStock(context.Context, *ProductRequest) (*ProductResponse, error)
	mustEmbedUnimplementedProdServiceServer()
}
*/

func (p *productService) GetProductStock(ctx context.Context, request *ProductRequest) (*ProductResponse, error) {
	return &ProductResponse{
		ProStock: request.ProId,
	}, nil
}
