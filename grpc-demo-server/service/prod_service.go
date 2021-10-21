package service

import "golang.org/x/net/context"

type ProdService struct {
}

func (p *ProdService) GetProdInfo(ctx context.Context, req *ProdRequest) (*ProdModel, error) {
	ret := &ProdModel{ProdId: req.ProdId, ProdName: "苹果", ProdPrice: 5499.99}
	return ret, nil
}

func (p *ProdService) GetProdStocks(ctx context.Context, size *QuerySize) (*ProdResponseList, error) {
	var prods []*ProdResponse
	for i := 1; i <= int(size.Size); i++ {
		prods = append(prods, &ProdResponse{ProdStock: int32(30 + i)})
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
