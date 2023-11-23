package socialmedia_pg

import (
	"database/sql"
	"errors"
	"project2/entity"
	"project2/pkg/errs"
	"project2/repository/socialmedia_repository"
)

const (
	GetuserIdQuery = `
SELECT id FROM "user" WHERE id = $1
`

	CreateSocialMediaQuery = `
INSERT INTO "socialmedia"
(
	name,
	social_media_url,
	user_id
	
)
VALUES($1, $2, $3)
RETURNING id,user_id,created_at
`

	GetSocialMediaWithUser = `
SELECT "s"."id", "s"."name", "s"."social_media_url", "s"."user_id", "s"."created_at", "s"."updated_at", "u"."id","u"."username"
FROM "socialmedia" as "s"
LEFT JOIN "user" as "u" ON "s"."user_id" = "u"."id"
`
	GetSocialMediaById = `
SELECT id, name, social_media_url, user_id, created_at, updated_at
FROM "socialmedia"
WHERE id=$1;
`
	UpdateSocialMedia = `
UPDATE "socialmedia"
SET name=$2,
social_media_url=$3
WHERE id=$1
RETURNING id,name,social_media_url,user_id,updated_at
`

	DeleteSocialMedia = `
DELETE FROM "socialmedia"
WHERE "id" = $1;
`
)

type socialmediaPG struct {
	db *sql.DB
}

func NewSocialMediaPG(db *sql.DB) socialmedia_repository.Repository {
	return &socialmediaPG{
		db: db,
	}
}

func (s *socialmediaPG) CreateNewSocialMedia(newSocialMedia entity.SocialMedia, userId int) (*entity.SocialMedia, errs.MessageErr) {
	var userRow int
	var socialMedia entity.SocialMedia

	err := s.db.QueryRow(GetuserIdQuery, userId).Scan(&userRow)

	if err != nil {

		if errors.Is(sql.ErrNoRows, err) {
			return nil, errs.NewNotFoundError("user not found")
		}
		return nil, errs.NewInternalServerError("something went wrong")
	}

	err = s.db.QueryRow(CreateSocialMediaQuery, newSocialMedia.Name, newSocialMedia.Social_media_url, userRow).Scan(&socialMedia.Id, &socialMedia.User_id, &socialMedia.Created_at)

	if err != nil {

		if errors.Is(sql.ErrNoRows, err) {
			return nil, errs.NewNotFoundError("Social Media not found")
		}
		return nil, errs.NewInternalServerError("something went wrong")
	}

	return &socialMedia, nil
}

func (s *socialmediaPG) GetSocialMedia() (*[]socialmedia_repository.SocialMediaUser, errs.MessageErr) {

	rows, err := s.db.Query(GetSocialMediaWithUser)

	if err != nil {

		return nil, errs.NewInternalServerError("something went wrong")
	}

	socialMediaUsers := []socialmedia_repository.SocialMediaUser{}

	for rows.Next() {
		var socialMediaUser socialmedia_repository.SocialMediaUser

		err = rows.Scan(
			&socialMediaUser.Socialmedia.Id, &socialMediaUser.Socialmedia.Name, &socialMediaUser.Socialmedia.Social_media_url, &socialMediaUser.Socialmedia.User_id, &socialMediaUser.Socialmedia.Created_at, &socialMediaUser.Socialmedia.Updated_at,
			&socialMediaUser.User.Id, &socialMediaUser.User.Username,
		)

		if err != nil {
			return nil, errs.NewInternalServerError("something went wrong")
		}

		socialMediaUsers = append(socialMediaUsers, socialMediaUser)
	}
	return &socialMediaUsers, nil
}

func (s *socialmediaPG) UpdateSocialMedia(socialMediaId int, updateSocialMedia entity.SocialMedia) (*entity.SocialMedia, errs.MessageErr) {
	var socialMediaUpdate entity.SocialMedia

	rows := s.db.QueryRow(UpdateSocialMedia, socialMediaId, updateSocialMedia.Name, updateSocialMedia.Social_media_url)

	err := rows.Scan(&socialMediaUpdate.Id, &socialMediaUpdate.Name, &socialMediaUpdate.Social_media_url, &socialMediaUpdate.User_id, &socialMediaUpdate.Updated_at)

	if err != nil {

		if errors.Is(sql.ErrNoRows, err) {
			return nil, errs.NewNotFoundError("Social Media not found")
		}
		return nil, errs.NewInternalServerError("something went wrong")
	}
	return &socialMediaUpdate, nil
}

func (s *socialmediaPG) GetSocialMediaById(socialMediaId int) (*entity.SocialMedia, errs.MessageErr) {
	var getById entity.SocialMedia

	rows := s.db.QueryRow(GetSocialMediaById, socialMediaId)

	err := rows.Scan(&getById.Id, &getById.Name, &getById.Social_media_url, &getById.User_id, &getById.Created_at, &getById.Updated_at)

	if err != nil {
		if errors.Is(sql.ErrNoRows, err) {
			return nil, errs.NewNotFoundError("Social Media not found")
		}
		return nil, errs.NewInternalServerError("something went wrong")
	}

	return &getById, nil
}

func (s *socialmediaPG) DeleteSocialMedia(socialMediaId int) errs.MessageErr {
	_, err := s.db.Exec(DeleteSocialMedia, socialMediaId)

	if err != nil {

		return errs.NewInternalServerError("something went wrong")
	}
	return nil

}
