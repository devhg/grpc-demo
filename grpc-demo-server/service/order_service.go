package service

import (
	"context"
	"fmt"
)

type OrderService struct {
}

func (o OrderService) NewOrder(ctx context.Context, om *OrderMain) (*OrderResponse, error) {
	fmt.Println(om)
	ret := &OrderResponse{Status: "200", Message: "ok!"}
	return ret, nil
}
