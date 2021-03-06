package main

import (
	"github.com/devhg/grpc-demo/grpc-demo-server/helper"
	"github.com/devhg/grpc-demo/grpc-demo-server/service"
	"google.golang.org/grpc"
	"net"
)

func main() {

	rpcServer := grpc.NewServer(grpc.Creds(helper.GetServerCreds()))

	service.RegisterProdServiceServer(rpcServer, new(service.ProdService))           // 商品服务
	service.RegisterOrderServiceServer(rpcServer, new(service.OrderService))         // 订单服务
	service.RegisterUserScoreServiceServer(rpcServer, new(service.UserScoreService)) // 积分服务

	// 使用tcp
	listen, _ := net.Listen("tcp", ":9305")

	rpcServer.Serve(listen)

	// 使用http

	//server := &http.Server{
	//	Addr:    ":9305",
	//	Handler: nil,
	//}
	//
	//http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
	//	fmt.Println(request)
	//	//writer.Write([]byte("122"))
	//	rpcServer.ServeHTTP(writer, request)
	//})

	//server.ListenAndServeTLS("keys/server.crt", "keys/server_no_password.key") // 自签证书
	//server.ListenAndServeTLS("cert/server.pem", "cert/server.key") // ca证书
}
