package models

import "time"

type Post struct {
	Id        int
	UserId    int
	Content   string
	CreatedAt time.Time
	UpdatedAt *time.Time
}
