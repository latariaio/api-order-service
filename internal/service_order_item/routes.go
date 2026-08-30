package service_order_item

import (
	"github.com/gin-gonic/gin"
	"github.com/latariaio/api-order-service/internal/service"
	"github.com/latariaio/api-order-service/internal/service_order"
	"gorm.io/gorm"
)

func RegisterRoutes(router *gin.Engine, db *gorm.DB) {
	itemRepo := NewServiceOrderItemRepository(db)
	orderRepo := service_order.NewServiceOrderRepository(db)
	serviceRepo := service.NewServiceRepository(db)

	svc := NewServiceOrderItemService(itemRepo, orderRepo, serviceRepo)
	handler := NewServiceOrderItemHandler(svc)

	router.POST("/service-orders/:id/items", handler.AddItem)
	router.GET("/service-orders/:id/items", handler.ListItems)
	router.PUT("/service-orders/:id/items/:itemId", handler.UpdateItem)
	router.DELETE("/service-orders/:id/items/:itemId", handler.RemoveItem)
}
