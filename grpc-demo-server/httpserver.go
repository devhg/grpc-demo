package main

import (
	"context"
	"github.com/QXQZX/grpc-demo/grpc-demo-server/helper"
	"github.com/QXQZX/grpc-demo/grpc-demo-server/service"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"google.golang.org/grpc"
	"log"
	"net/http"
)

func main() {
	//由grpc-gateway实现同时支持http和grpc服务
	//grpc-gateway为http提供代理访问grpc
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	// Register gRPC server endpoint
	// Note: Make sure the gRPC server is running properly and accessible
	mux := runtime.NewServeMux()
	opts := []grpc.DialOption{grpc.WithTransportCredentials(helper.GetClientCreds())}

	// 商品服务
	err := service.RegisterProdServiceHandlerFromEndpoint(ctx, mux,
		"localhost:9305", opts) // 此地址对应grpc的端口

	if err != nil {
		log.Fatal(err)
	}

	// 订单服务
	err = service.RegisterOrderServiceHandlerFromEndpoint(ctx, mux,
		"localhost:9305", opts) // 此地址对应grpc的端口
	if err != nil {
		log.Fatal(err)
	}

	//server := &http.Server{Addr: ":8081", Handler: mux}
	//server.ListenAndServe()

	// Start HTTP server (and proxy calls to gRPC server endpoint)
	log.Fatal(http.ListenAndServe(":8081", mux))
}
