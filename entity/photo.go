package entity

import "time"

type Photo struct {
	Id         int
	Title      string
	Caption    string
	Photo_url  string
	User_id    int
	Created_at time.Time
	Updated_at time.Time
}
