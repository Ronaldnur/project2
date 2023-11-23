package photo_repository

import (
	"project2/entity"
	"project2/pkg/errs"
)

type Repository interface {
	CreateNewPhoto(newPhoto entity.Photo, userId int) (*entity.Photo, errs.MessageErr)
	GetPhotos() (*[]PhotoUser, errs.MessageErr)
	UpdatePhoto(photoId int, newUpdatePhoto entity.Photo) (*entity.Photo, errs.MessageErr)
	GetPhotoById(photoId int) (*entity.Photo, errs.MessageErr)
	DeletePhoto(photoId int) errs.MessageErr
}
