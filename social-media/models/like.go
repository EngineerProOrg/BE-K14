package models

import "time"

type Like struct {
	Id        int
	PostId    int
	UserId    int
	CreatedAt time.Time
	UpdatedAt *time.Time
}
