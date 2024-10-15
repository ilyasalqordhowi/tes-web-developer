package repository

import (
	"soal-test/models"
	"time"
)


type TimeRepository interface {
	GetCurrentTime() models.TimeModel
}


type TimeRepo struct{}

func (tr *TimeRepo) GetCurrentTime() models.TimeModel {
	return models.TimeModel{
		Tanggal: time.Now().Format("2006-01-02"),
		Waktu:   time.Now().Format("15:04:05"),
	}
}
