package main

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"github.com/QXQZX/grpc-demo/grpc-demo-client/service"
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

	client := service.NewProdServiceClient(conn)
	prodResponse, err := client.GetProdService(context.Background(),
		&service.ProdRequest{ProdId: 12, ProdArea: 1})
	stocks, err := client.GetProdStocks(context.Background(), &service.QuerySize{Size: 3})

	prodInfo, err := client.GetProdInfo(context.Background(), &service.ProdRequest{ProdId: 1222})

	if err != nil {
		log.Fatal(err)
	}

	// 普通类型
	fmt.Println(prodResponse)
	// 数组类型
	fmt.Println("stocks", stocks.Prods)
	fmt.Println("stocks", stocks.Prods[2].ProdStock)

	fmt.Println(prodInfo)
}
