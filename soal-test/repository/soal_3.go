package repository

import (
	"context"
	"soal-test/lib"
	"soal-test/models"
)


func SaveMessage(pesan models.Message) error {
	db := lib.DB()
	defer db.Close(context.Background())

	sql := `INSERT INTO pesan (name) VALUES ($1)`
	_, err := db.Exec(context.Background(), sql, pesan.Pesan)
	return err
}
