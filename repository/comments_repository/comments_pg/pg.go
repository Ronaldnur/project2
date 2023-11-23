package comments_pg

import (
	"database/sql"
	"errors"
	"project2/entity"
	"project2/pkg/errs"
	"project2/repository/comments_repository"
)

const (
	GetuserIdQuery = `
	SELECT id FROM "user" WHERE id = $1
	`
	CheckPhotoExistenceQuery = `
	SELECT id FROM "photo" WHERE id = $1
	`
	CreateCommentQuery = `
	INSERT INTO "comment"
	(
		message,
		photo_id,
		user_id
	)
	VALUES($1, $2, $3)
	RETURNING id,user_id,created_at
	`

	GetCommentsPhotoUser = `
	SELECT "c"."id","c"."message","c"."photo_id","c"."user_id","c"."updated_at","c"."created_at","u"."id","u"."email","u"."username","p"."id", "p"."title", "p"."caption", "p"."photo_url", "p"."user_id"
	FROM "comment" as "c"
	LEFT JOIN
	"user" AS "u" ON "c"."user_id" = "u"."id"
	LEFT JOIN
	"photo" AS "p" ON "c"."photo_id" = "p"."id"
	`

	GetCommentById = `
SELECT id,user_id,photo_id,message,created_at,updated_at
FROM "comment"
WHERE id = $1
`

	UpdateComment = `
UPDATE "comment"
SET message=$2
WHERE id=$1
RETURNING id,photo_id,user_id,created_at
`

	DeleteComment = `
DELETE FROM "comment"
WHERE "id" = $1;
`
)

type commentPG struct {
	db *sql.DB
}

func NewCommentPG(db *sql.DB) comments_repository.Repository {
	return &commentPG{
		db: db,
	}
}
func (c *commentPG) CreateComment(newComment entity.Comment, userId int) (*entity.Comment, errs.MessageErr) {
	var userRow, photoRow int
	var comment entity.Comment

	err := c.db.QueryRow(GetuserIdQuery, userId).Scan(&userRow)

	if err != nil {
		if errors.Is(sql.ErrNoRows, err) {
			return nil, errs.NewNotFoundError("user not found")
		}
		return nil, errs.NewInternalServerError("something went wrong")
	}

	err = c.db.QueryRow(CheckPhotoExistenceQuery, newComment.Photo_id).Scan(&photoRow)
	if err != nil {
		if errors.Is(sql.ErrNoRows, err) {
			return nil, errs.NewNotFoundError("photo not found")
		}
		return nil, errs.NewInternalServerError("something went wrong")
	}

	err = c.db.QueryRow(CreateCommentQuery, newComment.Message, newComment.Photo_id, userRow).Scan(&comment.Id, &comment.User_id, &comment.Created_at)

	if err != nil {
		if errors.Is(sql.ErrNoRows, err) {
			return nil, errs.NewNotFoundError("comment not found")
		}
		return nil, errs.NewInternalServerError("something went wrong")
	}
	return &comment, nil
}

func (c *commentPG) GetComments() (*[]comments_repository.CommentUserPhoto, errs.MessageErr) {
	rows, err := c.db.Query(GetCommentsPhotoUser)

	if err != nil {
		return nil, errs.NewInternalServerError("something went wrong")
	}

	commentPhotoUsers := []comments_repository.CommentUserPhoto{}
	for rows.Next() {
		var commentPhotoUser comments_repository.CommentUserPhoto

		err = rows.Scan(
			&commentPhotoUser.Comment.Id, &commentPhotoUser.Comment.Message, &commentPhotoUser.Comment.Photo_id, &commentPhotoUser.Comment.User_id, &commentPhotoUser.Comment.Created_at, &commentPhotoUser.Comment.Updated_at,
			&commentPhotoUser.User.Id, &commentPhotoUser.User.Email, &commentPhotoUser.User.Username,
			&commentPhotoUser.Photo.Id, &commentPhotoUser.Photo.Title, &commentPhotoUser.Photo.Caption, &commentPhotoUser.Photo.Photo_url, &commentPhotoUser.Photo.User_id,
		)
		if err != nil {
			return nil, errs.NewInternalServerError("something went wrong")
		}

		commentPhotoUsers = append(commentPhotoUsers, commentPhotoUser)
	}
	return &commentPhotoUsers, nil
}

func (c *commentPG) UpdateComment(commentId int, newUpdateComment entity.Comment) (*entity.Comment, errs.MessageErr) {
	var commentUpdate entity.Comment

	rows := c.db.QueryRow(UpdateComment, commentId, newUpdateComment.Message)
	err := rows.Scan(&commentUpdate.Id, &commentUpdate.Photo_id, &commentUpdate.User_id, &commentUpdate.Updated_at)

	if err != nil {

		if errors.Is(sql.ErrNoRows, err) {
			return nil, errs.NewNotFoundError("comment not found")
		}
		return nil, errs.NewInternalServerError("something went wrong")
	}
	return &commentUpdate, nil
}
func (c *commentPG) GetCommentById(commentId int) (*entity.Comment, errs.MessageErr) {
	var getComment entity.Comment

	rows := c.db.QueryRow(GetCommentById, commentId)

	err := rows.Scan(&getComment.Id, &getComment.User_id, &getComment.Photo_id, &getComment.Message, &getComment.Created_at, &getComment.Updated_at)

	if err != nil {

		if errors.Is(sql.ErrNoRows, err) {
			return nil, errs.NewNotFoundError("comment not found")
		}
		return nil, errs.NewInternalServerError("something went wrong")
	}
	return &getComment, nil
}

func (c *commentPG) DeleteComment(commentId int) errs.MessageErr {

	_, err := c.db.Exec(DeleteComment, commentId)
	if err != nil {
		return errs.NewInternalServerError("something went wrong")
	}
	return nil
}
