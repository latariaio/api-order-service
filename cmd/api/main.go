package main

import (
	"log"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/latariaio/api-order-service/internal/customer"
	"github.com/latariaio/api-order-service/internal/database"
	"github.com/latariaio/api-order-service/internal/service"
	"github.com/latariaio/api-order-service/internal/service_order"
	"github.com/latariaio/api-order-service/internal/service_order_item"
)

func main() {
	if err := database.RunMigrations(); err != nil {
		log.Fatalf("failed to run migrations: %v", err)
	}

	db := database.Connect()

	router := gin.Default()

	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:5173"}, // porta padrão do Vite
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong!",
		})
	})

	customer.RegisterRoutes(router, db)
	service.RegisterRoutes(router, db)
	service_order.RegisterRoutes(router, db)
	service_order_item.RegisterRoutes(router, db)

	router.Run()
}
