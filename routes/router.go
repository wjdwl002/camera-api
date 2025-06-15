package routes

import (
	"camera-app/backend/internal/camera_brand"
	"camera-app/backend/internal/health"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func InitRouter(db *gorm.DB) *gin.Engine {
	r := gin.Default()

	healthGroup := r.Group("/health")
	{
		healthGroup.GET("", health.HealthCheck)
	}

	/** Camera brand */
	repo := camera_brand.NewRepository(db)
	service := camera_brand.NewService(repo)
	handler := camera_brand.NewHandler(service)
	handler.RegisterRoutes(r)

	return r
}
