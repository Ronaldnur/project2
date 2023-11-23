package service

import (
	"net/http"
	"project2/dto"
	"project2/entity"
	"project2/pkg/errs"
	"project2/pkg/helpers"
	"project2/repository/user_repository"
)

type userService struct {
	userRepo user_repository.Repository
}

type UserService interface {
	CreateNewUser(newUserRequsest dto.NewUserRequest) (*dto.NewUserResponse, errs.MessageErr)
	Login(payload dto.NewUserLogin) (*dto.LoginResponse, errs.MessageErr)
	UpdateUser(userId int, newUpdate dto.NewUserUpdate) (*dto.UserUpdateResponse, errs.MessageErr)
	DeleteUser(userId int) (*dto.DeleteResponse, errs.MessageErr)
}

func NewUserService(userRepo user_repository.Repository) UserService {
	return &userService{
		userRepo: userRepo,
	}
}

func (u *userService) CreateNewUser(newUserRequsest dto.NewUserRequest) (*dto.NewUserResponse, errs.MessageErr) {
	err := helpers.ValidateStruct(newUserRequsest)
	if len(newUserRequsest.Password) < 6 {
		return nil, errs.NewBadRequest("Password should be at least 6 characters long")
	}

	if newUserRequsest.Age <= 8 {
		return nil, errs.NewBadRequest("Age should be above 8 years")
	}

	if err != nil {
		return nil, err
	}
	existingUser, err := u.userRepo.GetUserByEmail(newUserRequsest.Email)
	if err != nil && err.Status() == http.StatusInternalServerError {
		return nil, err
	}

	if existingUser != nil {
		return nil, errs.NewDuplicateDataError("Please Try Another Email")
	}

	exitingUsername, err := u.userRepo.GetUserByUsername(newUserRequsest.Username)
	if err != nil && err.Status() == http.StatusInternalServerError {
		return nil, err
	}

	if exitingUsername != nil {
		return nil, errs.NewDuplicateDataError("Please Try Another Username")
	}

	user := entity.User{
		Username: newUserRequsest.Username,
		Email:    newUserRequsest.Email,
		Password: newUserRequsest.Password,
		Age:      newUserRequsest.Age,
	}
	err = user.HashPassword()
	if err != nil {
		return nil, err
	}
	userId, err := u.userRepo.CreateNewUser(user)

	if err != nil {
		return nil, err
	}
	response := dto.NewUserResponse{
		Result:     "success",
		StatusCode: http.StatusCreated,
		Message:    "user registered successfully",
		Data: dto.UserDataResponse{
			Age:      newUserRequsest.Age,
			Email:    newUserRequsest.Email,
			Id:       userId,
			Username: newUserRequsest.Username,
		},
	}

	return &response, nil
}

func (u *userService) Login(payload dto.NewUserLogin) (*dto.LoginResponse, errs.MessageErr) {
	err := helpers.ValidateStruct(payload)

	if err != nil {
		return nil, err
	}

	user, err := u.userRepo.GetUserByEmail(payload.Email)
	if err != nil {
		if err.Status() == http.StatusNotFound {
			return nil, errs.NewBadRequest("invalid email/password")
		}
		return nil, err
	}

	isValidPassword := user.ComparePassword(payload.Password)
	if !isValidPassword {
		return nil, errs.NewBadRequest("invalid email/password")
	}

	token := user.GenerateToken()

	response := dto.LoginResponse{
		Result:     "success",
		StatusCode: http.StatusOK,
		Message:    "successfully logged in",
		Data: dto.TokenResponse{
			Token: token,
		},
	}

	return &response, nil
}

func (u *userService) UpdateUser(userId int, newUpdate dto.NewUserUpdate) (*dto.UserUpdateResponse, errs.MessageErr) {
	err := helpers.ValidateStruct(newUpdate)
	if err != nil {
		return nil, err
	}

	payload := entity.User{
		Email:    newUpdate.Email,
		Username: newUpdate.Username,
	}

	update, err := u.userRepo.UpdateUserById(userId, payload)

	if err != nil {
		return nil, err
	}

	response := dto.UserUpdateResponse{
		Result:     "success",
		StatusCode: http.StatusOK,
		Message:    "successfully updated",
		Data: dto.MakeUserUpdate{
			Id:         update.Id,
			Email:      newUpdate.Email,
			Username:   newUpdate.Username,
			Age:        update.Age,
			Updated_at: update.Updated_at,
		},
	}

	return &response, nil
}

func (u *userService) DeleteUser(userId int) (*dto.DeleteResponse, errs.MessageErr) {

	_, err := u.userRepo.GetUserById(userId)

	if err != nil {
		return nil, err
	}

	err = u.userRepo.DeleteUserById(userId)
	if err != nil {
		return nil, err
	}

	response := dto.DeleteResponse{
		StatusCode: http.StatusOK,
		Message:    "Your account has been succsessfully deleted",
	}

	return &response, nil
}
