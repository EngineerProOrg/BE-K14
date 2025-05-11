package sharedmodels

import "time"

type UserBaseViewModel struct {
	UserId    int64     `json:"userId"`
	FirstName string    `json:"firstName" binding:"required,notblank"`
	LastName  string    `json:"lastName" binding:"required,notblank"`
	Name      string    `json:"fullName"`
	Birthday  time.Time `json:"birthday" binding:"required"`
	Email     string    `json:"email" binding:"required,notblank,email"`
	Username  string    `json:"username"`
	Avatar    *string   `json:"avatar"`
}
