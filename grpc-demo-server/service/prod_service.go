package service

import "golang.org/x/net/context"

type ProdService struct {
}

func (p *ProdService) GetProdService(context.Context, *ProdRequest) (*ProdResponse, error) {
	return &ProdResponse{ProdStock: 25}, nil
}
