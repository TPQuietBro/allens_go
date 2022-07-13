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
}

func (p *productService) GetProductStock(ctx context.Context, request *ProductRequest) (*ProductResponse, error) {
	return &ProductResponse{
		ProStock: request.ProId,
	}, nil
}
