package repositories

import (
	"errors"
	"homework-caching-and-redis/constants"
	"homework-caching-and-redis/models"
	"homework-caching-and-redis/utils"
)

func CreateANewUser(user *models.User) error {
	insertQuery := `
	INSERT INTO users(name, email, password, createdAt)
	VALUES (?, ?, ?, ?)`

	stmt, err := DbContext.Prepare(insertQuery)
	if err != nil {
		return err
	}
	defer stmt.Close()
	hashedPasword, err := utils.HashPassword(user.Password)

	if err != nil {
		return err
	}
	result, err := stmt.Exec(
		user.Name,
		user.Email,
		hashedPasword,
		user.CreatedAt.Format(constants.FormatDateTime))

	if err != nil {
		return err
	}

	id, err := result.LastInsertId()
	user.Id = id
	return err
}

func Login(u *models.User) error {
	query := `SELECT id, password FROM users WHERE email = ?`

	row := DbContext.QueryRow(query, u.Email)
	var hashedPasword string
	err := row.Scan(&u.Id, &hashedPasword)

	if err != nil {
		return errors.New("email or password is invalid")
	}

	validPassword := utils.CheckPasswordHash(u.Password, hashedPasword)
	if !validPassword {
		return errors.New("email or password is invalid")
	}

	return nil
}
