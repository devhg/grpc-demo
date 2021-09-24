package helper

import (
	"crypto/tls"
	"crypto/x509"
	"io/ioutil"
	
	"google.golang.org/grpc/credentials"
)

// GetServerCreds 服务端
func GetServerCreds() credentials.TransportCredentials {
	/*
		// TLS证书  (自签证书)
		creds, err := credentials.NewServerTLSFromFile("./keys/server.crt",
			"./keys/server_no_password.key")
	
		if err != nil {
			log.Fatal(err)
		}
	*/
	
	cert, _ := tls.LoadX509KeyPair("cert/server.pem", "cert/server.key")
	certPool := x509.NewCertPool()
	
	ca, _ := ioutil.ReadFile("cert/ca.pem")
	certPool.AppendCertsFromPEM(ca)
	
	creds := credentials.NewTLS(&tls.Config{
		Certificates: []tls.Certificate{cert}, // 服务端证书
		ClientAuth:   tls.RequireAndVerifyClientCert,
		ClientCAs:    certPool,
	})
	return creds
}

// GetClientCreds 客户端
func GetClientCreds() credentials.TransportCredentials {
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
	return creds
}
