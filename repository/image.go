package repository

import (
	"github.com/daniilmikhaylov2005/mangagram/models"
)

func SelectAllImages() ([]models.Image, error) {
	db := getConnection()
	defer db.Close()

	query := `SELECT * FROM images`
	rows, err := db.Query(query)

	if err != nil {
		return []models.Image{}, err
	}

	var images []models.Image

	defer rows.Close()

	for rows.Next() {
		var image models.Image

		if err := rows.Scan(&image.ID, &image.Url, &image.Author); err != nil {
			return []models.Image{}, err
		}

		images = append(images, image)
	}

	return images, nil
}

func SelectImageById(id int) (models.Image, error) {
	db := getConnection()
	defer db.Close()

	query := `SELECT * FROM images WHERE id=$1`
	row := db.QueryRow(query, id)

	var image models.Image

	if err := row.Scan(&image.ID, &image.Url, &image.Author); err != nil {
		return models.Image{}, err
	}

	return image, nil
}

func InsertImage(image models.Image) error {
	db := getConnection()
	defer db.Close()

	query := `INSERT INTO images (url, author) VALUES ($1, $2)`
	_, err := db.Exec(query, image.Url, image.Author)

	if err != nil {
		return err
	}

	return nil
}
