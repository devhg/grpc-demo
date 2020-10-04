package service

import (
	"context"
	"fmt"
)

type OrderService struct {
}

func (o OrderService) NewOrder(ctx context.Context, or *OrderRequest) (*OrderResponse, error) {
	fmt.Println(or.OrderMain)
	fmt.Println(or.OrderMain.Details)
	ret := &OrderResponse{Status: "200", Message: "ok!"}
	return ret, nil
}
