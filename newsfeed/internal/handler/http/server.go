package http

import (
	"context"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/protobuf/proto"

	"ep.k14/newsfeed/internal/handler/proto/user"
)

type HttpServer struct {
	router *gin.Engine

	userGrpcClient user.UserServiceClient
}

func New() (*HttpServer, error) {
	router := gin.Default()

	h := &HttpServer{}

	conn, err := grpc.NewClient("localhost:8081", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		fmt.Println("err init user grpc client:", err)
		return nil, err
	}
	// defer conn.Close()
	userGrpcCli := user.NewUserServiceClient(conn)

	userRouter := router.Group("/user")
	userRouter.POST("/signup", h.Signup)
	userRouter.POST("/login", h.Login)
	// TODO:
	// add

	h.router = router
	h.userGrpcClient = userGrpcCli

	return h, nil
}

func (h *HttpServer) Start() error {
	err := h.router.Run() // listen and serve on 0.0.0.0:8080
	return err
}

type SignupRequest struct {
	Username    string `json:"user_name"`
	Password    string `json:"password"`
	Email       string `json:"email"`
	DisplayName string `json:"display_name"`
	Dob         string `json:"dob"`
}

type UserData struct {
	Username    string `json:"user_name"`
	Email       string `json:"email"`
	DisplayName string `json:"display_name"`
	Dob         string `json:"dob"`
}

func (h *HttpServer) Signup(c *gin.Context) {
	req := &SignupRequest{}

	if err := c.ShouldBind(req); err != nil {
		fmt.Println("err when parse signup request", err)
		c.JSON(http.StatusBadRequest, &CommonResponse{
			Code:    CodeInvalidRequest,
			Message: err.Error(), // TODO: custom error msg
			Data:    nil,
		})
		return
	}

	grpcReq := &user.SignupRequest{
		UserName:    proto.String(req.Username),
		Password:    proto.String(req.Password),
		DisplayName: proto.String(req.DisplayName),
	}
	ctx := context.Background()
	_, err := h.userGrpcClient.Signup(ctx, grpcReq)
	if err != nil {
		fmt.Println("err when call signup grpc", err)
		c.JSON(http.StatusBadRequest, &CommonResponse{
			Code:    CodeGrpcCall,
			Message: err.Error(), // TODO: custom error msg
			Data:    nil,
		})
		return
	}

	userData := &UserData{
		Username:    req.Username,
		Email:       req.Email,
		DisplayName: req.DisplayName,
		Dob:         req.Dob,
	}

	c.JSON(http.StatusOK, &CommonResponse{
		Code:    CodeSuccess,
		Message: "success",
		Data:    userData,
	})
}

func (h *HttpServer) Login(c *gin.Context) {
}
