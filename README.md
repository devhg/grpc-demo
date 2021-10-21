# grpc-example

[TLS+CA认证](./conf)

证书过期需要重新生成

## show

rpc-server run
`go run server.go`

http-server proxy by grpc-gateway
`go run httpserver.go`

* http://127.0.0.1:8081/v1/prod/123123
* http://127.0.0.1:8081/v1/prodInfo/123123

run client
`go run client.go`
`go test client_test.go`
