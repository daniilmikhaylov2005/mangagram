package repository

import (
	"github.com/daniilmikhaylov2005/mangagram/models"
)

func InsertUser(user models.User) (models.User, error) {
	db := getConnection()
	defer db.Close()

	tx, err := db.Begin()

	if err != nil {
		return models.User{}, err
	}

	query := `INSERT INTO users (username, email, password) VALUES ($1, $2, $3) RETURNING *`
	row := tx.QueryRow(query, user.Username, user.Email, user.Password)
	var rUser models.User

	if err := row.Scan(&rUser.ID, &rUser.Username, &rUser.Email, &rUser.Password); err != nil {
		tx.Rollback()
		return models.User{}, err
	}

	tx.Commit()
	return rUser, nil
}

func SelectUserByEmail(email string) (models.User, error) {
	db := getConnection()
	defer db.Close()

	query := `SELECT * FROM users WHERE email=$1`
	row := db.QueryRow(query, email)

	var user models.User

	if err := row.Scan(&user.ID, &user.Username, &user.Email, &user.Password); err != nil {
		return models.User{}, err
	}

	return user, nil
}
