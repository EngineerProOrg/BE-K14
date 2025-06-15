package models

import "time"

type User struct {
	Id        int64
	Name      string    `json:"name"`
	Email     string    `json:"email" binding:"required"`
	Password  string    `json:"password" binding:"required"`
	CreatedAt time.Time `json:"createdAt"`
}
