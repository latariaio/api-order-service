package service

import (
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(router *gin.Engine, handler *HandlerService) {
	router.POST("/services", handler.Create)
	router.GET("/services", handler.GetServices)
	router.GET("/services/:id", handler.GetServiceById)
	router.PATCH("/services/:id", handler.UpdateService)
	router.DELETE("/services/:id", handler.DeleteService)
}
