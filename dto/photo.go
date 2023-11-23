package dto

import "time"

type NewPhotoRequest struct {
	Title     string `json:"title" valid:"required"`
	Caption   string `json:"caption"`
	Photo_url string `json:"photo_url" valid:"required"`
}

type MakeDataPhoto struct {
	Id         int       `json:"id"`
	Title      string    `json:"title"`
	Caption    string    `json:"caption"`
	Photo_url  string    `json:"photo_url"`
	User_id    int       `json:"user_id"`
	Created_at time.Time `json:"created_at"`
}

type NewPhotoResponse struct {
	Result     string        `json:"result"`
	StatusCode int           `json:"statusCode"`
	Message    string        `json:"message"`
	Data       MakeDataPhoto `json:"data"`
}

type NewGetPhotoRequest struct {
	Id         int       `json:"id"`
	Title      string    `json:"title"`
	Caption    string    `json:"caption"`
	Photo_url  string    `json:"photo_url"`
	User_id    int       `json:"user_id"`
	Created_at time.Time `json:"created_at"`
	Updated_at time.Time `json:"updated_at"`
	User       GetUser   `json:"User"`
}

type GetPhotoResponse struct {
	Result     string               `json:"result"`
	StatusCode int                  `json:"statusCode"`
	Message    string               `json:"message"`
	Data       []NewGetPhotoRequest `json:"data"`
}
type MakeUpdatePhoto struct {
	Id         int       `json:"id"`
	Title      string    `json:"title"`
	Caption    string    `json:"caption"`
	Photo_url  string    `json:"photo_url"`
	User_id    int       `json:"user_id"`
	Updated_at time.Time `json:"updated_at"`
}

type NewUpdateResponse struct {
	Result     string          `json:"result"`
	StatusCode int             `json:"statusCode"`
	Message    string          `json:"message"`
	Data       MakeUpdatePhoto `json:"data"`
}

type DeletePhotoResponse struct {
	StatusCode int    `json:"statusCode"`
	Message    string `json:"message"`
}

type GetPhotoForComment struct {
	Id        int    `json:"id"`
	Title     string `json:"title"`
	Caption   string `json:"caption"`
	Photo_url string `json:"photo_url"`
	User_id   int    `json:"user_id"`
}
