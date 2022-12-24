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
