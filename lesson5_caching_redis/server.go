package main

import (
	"context"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	redis "github.com/go-redis/redis/v8"
	"time"
)

var (
	ip   = "localhost"
	port = "6379"
)

var contextRedis = context.Background()
var redisClient = redis.NewClient(
	&redis.Options{
		Addr:     ip + ":" + port, //todo: need change
		Password: "",              //password redis
		DB:       0,               //db redis
	},
)

type user struct {
	Username  string `json:"username"`
	SessionId string `json:"sessionId"`
}

func login(c *gin.Context) {
	username := c.Query("username")
	sessionId := c.Query("sessionId")
	err := redisClient.Set(contextRedis, sessionId, username, time.Hour).Err()
	if err != nil {
		fmt.Println("redis set error", err)
		return
	}
}

func getUserInfo(c *gin.Context) {
	key := c.Query("sessionId")
	value, err := redisClient.Get(contextRedis, key).Result()
	if err != nil {
		fmt.Println("redis get error", err)
		return
	} else if errors.Is(err, redis.Nil) {
		fmt.Printf("Key '%s' does not exist\n", key)
	} else {
		fmt.Printf("value: '%s' for key: '%s'", key, value)
	}
}

func main() {
	r := gin.Default()

	groupApi := r.Group("/api")
	groupApi.GET("/login", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "login",
		})
		login(c)
	})

	groupApi.GET("/show-info", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "show data of redis",
		})
		getUserInfo(c)
	})
	err := r.Run(":8080")
	if err != nil {
		return
	}
}
