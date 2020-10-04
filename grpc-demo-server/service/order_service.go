package service

import (
	"context"
	"fmt"
)

type OrderService struct {
}

func (o OrderService) NewOrder(ctx context.Context, or *OrderRequest) (*OrderResponse, error) {
	err := or.OrderMain.Validate()
	if err != nil {
		return &OrderResponse{Status: "error", Message: err.Error()}, nil
	}
	fmt.Println(or.OrderMain)
	fmt.Println(or.OrderMain.Details)
	ret := &OrderResponse{Status: "200", Message: "ok!"}
	return ret, nil
}
