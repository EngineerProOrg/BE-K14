package models

import "time"

type Comment struct {
	Id        int
	PostId    int
	UserId    int
	Content   string
	CreatedAt time.Time
	UpdatedAt *time.Time //  *time.Time allows nullable
}
