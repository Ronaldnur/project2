package comments_repository

import (
	"project2/entity"
	"project2/pkg/errs"
)

type Repository interface {
	CreateComment(newComment entity.Comment, userId int) (*entity.Comment, errs.MessageErr)
	GetComments() (*[]CommentUserPhoto, errs.MessageErr)
	UpdateComment(commentId int, newUpdateComment entity.Comment) (*entity.Comment, errs.MessageErr)
	GetCommentById(commentId int) (*entity.Comment, errs.MessageErr)
	DeleteComment(commentId int) errs.MessageErr
}
