package customer

import (
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(router *gin.Engine, handler *CustomerHandler) {
	router.POST("/customers", handler.CreateCustomer)
	router.GET("/customers", handler.GetCustomers)
	router.GET("/customers/:id", handler.GetCustomer)
	router.PATCH("/customers/:id", handler.UpdateCustomer)
	router.DELETE("/customers/:id", handler.DeleteCustomer)
}
