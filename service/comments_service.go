package service

import (
	"net/http"
	"project2/dto"
	"project2/entity"
	"project2/pkg/errs"
	"project2/pkg/helpers"
	"project2/repository/comments_repository"
)

type commentService struct {
	commentRepo comments_repository.Repository
}

type CommentService interface {
	CreateNewComment(newCommentRequest dto.NewCommentRequest, comment_Id int) (*dto.CommentResponse, errs.MessageErr)
	GetComments(userId int) (*dto.GetCommentsResponse, errs.MessageErr)
	UpdateComment(commentId int, newUpdateRequest dto.MakeCommentUpdate) (*dto.UpdateResponse, errs.MessageErr)
	DeleteComment(commentId int) (*dto.DeleteCommentResponse, errs.MessageErr)
}

func NewCommentService(commentRepo comments_repository.Repository) CommentService {
	return &commentService{
		commentRepo: commentRepo,
	}
}

func (c *commentService) CreateNewComment(newCommentRequest dto.NewCommentRequest, comment_Id int) (*dto.CommentResponse, errs.MessageErr) {
	err := helpers.ValidateStruct(newCommentRequest)

	if err != nil {
		return nil, err
	}

	comment := entity.Comment{
		Message:  newCommentRequest.Message,
		Photo_id: newCommentRequest.Photo_id,
	}

	commentReturn, err := c.commentRepo.CreateComment(comment, comment_Id)

	if err != nil {
		return nil, err
	}

	response := dto.CommentResponse{
		Result:     "Success",
		StatusCode: http.StatusCreated,
		Message:    "Comment Successfully Created",
		Data: dto.MakeDataComment{
			Id:         commentReturn.Id,
			Message:    newCommentRequest.Message,
			Photo_id:   newCommentRequest.Photo_id,
			User_id:    commentReturn.User_id,
			Created_at: commentReturn.Created_at,
		},
	}

	return &response, nil

}

func (c *commentService) GetComments(userId int) (*dto.GetCommentsResponse, errs.MessageErr) {
	comments, err := c.commentRepo.GetComments()

	if err != nil {
		return nil, err
	}

	commentResult := []dto.GetCommentsRequest{}

	for _, eachComment := range *comments {
		comment := dto.GetCommentsRequest{
			Id:         eachComment.Comment.Id,
			Message:    eachComment.Comment.Message,
			Photo_id:   eachComment.Comment.Photo_id,
			User_id:    eachComment.Comment.User_id,
			Updated_at: eachComment.Comment.Updated_at,
			Created_at: eachComment.Comment.Created_at,
			User: dto.GetUserForComment{
				Id:       eachComment.User.Id,
				Email:    eachComment.User.Email,
				Username: eachComment.User.Username,
			},
			Photo: dto.GetPhotoForComment{
				Id:        eachComment.Photo.Id,
				Title:     eachComment.Photo.Title,
				Caption:   eachComment.Photo.Caption,
				Photo_url: eachComment.Photo.Photo_url,
				User_id:   eachComment.Photo.User_id,
			},
		}
		commentResult = append(commentResult, comment)
	}
	response := dto.GetCommentsResponse{
		Result:     "Success",
		StatusCode: http.StatusOK,
		Message:    "Successfully Get All Comment data",
		Data:       commentResult,
	}
	return &response, nil
}

func (c *commentService) UpdateComment(commentId int, newUpdateRequest dto.MakeCommentUpdate) (*dto.UpdateResponse, errs.MessageErr) {
	err := helpers.ValidateStruct(newUpdateRequest)

	if err != nil {
		return nil, err
	}
	updateComment := entity.Comment{
		Message: newUpdateRequest.Message,
	}
	commentUpdate, err := c.commentRepo.UpdateComment(commentId, updateComment)
	if err != nil {
		return nil, err
	}

	response := dto.UpdateResponse{

		Result:     "Success",
		StatusCode: http.StatusOK,
		Message:    "Successfully Update Comment By Id",
		Data: dto.UpdateComment{
			Id:         commentUpdate.Id,
			Photo_id:   commentUpdate.Photo_id,
			User_id:    commentUpdate.User_id,
			Message:    newUpdateRequest.Message,
			Updated_at: commentUpdate.Updated_at,
		},
	}
	return &response, nil

}

func (c *commentService) DeleteComment(commentId int) (*dto.DeleteCommentResponse, errs.MessageErr) {
	err := c.commentRepo.DeleteComment(commentId)
	if err != nil {
		return nil, err
	}

	response := dto.DeleteCommentResponse{
		StatusCode: http.StatusOK,
		Message:    "Your comment has been successfully deleted",
	}

	return &response, nil
}
