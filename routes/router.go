package routes

import (
	"camera-app/backend/handlers"

	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	r := gin.Default()

	r.GET("/health", handlers.HealthCheck)

	return r
}
