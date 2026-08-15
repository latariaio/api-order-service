package service_order

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func RegisterRoutes(router *gin.Engine, db *gorm.DB) {
	repo := NewServiceOrderRepository(db)
	service := NewServiceOrderService(repo)
	handler := NewServiceOrderHandler(service)

	router.POST("/service-orders", handler.Create)
	router.GET("/service-orders", handler.GetAll)
	router.GET("/service-orders/:id", handler.GetByID)
	router.PUT("/service-orders/:id", handler.Update)
	router.DELETE("/service-orders/:id", handler.Delete)
}
