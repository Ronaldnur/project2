package photo_repository

import "project2/entity"

type PhotoUser struct {
	Photo entity.Photo
	User  entity.User
}
