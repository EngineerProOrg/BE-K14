package main

import (
	"github.com/gin-gonic/gin"
)

var (
	ip = "localhost"
	port = "6379"
)

contextRedis := context.Background()
redisClient := redis.NewClient(
	&redis.Options{
		Addr: ip+":"+port, //todo: need change
		Password: "",				//password redis
		DB: 0,					//db redis				
	}
)

type user struct {
	username string `json:"username"`
	sessionId string `json:"sessionId"`
}

func main() {
	r := gin.Default()

	groupApi := r.Group("/api")
	groupApi.GET("/login", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "login",
		})
	})
	r.Run(":8080")
}
