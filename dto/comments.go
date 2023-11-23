package dto

import "time"

type NewCommentRequest struct {
	Message  string `json:"message" valid:"required"`
	Photo_id int    `json:"photo_id" valid:"required"`
}

type MakeDataComment struct {
	Id         int       `json:"id"`
	Message    string    `json:"message"`
	Photo_id   int       `json:"photo_id"`
	User_id    int       `json:"user_id"`
	Created_at time.Time `json:"created_at"`
}

type CommentResponse struct {
	Result     string `json:"result"`
	StatusCode int    `json:"statusCode"`
	Message    string `json:"message"`
	Data       MakeDataComment
}

type GetCommentsRequest struct {
	Id         int                `json:"id"`
	Message    string             `json:"message"`
	Photo_id   int                `json:"photo_id"`
	User_id    int                `json:"user_id"`
	Updated_at time.Time          `json:"updated_at"`
	Created_at time.Time          `json:"created_at"`
	User       GetUserForComment  `json:"User"`
	Photo      GetPhotoForComment `json:"Photo"`
}

type GetCommentsResponse struct {
	Result     string               `json:"result"`
	StatusCode int                  `json:"statusCode"`
	Message    string               `json:"message"`
	Data       []GetCommentsRequest `json:"data"`
}

type MakeCommentUpdate struct {
	Message string `json:"message" valid:"required"`
}

type UpdateComment struct {
	Id         int       `json:"id"`
	Message    string    `json:"message"`
	Photo_id   int       `json:"photo_id"`
	User_id    int       `json:"user_id"`
	Updated_at time.Time `json:"updated_at"`
}

type UpdateResponse struct {
	Result     string        `json:"result"`
	StatusCode int           `json:"statusCode"`
	Message    string        `json:"message"`
	Data       UpdateComment `json:"data"`
}

type DeleteCommentResponse struct {
	StatusCode int    `json:"statusCode"`
	Message    string `json:"message"`
}
