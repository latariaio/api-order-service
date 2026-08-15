package main

import (
	"github.com/gin-gonic/gin"
	"github.com/latariaio/api-order-service/internal/customer"
	"github.com/latariaio/api-order-service/internal/database"
	"github.com/latariaio/api-order-service/internal/service"
	"github.com/latariaio/api-order-service/internal/service_order"
	"github.com/latariaio/api-order-service/internal/service_order_item"
)

func main() {
	db := database.Connect()
	db.AutoMigrate(&customer.Customer{}, &service.Service{})

	router := gin.Default()

	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	customer.RegisterRoutes(router, db)
	service.RegisterRoutes(router, db)
	service_order.RegisterRoutes(router, db)
	service_order_item.RegisterRoutes(router, db)

	router.Run()
}
