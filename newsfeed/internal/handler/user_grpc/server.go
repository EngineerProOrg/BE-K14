package user_grpc

import (
	"context"
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/protobuf/proto"

	"ep.k14/newsfeed/internal/handler/proto/user"
	"ep.k14/newsfeed/internal/service/model"
)

type UserService interface {
	Signup(ctx context.Context, user *model.User) (*model.User, error)
}

type UserGrpcServer struct {
	grpcServer *grpc.Server
}

func New(userService UserService) (*UserGrpcServer, error) {
	s := &UserGrpcServer{}

	// init handler
	userHandler := &userGrpcHandler{
		userService: userService,
	}

	// register handler into grpc server
	grpcServer := grpc.NewServer()
	user.RegisterUserServiceServer(grpcServer, userHandler)

	s.grpcServer = grpcServer

	return s, nil
}

func (s *UserGrpcServer) Start() error {
	// open port
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", 8081))
	if err != nil {
		fmt.Println("failed to listen", err)
		return fmt.Errorf("failed to listen port %s", "8081")
	}

	// server listen from port
	fmt.Println("user grpc server starting to serve on port 8081 ...")
	err = s.grpcServer.Serve(lis)
	if err != nil {
		fmt.Println("failed to serve on port", err)
		return fmt.Errorf("failed to serve port %s", "8081")
	}

	return nil
}

type userGrpcHandler struct {
	user.UnimplementedUserServiceServer

	userService UserService
}

func (h *userGrpcHandler) Signup(ctx context.Context, req *user.SignupRequest) (*user.SignupResponse, error) {
	log.Println("receive request", req)

	// service
	userModel, err := h.userService.Signup(ctx, &model.User{
		Username:    req.GetUserName(),
		Password:    req.GetPassword(),
		DisplayName: req.GetDisplayName(),
	})
	if err != nil {
		fmt.Println("err when service signup user", err)
		return nil, fmt.Errorf("failed to create user: %s", err)
	}

	resp := &user.SignupResponse{
		UserName:    proto.String(userModel.Username),
		DisplayName: proto.String(userModel.DisplayName),
	}
	log.Println("return response", resp)
	return resp, nil
}
