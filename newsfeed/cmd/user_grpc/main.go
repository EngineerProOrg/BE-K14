package main

import (
	"fmt"

	"ep.k14/newsfeed/internal/handler/user_grpc"
	"ep.k14/newsfeed/internal/service/user_service"
)

func main() {
	// create db conn -> db access object
	userService, _ := user_service.New()

	userGrpcServer, err := user_grpc.New(userService)
	if err != nil {
		fmt.Println("err init user grpc server", err)
		return
	}

	err = userGrpcServer.Start()
	if err != nil {
		fmt.Println("err start user grpc server", err)
		return
	}
}
