package service

import (
	"context"
	"fmt"
	"time"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"

	"github.com/devhg/grpc-demo/grpc-demo-server/helper"
)

type OrderService struct {
	auth *helper.Auth
}

func (o OrderService) GetOrderInfo(ctx context.Context, request *OrderRequest) (*OrderMain, error) {
	time.Sleep(3 * time.Second)

	if ctx.Err() == context.Canceled {
		return nil, status.Errorf(codes.Canceled, "SearchService.Search canceled")
	}

	return &OrderMain{
		OrderId:    request.OrderId,
		OrderNo:    "12313",
		UserId:     123,
		OrderMoney: 10,
		OrderTime:  timestamppb.Now(),
		Details:    []*OrderDetail{{OrderNo: "12313", DetailId: 1}},
	}, nil
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
