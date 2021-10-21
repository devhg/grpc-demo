# --proto_path 是proto的导入路径，路径规则是 ${--proto_path}/${import} 两者拼接起来
# --go_out 是go语言的输出目录，输出路径是 ${--go_out}/${option go_package=""} 两者拼接起来
# protoc --proto_path=./proto --go_out=./ proto/service/*.proto



# protoc --doc_out=. --proto_path=./proto --go_out=plugins=grpc:./ proto/service/*.proto

protoc --proto_path=./proto --go_out=plugins=grpc:./ proto/service/*.proto
protoc --proto_path=./proto --go_out=plugins=grpc:./ --validate_out=lang=go:./ proto/service/models.proto


protoc --proto_path=./proto --grpc-gateway_out=logtostderr=true:. proto/service/*.proto
#protoc --proto_path=./proto --grpc-gateway_out=logtostderr=true:. prod.proto
#protoc --proto_path=./proto --grpc-gateway_out=logtostderr=true:. models.proto
#protoc --proto_path=./proto --grpc-gateway_out=logtostderr=true:. order.proto
#protoc --proto_path=./proto --grpc-gateway_out=logtostderr=true:. user.proto