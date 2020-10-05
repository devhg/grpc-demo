package service

import (
	"io"
	"time"
)

type UserScoreService struct {
}

func (*UserScoreService) GetUserScoreByServerStream(in *UserScoreRequest,
	stream UserScoreService_GetUserScoreByServerStreamServer) error {
	var score int32 = 101
	users := make([]*UserScore, 0)
	for i, user := range in.Users {
		user.Score = score
		score++
		users = append(users, user)

		if (i+1)%2 == 0 && i > 0 {
			err := stream.Send(&UserScoreResponse{Users: users})
			if err != nil {
				return err
			}
			users = (users)[0:0]
			time.Sleep(2 * time.Second)
		}
	}
	if len(users) > 0 {
		err := stream.Send(&UserScoreResponse{Users: users})
		if err != nil {
			return err
		}
	}

	return nil
}

func (*UserScoreService) GetUserScoreByClientStream(stream UserScoreService_GetUserScoreByClientStreamServer) error {
	var score int32 = 101
	users := make([]*UserScore, 0)

	for {
		req, err := stream.Recv()
		if err == io.EOF {
			// 接收 并业务处理完 返回并关闭
			return stream.SendAndClose(&UserScoreResponse{Users: users})
		}
		if err != nil {
			return err
		}

		// 这里是服务端业务处理
		for _, user := range req.Users {
			user.Score = score
			score++
			users = append(users, user)
		}
	}
}
