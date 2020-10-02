protoc --go_out=plugins=grpc:../ prod.proto
protoc --grpc-gateway_out=logtostderr=true:../ prod.proto