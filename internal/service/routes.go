package service

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func RegisterRoutes(router *gin.Engine, db *gorm.DB) {
	repo := NewServiceRepository(db)
	service := NewServices(repo)
	handler := NewHandlerService(service)

	router.POST("/services", handler.Create)
	router.GET("/services", handler.GetServices)
	router.GET("/services/:id", handler.GetServiceById)
	router.PATCH("/services/:id", handler.UpdateService)
	router.DELETE("/services/:id", handler.DeleteService)
}
