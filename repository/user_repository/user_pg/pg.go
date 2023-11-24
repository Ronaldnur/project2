package user_pg

import (
	"database/sql"
	"errors"
	"project2/entity"
	"project2/pkg/errs"
	"project2/repository/user_repository"
)

const (
	createNewUser = `
		INSERT INTO "user"
		(
			username,
			email,
			password,
			age
		)
		VALUES ($1, $2, $3, $4)
		RETURNING id
	`

	getUserByEmailQuery = `
    SELECT id, username, email, password, age, created_at, updated_at
    FROM "user"
    WHERE email = $1
`

	updateUserByIdQuery = `
    UPDATE "user"
    SET email = $2, username = $3
    WHERE id = $1
	RETURNING id,username,age,updated_at
`

	retrieveUserById = `
		SELECT id, email, password from "user"
		WHERE id = $1;
	`

	deleteUserByIdQuery = `
	    DELETE FROM "user"
	    WHERE id = $1

`

	getUserByUsernameQuery = `
SELECT id, username, email, password, age, created_at, updated_at
FROM "user"
WHERE username = $1
`
)

type userPG struct {
	db *sql.DB
}

func NewUserPG(db *sql.DB) user_repository.Repository {
	return &userPG{
		db: db,
	}
}

func (u *userPG) CreateNewUser(user entity.User) (int, errs.MessageErr) {
	var userId int
	err := u.db.QueryRow(createNewUser, user.Username, user.Email, user.Password, user.Age).Scan(&userId)
	if err != nil {

		return 0, errs.NewInternalServerError("Failed to create new user")
	}
	return userId, nil
}

func (u *userPG) GetUserByEmail(userEmail string) (*entity.User, errs.MessageErr) {
	var user entity.User

	rows := u.db.QueryRow(getUserByEmailQuery, userEmail)

	err := rows.Scan(&user.Id, &user.Username, &user.Email, &user.Password, &user.Age, &user.Created_at, &user.Updated_at)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errs.NewNotFoundError("user not found")
		}

		return nil, errs.NewInternalServerError("Something went wrong")
	}
	return &user, nil
}
func (u *userPG) GetUserById(userId int) (*entity.User, errs.MessageErr) {
	var user entity.User

	row := u.db.QueryRow(retrieveUserById, userId)

	err := row.Scan(&user.Id, &user.Email, &user.Password)

	if err != nil {
		if errors.Is(sql.ErrNoRows, err) {
			return nil, errs.NewNotFoundError("user not found")
		}
		return nil, errs.NewInternalServerError("something went wrong")
	}

	return &user, nil
}
func (u *userPG) UpdateUserById(userId int, userUpdate entity.User) (*entity.User, errs.MessageErr) {
	var user entity.User
	err := u.db.QueryRow(updateUserByIdQuery, userId, userUpdate.Email, userUpdate.Username).Scan(&user.Id, &user.Username, &user.Age, &user.Updated_at)

	if err != nil {
		if errors.Is(sql.ErrNoRows, err) {
			return nil, errs.NewNotFoundError("user not found")
		}
		return nil, errs.NewInternalServerError("something went wrong")
	}
	return &user, nil
}

func (u *userPG) DeleteUserById(userId int) errs.MessageErr {
	_, err := u.db.Exec(deleteUserByIdQuery, userId)
	if err != nil {
		return errs.NewInternalServerError("something went wrong")
	}
	return nil
}

func (u *userPG) GetUserByUsername(userUsername string) (*entity.User, errs.MessageErr) {
	var user entity.User

	rows := u.db.QueryRow(getUserByUsernameQuery, userUsername)

	err := rows.Scan(&user.Id, &user.Username, &user.Email, &user.Password, &user.Age, &user.Created_at, &user.Updated_at)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errs.NewNotFoundError("username not found")
		}
		return nil, errs.NewInternalServerError("Something went wrong")
	}
	return &user, nil
}
