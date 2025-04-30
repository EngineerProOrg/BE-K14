package sharedmodels

import "time"

type UserBaseViewModel struct {
	UserId    int64     `json:"userId" binding:"required"`
	FirstName string    `json:"firstName" binding:"required"`
	LastName  string    `json:"lastName" binding:"required"`
	Name      string    `json:"fullName" binding:"required"`
	Birthday  time.Time `json:"birthday" binding:"required"`
	Email     string    `json:"email" binding:"required"`
	Avatar    *string   `json:"avatar"`
}
