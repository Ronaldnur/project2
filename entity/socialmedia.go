package entity

import "time"

type SocialMedia struct {
	Id               int
	Name             string
	Social_media_url string
	User_id          int
	Created_at       time.Time
	Updated_at       time.Time
}
