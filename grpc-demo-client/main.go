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
	prodResponse, err := client.GetProdService(context.Background(), &service.ProdRequest{ProdId: 12})

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(prodResponse)
}
