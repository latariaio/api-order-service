package customer

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func RegisterRoutes(router *gin.Engine, db *gorm.DB) {
	repo := NewCustomerRepository(db)
	service := NewCustomerService(repo)
	handler := NewCustomerHandler(service)

	router.POST("/customers", handler.CreateCustomer)
	router.GET("/customers", handler.GetCustomers)
	router.GET("/customers/:id", handler.GetCustomer)
	router.PATCH("/customers/:id", handler.UpdateCustomer)
	router.DELETE("/customers/:id", handler.DeleteCustomer)
}
