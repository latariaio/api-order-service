package service_order

import (
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(router *gin.Engine, handler *ServiceOrderHandler) {

	router.POST("/service-orders", handler.Create)
	router.GET("/service-orders", handler.GetAll)
	router.GET("/service-orders/:id", handler.GetByID)
	router.PUT("/service-orders/:id", handler.Update)
	router.PATCH("/service-orders/:id/status", handler.UpdateStatus)
	router.DELETE("/service-orders/:id", handler.Delete)
}
