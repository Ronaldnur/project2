package socialmedia_repository

import (
	"project2/entity"
	"project2/pkg/errs"
)

type Repository interface {
	CreateNewSocialMedia(newSocialMedia entity.SocialMedia, userId int) (*entity.SocialMedia, errs.MessageErr)
	GetSocialMedia() (*[]SocialMediaUser, errs.MessageErr)
	UpdateSocialMedia(socialMediaId int, updateSocialMedia entity.SocialMedia) (*entity.SocialMedia, errs.MessageErr)
	GetSocialMediaById(socialMediaId int) (*entity.SocialMedia, errs.MessageErr)
	DeleteSocialMedia(socialMediaId int) errs.MessageErr
}
