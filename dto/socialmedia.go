package dto

import "time"

type NewSocialMediaRequest struct {
	Name             string `json:"name" valid:"required"`
	Social_media_url string `json:"social_media_url" valid:"required"`
}

type SocialMediaReturn struct {
	Id               int       `json:"id"`
	Name             string    `json:"name"`
	Social_media_url string    `json:"social_media_url"`
	User_id          int       `json:"user_id"`
	Created_at       time.Time `json:"created_at"`
}

type SocialMediaResponse struct {
	Result     string            `json:"result"`
	StatusCode int               `json:"statusCode"`
	Message    string            `json:"message"`
	Data       SocialMediaReturn `json:"data"`
}

type GetSocialMedia struct {
	Id               int                   `json:"id"`
	Name             string                `json:"name"`
	Social_media_url string                `json:"social_media_url"`
	User_id          int                   `json:"user_id"`
	Created_at       time.Time             `json:"created_at"`
	Updated_at       time.Time             `json:"updated_at"`
	User             GetUserForSocialMedia `json:"User"`
}

type GetSocialMediaResponse struct {
	Result     string           `json:"result"`
	StatusCode int              `json:"statusCode"`
	Message    string           `json:"message"`
	Data       []GetSocialMedia `json:"social_medias"`
}

type UpdateSocialMediaReturn struct {
	Id               int       `json:"id"`
	Name             string    `json:"name" valid:"required"`
	Social_media_url string    `json:"social_media_url" valid:"required"`
	User_id          int       `json:"user_id"`
	Updated_at       time.Time `json:"created_at"`
}

type UpdateSocialMediaResponse struct {
	Result     string                  `json:"result"`
	StatusCode int                     `json:"statusCode"`
	Message    string                  `json:"message"`
	Data       UpdateSocialMediaReturn `json:"data"`
}

type DeleteResponseSocialMedia struct {
	StatusCode int    `json:"statusCode"`
	Message    string `json:"message"`
}
