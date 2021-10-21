package main

import (
	"context"
	"fmt"
	"github.com/devhg/grpc-demo/grpc-demo-client/helper"
	"github.com/devhg/grpc-demo/grpc-demo-client/service"
	"github.com/golang/protobuf/ptypes/timestamp"
	"google.golang.org/grpc"
	"io"
	"log"
	"testing"
	"time"
)

// 商品服务
func TestProdService(t *testing.T) {
	conn, err := grpc.Dial(":9305", grpc.WithTransportCredentials(helper.GetClientCreds()))
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	// 商品服务
	prodClient := service.NewProdServiceClient(conn)

	prodResponse, err := prodClient.GetProdService(
		context.Background(),
		&service.ProdRequest{
			ProdId:   12,
			ProdArea: service.ProdAreas_B,
		},
	)

	stocks, err := prodClient.GetProdStocks(
		context.Background(),
		&service.QuerySize{Size: 3},
	)

	prodInfo, err := prodClient.GetProdInfo(
		context.Background(),
		&service.ProdRequest{ProdId: 1222},
	)

	// 普通类型
	fmt.Println("prodResponse =", prodResponse)
	// 数组类型
	fmt.Println("prodResponseList =", stocks.Prods)
	fmt.Println("prodResponseList.Prods[2].ProdStock =", stocks.Prods[2].ProdStock)
	fmt.Println("prodInfo =", prodInfo)
}

// 订单服务
func TestOrderService(t *testing.T) {
	conn, err := grpc.Dial(":9305", grpc.WithTransportCredentials(helper.GetClientCreds()))
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	// 订单服务
	orderClient := service.NewOrderServiceClient(conn)

	orderResp, err := orderClient.NewOrder(
		context.Background(),
		&service.OrderRequest{OrderMain: &service.OrderMain{
			OrderId:    11,
			OrderNo:    "20201003",
			OrderMoney: 111,
			UserId:     233,
			OrderTime:  &timestamp.Timestamp{Seconds: time.Now().Unix()},
			Details: []*service.OrderDetail{
				{OrderNo: "10001", DetailId: 101},
				{OrderNo: "10002", DetailId: 102},
			},
		}},
	)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(orderResp)
}

// 从服务端流模式接收
func TestGetUserScoreByServerStream(t *testing.T) {
	conn, err := grpc.Dial(":9305", grpc.WithTransportCredentials(helper.GetClientCreds()))
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	// 从服务流模式接收
	userScoreClient := service.NewUserScoreServiceClient(conn)

	req := service.UserScoreRequest{}
	for i := 1; i < 6; i++ {
		req.Users = append(req.Users, &service.UserScore{UserId: int32(i)})
	}

	stream, err := userScoreClient.GetUserScoreByServerStream(context.Background(), &req)
	if err != nil {
		log.Fatal(err)
	}

	for {
		res, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}

		fmt.Println(res.Users)
	}
}

// 向服务端流模式发送
func TestGetUserScoreByClientStream(t *testing.T) {
	conn, err := grpc.Dial(":9305", grpc.WithTransportCredentials(helper.GetClientCreds()))
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	// 向服务端流模式发送
	userScoreClient := service.NewUserScoreServiceClient(conn)
	stream, err := userScoreClient.GetUserScoreByClientStream(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	for j := 0; j < 3; j++ {
		req := &service.UserScoreRequest{}
		req.Users = make([]*service.UserScore, 0)

		// 这里好比客户端一个比较耗时的过程
		for i := 1; i < 6; i++ {
			req.Users = append(req.Users, &service.UserScore{UserId: int32(i)})
		}

		err := stream.Send(req)
		if err != nil {
			log.Println(err)
		}
	}

	resp, _ := stream.CloseAndRecv()
	fmt.Println(resp.Users)
}

// 双向流模式
func TestGetUserScoreByStream(t *testing.T) {
	conn, err := grpc.Dial(":9305", grpc.WithTransportCredentials(helper.GetClientCreds()))
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	// 双向流模式
	userScoreClient := service.NewUserScoreServiceClient(conn)
	stream, err := userScoreClient.GetUserScoreByStream(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	var uid int32 = 1
	for j := 0; j < 3; j++ {
		req := &service.UserScoreRequest{}

		// 这里好比客户端一个比较耗时的过程
		for i := 1; i < 6; i++ {
			req.Users = append(req.Users, &service.UserScore{UserId: uid})
			uid++
		}

		err := stream.Send(req)
		if err != nil {
			log.Println(err)
		}

		resp, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Println(err)
		}
		fmt.Println(resp.Users)
	}
}
