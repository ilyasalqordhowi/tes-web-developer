package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)


func NamaHandler(c *gin.Context) {
	nama := c.Query("nama") 
	if nama == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Nama tidak boleh kosong"})
		return
	}


	message := "Halo, " + nama + "! Selamat datang!"
	c.JSON(http.StatusOK, gin.H{"message": message})
}
