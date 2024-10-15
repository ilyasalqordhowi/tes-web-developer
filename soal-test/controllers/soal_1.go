package controllers

import (
	"fmt"
	"soal-test/lib"
	"soal-test/repository"
	"strconv"

	"github.com/gin-gonic/gin"
)
func GetFilmByGenre(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	film, err := repository.GetFilmByGenre(id)
	fmt.Println(film)
	if err != nil {
		lib.HandlerNotFound(c, "data not found")
		return
	}

	lib.HandlerOK(c, "Get film by genre", film, nil)
}