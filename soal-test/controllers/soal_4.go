package controllers

import (
	"soal-test/lib"
	"soal-test/repository"

	"github.com/gin-gonic/gin"
)

func GetCurrentTimeHandler(c *gin.Context) {
	repo := repository.TimeRepo{}
	currentTime := repo.GetCurrentTime()

	if currentTime.Tanggal == "" && currentTime.Waktu == "" {
		lib.HandlerNotFound(c, "Time not found")
		return
	}

	lib.HandlerOK(c, "Current time retrieved successfully", currentTime, nil)
}
