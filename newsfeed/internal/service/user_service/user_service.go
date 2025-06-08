package user_service

import (
	"context"
	"log"

	"ep.k14/newsfeed/internal/service/model"
)

type UserService struct {
	// dai
}

func New() (*UserService, error) {
	return &UserService{}, nil
}

func (s *UserService) Signup(ctx context.Context, user *model.User) (*model.User, error) {
	log.Println("service: create user successfully")
	return user, nil
}
