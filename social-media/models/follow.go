package models

import "time"

type Follow struct {
	Id          int
	FollowerId  int
	FollowingId int
	CreatedAt   time.Time
	UpdatedAt   *time.Time //  *time.Time allows nullable
}
