package main

import (
	"context"
	"errors"
	"fmt"
	"sync"
	"time"

	"github.com/gin-gonic/gin"
	redis "github.com/go-redis/redis/v8"
)

var (
	ip = "localhost"

	port         = "4949"
	contextRedis = context.Background()
	redisClient  = redis.NewClient(
		&redis.Options{
			Addr:     ip + ":" + port, //todo: need change
			Password: "",              //password redis
			DB:       0,               //db redis
		},
	)
	mutex sync.Mutex
)

type user struct {
	Username  string `json:"username"`
	SessionId string `json:"sessionId"`
	PingCount int64  `json:"pingCount"`
}

// todo: 1. login function to set sessionId and username
func login(c *gin.Context) *user {
	username := c.Query("username")
	sessionId := c.Query("sessionId")
	user := user{
		Username:  username,
		SessionId: sessionId,
	}
	fmt.Printf("set data with username: %s, sessionId: %s\n", user.Username, user.SessionId)
	//todo: save sessionId, username to redis
	err := redisClient.Set(contextRedis, sessionId, username, time.Hour).Err()
	if err != nil {
		fmt.Println("redis set error", err)
		return nil
	}
	return &user
}

// getUserInfo function to get user info from redis
func getUserInfo(c *gin.Context) *user {
	key := c.Query("sessionId")
	value, err := redisClient.Get(contextRedis, key).Result()
	user := &user{
		Username:  key,
		SessionId: value,
	}
	if err != nil {
		fmt.Println("redis get error", err)
		return nil
	} else if errors.Is(err, redis.Nil) {
		fmt.Printf("Key '%s' does not exist\n", key)
		return nil
	} else {
		fmt.Printf("value: '%s' for key: '%s'\n", key, value)
		return user
	}
}

// todo: 2. ping function to ping redis
// only with one person can call at one time
func ping() string {
	// ping redis
	mutex.Lock()
	pong, err := redisClient.Ping(contextRedis).Result()
	mutex.Unlock()
	if err != nil {
		fmt.Println("redis ping error", err)
		return ""
	}
	fmt.Println("redis ping result: ", pong)
	time.Sleep(time.Second * 5)
	return pong
}

// todo: 3. count each person call ping api
func (user *user) count() int64 {
	user.PingCount++
	return user.PingCount
}

func main() {
	r := gin.Default()
	fmt.Println("Start server ...")
	groupApi := r.Group("/api")
	groupApi.GET("/login", func(c *gin.Context) {
		user := login(c)
		c.JSON(200, gin.H{
			"username":  user.Username,
			"sessionId": user.SessionId,
		})
	})

	groupApi.GET("/show-info", func(c *gin.Context) {
		user := getUserInfo(c)
		c.JSON(200, gin.H{
			"username":  user.Username,
			"sessionId": user.SessionId,
		})
	})

	groupApi.GET("/ping", func(c *gin.Context) {
		pong := ping()
		c.JSON(200, gin.H{
			"pong": pong,
		})
	})

	err := r.Run(":8080")
	if err != nil {
		return
	}
}
