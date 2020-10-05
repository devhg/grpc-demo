package service

import (
	"time"
)

type UserScoreService struct {
}

func (u UserScoreService) GetUserScore(in *UserScoreRequest,
	stream UserScoreService_GetUserScoreServer) error {
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
