package main

import (
	"context"
	"net/http"
	"time"

	"go.uber.org/fx"
	"go.uber.org/fx/fxevent"
	"go.uber.org/zap"
	"gorm.io/gorm"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"

	"github.com/latariaio/api-order-service/internal/customer"
	"github.com/latariaio/api-order-service/internal/database"
	"github.com/latariaio/api-order-service/internal/service"
	"github.com/latariaio/api-order-service/internal/service_order"
	"github.com/latariaio/api-order-service/internal/service_order_item"
)

func runMigrations() error {
	return database.RunMigrations()
}

func newRouter() *gin.Engine {
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

	return router
}

func startHTTPServer(lc fx.Lifecycle, router *gin.Engine) {
	srv := &http.Server{
		Addr: ":8080", Handler: router,
	}

	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			go func() {
				_ = srv.ListenAndServe()
			}()
			return nil
		},
		OnStop: func(ctx context.Context) error {
			return srv.Shutdown(ctx)
		},
	})
}

func main() {
	fx.New(
		// Troca o logger padrão (bem verboso) do Fx pelo zap
		fx.WithLogger(func(logger *zap.Logger) fxevent.Logger {
			return &fxevent.ZapLogger{Logger: logger}
		}),

		fx.Provide(
			zap.NewProduction,
			newRouter,
			database.Connect, // assume func() *gorm.DB

			// customer
			customer.NewCustomerRepository,
			customer.NewCustomerService,
			customer.NewCustomerHandler,

			// service
			service.NewServiceRepository,
			service.NewServices,
			service.NewHandlerService,

			// service_order
			service_order.NewServiceOrderRepository,
			service_order.NewServiceOrderService,
			service_order.NewServiceOrderHandler,

			// service_order_item
			service_order_item.NewServiceOrderItemRepository,
			service_order_item.NewServiceOrderItemService,
			service_order_item.NewServiceOrderItemHandler,
		),

		// Ordem de execução dos Invokes = ordem declarada abaixo
		fx.Invoke(
			runMigrations,
			customer.RegisterRoutes,
			service.RegisterRoutes,
			service_order.RegisterRoutes,
			service_order_item.RegisterRoutes,
			startHTTPServer,
		),
	).Run()
}

var _ = gorm.DB{}
