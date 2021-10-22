package service

import (
	"context"
	"fmt"

	"github.com/devhg/grpc-demo/grpc-demo-server/helper"
)

type OrderService struct {
	auth *helper.Auth
}

func (o OrderService) NewOrder(ctx context.Context, or *OrderRequest) (*OrderResponse, error) {
	// checkout auth token
	// we can move it to interceptor
	if err := o.auth.Check(ctx); err != nil {
		return nil, err
	}

	err := or.OrderMain.Validate()
	if err != nil {
		return &OrderResponse{Status: "error", Message: err.Error()}, nil
	}

	fmt.Println(or.OrderMain)
	fmt.Println(or.OrderMain.Details)
	ret := &OrderResponse{Status: "200", Message: "ok!"}
	return ret, nil
}
