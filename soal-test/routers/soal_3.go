package routers

import (
	"soal-test/controllers"

	"github.com/gin-gonic/gin"
)

func MessageRouter(r *gin.RouterGroup) {
	r.POST("/",controllers.SaveMessageHandler)
}