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

func (p *ProdService) GetProdService(ctx context.Context, req *ProdRequest) (*ProdResponse, error) {
	stock := 10
	if req.ProdArea == ProdAreas_A {
		stock = 11
	} else if req.ProdArea == ProdAreas_B {
		stock = 22
	} else if req.ProdArea == ProdAreas_C {
		stock = 33
	}
	return &ProdResponse{ProdStock: int32(stock)}, nil
}
