package user_repository

import (
	"project2/entity"
	"project2/pkg/errs"
)

type Repository interface {
	CreateNewUser(user entity.User) (int, errs.MessageErr)
	GetUserByEmail(userEmail string) (*entity.User, errs.MessageErr)
	UpdateUserById(userId int, userUpdate entity.User) (*entity.User, errs.MessageErr)
	GetUserById(userId int) (*entity.User, errs.MessageErr)
	DeleteUserById(userId int) errs.MessageErr
	GetUserByUsername(userUsername string) (*entity.User, errs.MessageErr)
}
