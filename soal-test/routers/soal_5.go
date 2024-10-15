package routers

import (
	"soal-test/controllers"

	"github.com/gin-gonic/gin"
)

func NamaRouter(r *gin.RouterGroup) {
	r.GET("/nama", controllers.NamaHandler)
}