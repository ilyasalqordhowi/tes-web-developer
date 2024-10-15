package routers

import (
	"soal-test/controllers"

	"github.com/gin-gonic/gin"
)

func FilmRouter(r *gin.RouterGroup) {
	r.GET("/",controllers.GetFilmByGenre)
}