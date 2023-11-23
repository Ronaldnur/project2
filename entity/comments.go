package entity

import "time"

type Comment struct {
	Id         int
	User_id    int
	Photo_id   int
	Message    string
	Created_at time.Time
	Updated_at time.Time
}
