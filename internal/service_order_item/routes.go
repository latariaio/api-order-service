package service_order_item

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func RegisterRoutes(router *gin.Engine, db *gorm.DB) {
	repo := NewServiceOrderItemRepository(db)
	service := NewServiceOrderItemService(repo)
	handler := NewServiceOrderItemHandler(service)

	router.GET("/service_order_items", handler.GetServiceOrderItems)
	router.POST("/service_order_items", handler.CreateServiceOrderItem)
	router.GET("/service_order_items/:id", handler.GetServiceOrderItem)
	router.PUT("/service_order_items/:id", handler.UpdateServiceOrderItem)
	router.DELETE("/service_order_items/:id", handler.DeleteServiceOrderItem)
}
