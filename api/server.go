package api

import (
	"github.com/gin-gonic/gin"
)

func SetupServer() *gin.Engine {

	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()

	r.GET("/metrics")
	r.GET("/health")
	r.POST("/task", task)

	return r
}
