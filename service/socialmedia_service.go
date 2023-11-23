package service

import (
	"net/http"
	"project2/dto"
	"project2/entity"
	"project2/pkg/errs"
	"project2/pkg/helpers"
	"project2/repository/socialmedia_repository"
)

type socialMediaService struct {
	socialRepo socialmedia_repository.Repository
}

type SocialMediaService interface {
	MakeSocialMedia(userId int, newSocialMedia dto.NewSocialMediaRequest) (*dto.SocialMediaResponse, errs.MessageErr)
	GetSocialMedia(userId int) (*dto.GetSocialMediaResponse, errs.MessageErr)
	UpdateSocialMedia(socialMediaId int, updateSocialMedia dto.NewSocialMediaRequest) (*dto.UpdateSocialMediaResponse, errs.MessageErr)
	DeleteSocialMedia(socialMediaId int) (*dto.DeleteResponseSocialMedia, errs.MessageErr)
}

func NewSocialMediaService(socialRepo socialmedia_repository.Repository) SocialMediaService {
	return &socialMediaService{
		socialRepo: socialRepo,
	}
}

func (s *socialMediaService) MakeSocialMedia(userId int, newSocialMedia dto.NewSocialMediaRequest) (*dto.SocialMediaResponse, errs.MessageErr) {
	err := helpers.ValidateStruct(newSocialMedia)

	if err != nil {
		return nil, err
	}

	socialMedia := entity.SocialMedia{
		Name:             newSocialMedia.Name,
		Social_media_url: newSocialMedia.Social_media_url,
	}

	newSocial, err := s.socialRepo.CreateNewSocialMedia(socialMedia, userId)
	if err != nil {
		return nil, err
	}

	response := dto.SocialMediaResponse{
		Result:     "success",
		StatusCode: http.StatusCreated,
		Message:    "Social Media successfully maked",
		Data: dto.SocialMediaReturn{
			Id:               newSocial.Id,
			Name:             newSocialMedia.Name,
			Social_media_url: newSocialMedia.Social_media_url,
			User_id:          newSocial.User_id,
			Created_at:       newSocial.Created_at,
		},
	}
	return &response, nil
}
func (s *socialMediaService) GetSocialMedia(userId int) (*dto.GetSocialMediaResponse, errs.MessageErr) {
	socialMedias, err := s.socialRepo.GetSocialMedia()

	if err != nil {
		return nil, err
	}

	socialMediaResult := []dto.GetSocialMedia{}

	for _, eachSocialMedia := range *socialMedias {
		social := dto.GetSocialMedia{
			Id:               eachSocialMedia.Socialmedia.Id,
			Name:             eachSocialMedia.Socialmedia.Name,
			Social_media_url: eachSocialMedia.Socialmedia.Social_media_url,
			User_id:          eachSocialMedia.Socialmedia.User_id,
			Created_at:       eachSocialMedia.Socialmedia.Created_at,
			Updated_at:       eachSocialMedia.Socialmedia.Updated_at,
			User: dto.GetUserForSocialMedia{
				Id:       eachSocialMedia.User.Id,
				Username: eachSocialMedia.User.Username,
			},
		}
		socialMediaResult = append(socialMediaResult, social)
	}

	response := dto.GetSocialMediaResponse{
		Result:     "success",
		StatusCode: http.StatusOK,
		Message:    "Get Social Media successfully",
		Data:       socialMediaResult,
	}
	return &response, nil
}

func (s *socialMediaService) UpdateSocialMedia(socialMediaId int, updateSocialMedia dto.NewSocialMediaRequest) (*dto.UpdateSocialMediaResponse, errs.MessageErr) {
	err := helpers.ValidateStruct(updateSocialMedia)

	if err != nil {
		return nil, err
	}

	updateSocial := entity.SocialMedia{
		Name:             updateSocialMedia.Name,
		Social_media_url: updateSocialMedia.Social_media_url,
	}

	socialMediaUpdate, err := s.socialRepo.UpdateSocialMedia(socialMediaId, updateSocial)

	if err != nil {

		return nil, err
	}

	response := dto.UpdateSocialMediaResponse{
		Result:     "success",
		StatusCode: http.StatusOK,
		Message:    " Social Media successfully updated",
		Data: dto.UpdateSocialMediaReturn{
			Id:               socialMediaUpdate.Id,
			Name:             updateSocial.Name,
			Social_media_url: updateSocial.Social_media_url,
			User_id:          socialMediaUpdate.User_id,
			Updated_at:       socialMediaUpdate.Updated_at,
		},
	}

	return &response, nil
}

func (s *socialMediaService) DeleteSocialMedia(socialMediaId int) (*dto.DeleteResponseSocialMedia, errs.MessageErr) {
	err := s.socialRepo.DeleteSocialMedia(socialMediaId)

	if err != nil {

		return nil, err
	}

	response := dto.DeleteResponseSocialMedia{
		StatusCode: http.StatusOK,
		Message:    " Your social media has been  successfully deleted",
	}

	return &response, nil

}
