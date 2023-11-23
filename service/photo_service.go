package service

import (
	"net/http"
	"project2/dto"
	"project2/entity"
	"project2/pkg/errs"
	"project2/pkg/helpers"
	"project2/repository/photo_repository"
)

type photoService struct {
	photoRepo photo_repository.Repository
}

type PhotoService interface {
	PostPhoto(userId int, newPhotoRequest dto.NewPhotoRequest) (*dto.NewPhotoResponse, errs.MessageErr)
	GetPhotoUsers(userId int) (*dto.GetPhotoResponse, errs.MessageErr)
	PhotoUpdate(photoId int, newUpdateRequest dto.NewPhotoRequest) (*dto.NewUpdateResponse, errs.MessageErr)
	PhotoDelete(photoId int) (*dto.DeletePhotoResponse, errs.MessageErr)
}

func NewPhotoService(photoRepo photo_repository.Repository) PhotoService {
	return &photoService{
		photoRepo: photoRepo,
	}
}

func (p *photoService) PostPhoto(userId int, newPhotoRequest dto.NewPhotoRequest) (*dto.NewPhotoResponse, errs.MessageErr) {
	err := helpers.ValidateStruct(newPhotoRequest)

	if err != nil {
		return nil, err
	}

	photo := entity.Photo{
		Title:     newPhotoRequest.Title,
		Caption:   newPhotoRequest.Caption,
		Photo_url: newPhotoRequest.Photo_url,
	}

	newPhoto, err := p.photoRepo.CreateNewPhoto(photo, userId)

	if err != nil {
		return nil, err
	}

	response := dto.NewPhotoResponse{
		Result:     "success",
		StatusCode: http.StatusCreated,
		Message:    "photo successfully posted",
		Data: dto.MakeDataPhoto{
			Id:         newPhoto.Id,
			Title:      newPhotoRequest.Title,
			Caption:    newPhotoRequest.Caption,
			Photo_url:  newPhotoRequest.Photo_url,
			User_id:    newPhoto.User_id,
			Created_at: newPhoto.Created_at,
		},
	}

	return &response, nil
}

func (p *photoService) GetPhotoUsers(userId int) (*dto.GetPhotoResponse, errs.MessageErr) {
	photos, err := p.photoRepo.GetPhotos()

	if err != nil {
		return nil, err
	}
	photoResult := []dto.NewGetPhotoRequest{}

	for _, eachPhoto := range *photos {
		photo := dto.NewGetPhotoRequest{
			Id:         eachPhoto.Photo.Id,
			Title:      eachPhoto.Photo.Title,
			Caption:    eachPhoto.Photo.Caption,
			Photo_url:  eachPhoto.Photo.Photo_url,
			User_id:    eachPhoto.Photo.User_id,
			Created_at: eachPhoto.Photo.Created_at,
			Updated_at: eachPhoto.Photo.Updated_at,
			User: dto.GetUser{
				Email:    eachPhoto.User.Email,
				Username: eachPhoto.User.Username,
			},
		}
		photoResult = append(photoResult, photo)
	}

	response := dto.GetPhotoResponse{
		Result:     "success",
		StatusCode: http.StatusOK,
		Message:    "get photo data successfully",
		Data:       photoResult,
	}
	return &response, nil
}

func (p *photoService) PhotoUpdate(photoId int, newUpdateRequest dto.NewPhotoRequest) (*dto.NewUpdateResponse, errs.MessageErr) {
	err := helpers.ValidateStruct(newUpdateRequest)
	if err != nil {
		return nil, err
	}
	updatePhoto := entity.Photo{
		Title:     newUpdateRequest.Title,
		Caption:   newUpdateRequest.Caption,
		Photo_url: newUpdateRequest.Photo_url,
	}

	photoUdpate, err := p.photoRepo.UpdatePhoto(photoId, updatePhoto)

	if err != nil {
		return nil, err
	}

	response := dto.NewUpdateResponse{
		Result:     "success",
		StatusCode: http.StatusOK,
		Message:    "update photo data successfully",
		Data: dto.MakeUpdatePhoto{
			Id:         photoUdpate.Id,
			Title:      newUpdateRequest.Title,
			Caption:    newUpdateRequest.Caption,
			Photo_url:  newUpdateRequest.Photo_url,
			User_id:    photoUdpate.User_id,
			Updated_at: photoUdpate.Updated_at,
		},
	}

	return &response, nil
}

func (p *photoService) PhotoDelete(photoId int) (*dto.DeletePhotoResponse, errs.MessageErr) {
	err := p.photoRepo.DeletePhoto(photoId)

	if err != nil {
		return nil, err
	}
	response := dto.DeletePhotoResponse{
		StatusCode: http.StatusOK,
		Message:    "Your photo has been successfully deleted",
	}

	return &response, nil
}
