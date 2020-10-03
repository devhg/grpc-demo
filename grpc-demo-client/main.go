package main

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"github.com/QXQZX/grpc-demo/grpc-demo-client/service"
	"github.com/golang/protobuf/ptypes/timestamp"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"io/ioutil"
	"log"
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
	prodResponse, err := prodClient.GetProdService(context.Background(), &service.ProdRequest{ProdId: 12, ProdArea: 1})
	stocks, err := prodClient.GetProdStocks(context.Background(), &service.QuerySize{Size: 3})
	prodInfo, err := prodClient.GetProdInfo(context.Background(), &service.ProdRequest{ProdId: 1222})

	// 订单服务
	orderClient := service.NewOrderServiceClient(conn)

	order, err := orderClient.NewOrder(context.Background(),
		&service.OrderMain{
			OrderId:    11,
			OrderNo:    "20201003",
			OrderMoney: 1.1,
			UserId:     233,
			OrderTime:  &timestamp.Timestamp{Seconds: time.Now().Unix()},
		})

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
}
