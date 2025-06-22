package user_service

import (
	"context"
	"fmt"

	"ep.k14/newsfeed/internal/service/model"
)

type UserDAI interface {
	Signup(ctx context.Context, user *model.User) (*model.User, error)
}

type UserService struct {
	dai UserDAI
}

func New(userDai UserDAI) (*UserService, error) {
	return &UserService{
		dai: userDai,
	}, nil
}

func (s *UserService) Signup(ctx context.Context, user *model.User) (*model.User, error) {
	res, err := s.dai.Signup(ctx, user)
	if err != nil {
		fmt.Printf("err dai signup: %s\n", err)
		return nil, fmt.Errorf("err dai signup: %s", err)
	}
	user = res
	return user, nil
}
