package service

import "golang.org/x/net/context"

type ProdService struct {
}

func (p *ProdService) GetProdStocks(context.Context, *QuerySize) (*ProdResponseList, error) {
	prods := []*ProdResponse{
		&ProdResponse{ProdStock: 31},
		&ProdResponse{ProdStock: 32},
		&ProdResponse{ProdStock: 33},
		&ProdResponse{ProdStock: 34},
	}
	return &ProdResponseList{Prods: prods}, nil
}

func (p *ProdService) GetProdService(context.Context, *ProdRequest) (*ProdResponse, error) {
	return &ProdResponse{ProdStock: 25}, nil
}
