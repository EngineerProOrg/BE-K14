package main

import (
	"fmt"
	"log"

	"ep.k14/newsfeed/internal/dai/user_dai"
	"ep.k14/newsfeed/internal/handler/user_grpc"
	"ep.k14/newsfeed/internal/service/user_service"
)

func main() {
	// create db conn -> db access object
	userDai, err := user_dai.New(&user_dai.UserDbConfig{
		Username:     "root",
		Password:     "123456",
		Host:         "localhost",
		Port:         3306,
		DatabaseName: "newsfeed",
	})
	if err != nil {
		log.Println("err init user dai", err)
		return
	}

	userService, err := user_service.New(userDai)
	if err != nil {
		log.Println("err init user service", err)
		return
	}

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
