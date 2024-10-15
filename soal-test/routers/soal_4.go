package routers

import (
	"soal-test/controllers"

	"github.com/gin-gonic/gin"
)

func WaktuRouter(r *gin.RouterGroup) {
	r.GET("/",controllers.GetCurrentTimeHandler)
}