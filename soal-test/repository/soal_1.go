package repository

import (
	"context"
	"soal-test/lib"
	"soal-test/models"

	"github.com/jackc/pgx/v5"
)


func GetFilmByGenre(id int) ([]models.Film, error) {
	db := lib.DB()
	defer db.Close(context.Background())

	sql := `SELECT fg.id, f.name, f."imgFilm"
			FROM film f
			JOIN "filmByGenre" fg ON fg."filmId" = f.id
			JOIN "genreFilm" c ON c.id = fg."genreFilmId"
			WHERE fg."genreFilmId" = $1;
			`
	rows, err := db.Query(context.Background(), sql, id)
	if err != nil {
		return nil, err
	}

	film, err := pgx.CollectRows(rows, pgx.RowToStructByName[models.Film])
	if err != nil {
		return nil, err
	}

	return film, nil
}