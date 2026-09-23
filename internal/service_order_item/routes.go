package service_order_item

import (
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(router *gin.Engine, handler *ServiceOrderItemHandler) {
	router.POST("/service-orders/:id/items", handler.AddItem)
	router.GET("/service-orders/:id/items", handler.ListItems)
	router.PUT("/service-orders/:id/items/:itemId", handler.UpdateItem)
	router.DELETE("/service-orders/:id/items/:itemId", handler.RemoveItem)
}
