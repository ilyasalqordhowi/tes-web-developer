package repository

import (
	"context"
	"soal-test/lib"
	"soal-test/models"
)

func Nilai(luas float64, nilai models.Rectangle) error {
	db := lib.DB()
	defer db.Close(context.Background())

	
	sql := `INSERT INTO nilai (panjang, lebar, luas) VALUES ($1, $2, $3)`
	_, err := db.Exec(context.Background(), sql, nilai.Panjang, nilai.Lebar, luas)
	if err != nil {
		return err 
	}

	return nil
}
