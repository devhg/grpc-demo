package main

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"time"
	
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	
	"github.com/devhg/grpc-demo/grpc-demo-client/service"
	"github.com/golang/protobuf/ptypes/timestamp"
)

func main() {
	
	/*  自签证书
	creds, err := credentials.NewClientTLSFromFile("keys/server.crt", "devhui.org")
	
	if err != nil {
		log.Fatal(err)
	}
	*/
	
	cert, _ := tls.LoadX509KeyPair("cert/client.pem", "cert/client.key")
	certPool := x509.NewCertPool()
	
	ca, _ := ioutil.ReadFile("cert/ca.pem")
	certPool.AppendCertsFromPEM(ca)
	
	creds := credentials.NewTLS(&tls.Config{
		Certificates: []tls.Certificate{cert},
		ServerName:   "localhost",
		RootCAs:      certPool,
	})
	
	conn, err := grpc.Dial(":9305", grpc.WithTransportCredentials(creds))
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
			ProdArea: 1,
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
	
	// 订单服务
	orderClient := service.NewOrderServiceClient(conn)
	
	order, err := orderClient.NewOrder(
		context.Background(),
		&service.OrderRequest{OrderMain: &service.OrderMain{
			OrderId:    11,
			OrderNo:    "20201003",
			OrderMoney: 111,
			UserId:     233,
			OrderTime:  &timestamp.Timestamp{Seconds: time.Now().Unix()},
		}},
	)
	
	if err != nil {
		log.Fatal(err)
	}
	
	// 普通类型
	fmt.Println(prodResponse)
	// 数组类型
	fmt.Println("stocks", stocks.Prods)
	fmt.Println("stocks", stocks.Prods[2].ProdStock)
	fmt.Println(prodInfo)
	fmt.Println(order)
	
	/*
		// 从服务流模式接收
		userScoreClient := service.NewUserScoreServiceClient(conn)
		req := service.UserScoreRequest{}
		req.Users = make([]*service.UserScore, 0)
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
	*/
	/*
		// 向服务端流模式发送
		userScoreClient := service.NewUserScoreServiceClient(conn)
		stream, err := userScoreClient.GetUserScoreByClientStream(context.Background())
	
		if err != nil {
			log.Fatal(err)
		}
	
		for j := 0; j < 3; j++ {
			req := service.UserScoreRequest{}
			req.Users = make([]*service.UserScore, 0)
	
			// 这里好比客户端一个比较耗时的过程
			for i := 1; i < 6; i++ {
				req.Users = append(req.Users, &service.UserScore{UserId: int32(i)})
			}
			err := stream.Send(&req)
			if err != nil {
				log.Println(err)
			}
		}
	
		recv, _ := stream.CloseAndRecv()
		fmt.Println(recv.Users)
	*/
	
	// 双向流模式
	userScoreClient := service.NewUserScoreServiceClient(conn)
	stream, err := userScoreClient.GetUserScoreByStream(context.Background())
	
	if err != nil {
		log.Fatal(err)
	}
	
	var uid int32 = 1
	for j := 0; j < 3; j++ {
		req := service.UserScoreRequest{}
		req.Users = make([]*service.UserScore, 0)
		
		// 这里好比客户端一个比较耗时的过程
		for i := 1; i < 6; i++ {
			req.Users = append(req.Users, &service.UserScore{UserId: uid})
			uid++
		}
		err := stream.Send(&req)
		if err != nil {
			log.Println(err)
		}
		
		recv, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Println(err)
		}
		fmt.Println(recv.Users)
	}
	
}
