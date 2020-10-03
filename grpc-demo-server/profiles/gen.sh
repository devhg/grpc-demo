protoc --go_out=plugins=grpc:../ prod.proto
protoc --go_out=plugins=grpc:../ models.proto
protoc --grpc-gateway_out=logtostderr=true:../ prod.proto
protoc --grpc-gateway_out=logtostderr=true:../ models.proto