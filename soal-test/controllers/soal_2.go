package controllers

import (
	"net/http"
	"soal-test/models"
	"soal-test/repository"

	"github.com/gin-gonic/gin"
)


func HitungLuas(ctx *gin.Context) {
	var nilai models.Rectangle 

	
	if err := ctx.ShouldBindJSON(&nilai); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}


	luas := nilai.Panjang * nilai.Lebar

	if err := repository.Nilai(luas, nilai); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to insert data into the database"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"panjang": nilai.Panjang,
		"lebar":   nilai.Lebar,
		"luas":    luas,
	})
}
