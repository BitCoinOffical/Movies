package storage

import (
	"database/sql"
	"errors"

	"main.go/internal/models"
)

func (d *DataBase) AddMovies(models *models.Movies) error {
	insert := `INSERT INTO movies(title, genre, year, description) VALUES (?, ?, ?, ?)`
	res, err := d.DB.Exec(insert, models.Title, models.Genre, models.Year, models.Description)
	if err != nil {
		return err
	}
	id, err := res.LastInsertId()
	if err != nil {
		return err
	}
	models.ID = int(id)
	return nil
}

func (d *DataBase) GetMovies() ([]models.Movies, error) {
	selectall := `SELECT id, title, genre, year, description FROM movies ORDER BY id DESC`
	rows, err := d.DB.Query(selectall)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var movies []models.Movies
	for rows.Next() {
		var movie models.Movies
		err := rows.Scan(
			&movie.ID,
			&movie.Title,
			&movie.Genre,
			&movie.Year,
			&movie.Description,
		)
		if err != nil {
			return nil, err
		}
		movies = append(movies, movie)
	}
	return movies, nil
}

func (d *DataBase) GetMoviesByID(id int) (*models.Movies, error) {
	selectall := `SELECT id, title, genre, year, description FROM movies WHERE id = ?`
	row := d.DB.QueryRow(selectall, id)
	var movie models.Movies
	err := row.Scan(
		&movie.ID,
		&movie.Title,
		&movie.Genre,
		&movie.Year,
		&movie.Description,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.New("movie not found")
		}
		return nil, err
	}

	return &movie, nil
}

func (d *DataBase) DeleteMovies(id int) (bool, error) {
	delete := `DELETE FROM movies WHERE id = ?`
	res, err := d.DB.Exec(delete, id)
	if err != nil {
		return false, err
	}
	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return false, err
	}
	return rowsAffected > 0, nil
}

func (d *DataBase) UpdateMoviesByID(id int, updates struct {
	Title       string `json:"title"`
	Genre       string `json:"genre"`
	Year        int    `json:"year"`
	Description string `json:"description"`
}) error {
	query := `UPDATE movies SET 
        title = ?,
        genre = ?,
        year = ?,
        description = ?
    WHERE id = ?`

	_, err := d.DB.Exec(query,
		updates.Title,
		updates.Genre,
		updates.Year,
		updates.Description,
		id,
	)
	return err
}
