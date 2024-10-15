package controllers

import (
	"net/http"
	"soal-test/models"
	"soal-test/repository"

	"github.com/gin-gonic/gin"
)


func SaveMessageHandler(ctx *gin.Context) {
	 pesan := models.Message{}

	
	if err := ctx.ShouldBindJSON(&pesan); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}


	if err := repository.SaveMessage(pesan); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"status": "gagal", "error": "Failed to save message"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"status": "sukses", "pesan": pesan.Pesan})
}
