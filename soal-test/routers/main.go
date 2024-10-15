package routers

import (
	"github.com/gin-gonic/gin"
)

func RouterCombine(r *gin.Engine){	
	FilmRouter(r.Group("/genre"))
	HasilRouter(r.Group("/hasil"))
	MessageRouter(r.Group("/pesan"))
	WaktuRouter(r.Group("/waktu"))
	NamaRouter(r.Group("/menyapa"))
}