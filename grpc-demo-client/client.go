package main

import (
	"context"
	"fmt"
	"github.com/devhg/grpc-demo/grpc-demo-client/service"
	"google.golang.org/grpc"
	"io"
	"log"

	"github.com/devhg/grpc-demo/grpc-demo-client/helper"
)

func main() {
	conn, err := grpc.Dial(":9305", grpc.WithTransportCredentials(helper.GetClientCreds()))
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	// 双向流模式
	userScoreClient := service.NewUserScoreServiceClient(conn)
	stream, err := userScoreClient.GetUserScoreByStream(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	var uid int32 = 1
	for j := 0; j < 3; j++ {
		req := &service.UserScoreRequest{}

		// 这里好比客户端一个比较耗时的过程
		for i := 1; i < 6; i++ {
			req.Users = append(req.Users, &service.UserScore{UserId: uid})
			uid++
		}

		err := stream.Send(req)
		if err != nil {
			log.Println(err)
		}

		resp, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Println(err)
		}
		fmt.Println(resp.Users)
	}
}
