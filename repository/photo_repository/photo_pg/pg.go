package photo_pg

import (
	"database/sql"
	"errors"
	"project2/entity"
	"project2/pkg/errs"
	"project2/repository/photo_repository"
)

const (
	CreatePhotoQuery = `
INSERT INTO "photo"
(
	title,
	caption,
	photo_url,
	user_id
	
)
VALUES($1, $2, $3, $4)
RETURNING id,user_id,created_at
`

	GetuserIdQuery = `
SELECT id FROM "user" WHERE id = $1
`
	GetPhotoWithUser = `
SELECT "p"."id", "p"."title", "p"."caption", "p"."photo_url", "p"."user_id", "p"."created_at", "p"."updated_at", "u"."email", "u"."username"
FROM "photo" as "p"
INNER JOIN "user" as "u" ON "p"."user_id" = "u"."id"
`

	UpdatePhoto = `
UPDATE "photo"
SET title=$2,
caption=$3,
photo_url=$4
WHERE id=$1
RETURNING id,title,caption,photo_url,user_id,updated_at
`

	GetPhotoById = `
SELECT id, title, caption, photo_url, user_id, created_at, updated_at
FROM "photo"
WHERE id = $1;
`

	DeletePhoto = `
	DELETE FROM "photo"
	WHERE "id" = $1;
`
)

type photoPG struct {
	db *sql.DB
}

func NewPhotoPG(db *sql.DB) photo_repository.Repository {
	return &photoPG{
		db: db,
	}
}

func (p *photoPG) CreateNewPhoto(newPhoto entity.Photo, userId int) (*entity.Photo, errs.MessageErr) {
	var photo entity.Photo
	var userRow int

	err := p.db.QueryRow(GetuserIdQuery, userId).Scan(&userRow)

	if err != nil {
		if errors.Is(sql.ErrNoRows, err) {
			return nil, errs.NewNotFoundError("user not found")
		}
		return nil, errs.NewInternalServerError("something went wrong")
	}
	err = p.db.QueryRow(CreatePhotoQuery, newPhoto.Title, newPhoto.Caption, newPhoto.Photo_url, userRow).Scan(&photo.Id, &photo.User_id, &photo.Created_at)

	if err != nil {
		if errors.Is(sql.ErrNoRows, err) {
			return nil, errs.NewNotFoundError("photo not found")
		}
		return nil, errs.NewInternalServerError("something went wrong")
	}
	return &photo, nil
}

func (p *photoPG) GetPhotos() (*[]photo_repository.PhotoUser, errs.MessageErr) {
	rows, err := p.db.Query(GetPhotoWithUser)

	if err != nil {

		return nil, errs.NewInternalServerError("something went wrong")
	}

	photoUsers := []photo_repository.PhotoUser{}
	for rows.Next() {
		var photoUser photo_repository.PhotoUser

		err = rows.Scan(
			&photoUser.Photo.Id, &photoUser.Photo.Title, &photoUser.Photo.Caption, &photoUser.Photo.Photo_url, &photoUser.Photo.User_id, &photoUser.Photo.Created_at, &photoUser.Photo.Updated_at,
			&photoUser.User.Email, &photoUser.User.Username,
		)

		if err != nil {
			return nil, errs.NewInternalServerError("something went wrong")
		}

		photoUsers = append(photoUsers, photoUser)
	}

	return &photoUsers, nil
}

func (p *photoPG) UpdatePhoto(photoId int, newUpdatePhoto entity.Photo) (*entity.Photo, errs.MessageErr) {
	var photoUpdate entity.Photo

	rows := p.db.QueryRow(UpdatePhoto, photoId, newUpdatePhoto.Title, newUpdatePhoto.Caption, newUpdatePhoto.Photo_url)

	err := rows.Scan(&photoUpdate.Id, &photoUpdate.Title, &photoUpdate.Caption, &photoUpdate.Photo_url, &photoUpdate.User_id, &photoUpdate.Updated_at)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errs.NewNotFoundError("photo not found")
		}

		return nil, errs.NewInternalServerError("Something went wrong")
	}
	return &photoUpdate, nil
}

func (p *photoPG) GetPhotoById(photoId int) (*entity.Photo, errs.MessageErr) {
	var getPhoto entity.Photo

	rows := p.db.QueryRow(GetPhotoById, photoId)

	err := rows.Scan(&getPhoto.Id, &getPhoto.Title, &getPhoto.Caption, &getPhoto.Photo_url, &getPhoto.User_id, &getPhoto.Created_at, &getPhoto.Updated_at)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errs.NewNotFoundError("photo not found")
		}

		return nil, errs.NewInternalServerError("Something went wrong")
	}
	return &getPhoto, nil
}

func (p *photoPG) DeletePhoto(photoId int) errs.MessageErr {
	_, err := p.db.Exec(DeletePhoto, photoId)
	if err != nil {
		return errs.NewInternalServerError("something went wrong")
	}
	return nil
}
