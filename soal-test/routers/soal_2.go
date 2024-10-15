package routers

import (
	"soal-test/controllers"

	"github.com/gin-gonic/gin"
)

func HasilRouter(r *gin.RouterGroup) {
	r.POST("/",controllers.HitungLuas)
}