package dto

import "time"

type NewUserRequest struct {
	Email    string `json:"email" valid:"email,required"`
	Username string `json:"username" valid:"required"`
	Password string `json:"password" valid:"required"`
	Age      int    `json:"age" valid:"required"`
}
type UserDataResponse struct {
	Age      int    `json:"age"`
	Email    string `json:"email"`
	Id       int    `json:"id"`
	Username string `json:"username"`
}
type NewUserResponse struct {
	Result     string           `json:"result"`
	StatusCode int              `json:"statusCode"`
	Message    string           `json:"message"`
	Data       UserDataResponse `json:"data"`
}
type NewUserLogin struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type NewUserUpdate struct {
	Email    string `json:"email" valid:"email,required"`
	Username string `json:"username" valid:"required"`
}
type MakeUserUpdate struct {
	Id         int       `json:"id"`
	Email      string    `json:"email"`
	Username   string    `json:"username"`
	Age        int       `json:"age"`
	Updated_at time.Time `json:"updated_at"`
}

type UserUpdateResponse struct {
	Result     string         `json:"result"`
	StatusCode int            `json:"statusCode"`
	Message    string         `json:"message"`
	Data       MakeUserUpdate `json:"data"`
}
type TokenResponse struct {
	Token string `json:"token"`
}

type LoginResponse struct {
	Result     string        `json:"result"`
	StatusCode int           `json:"statusCode"`
	Message    string        `json:"message"`
	Data       TokenResponse `json:"token"`
}

type DeleteResponse struct {
	StatusCode int    `json:"statusCode"`
	Message    string `json:"message"`
}

type GetUser struct {
	Email    string `json:"email"`
	Username string `json:"username"`
}

type GetUserForComment struct {
	Id       int    `json:"id"`
	Email    string `json:"email"`
	Username string `json:"username"`
}

type GetUserForSocialMedia struct {
	Id       int    `json:"id"`
	Username string `json:"username"`
}
