package comments_repository

import "project2/entity"

type CommentUserPhoto struct {
	Comment entity.Comment
	User    entity.User
	Photo   entity.Photo
}
