package service

import (
	"errors"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/latariaio/api-order-service/internal/httputil"
)

type HandlerService struct {
	service *Services
}

func NewHandlerService(service *Services) *HandlerService {
	return &HandlerService{service: service}
}

func (h *HandlerService) Create(ctx *gin.Context) {
	var request CreateServiceRequest
	if err := ctx.ShouldBindJSON(&request); err != nil {
		if fieldErrors := httputil.FormatValidationError(err); fieldErrors != nil {
			ctx.JSON(400, gin.H{"errors": fieldErrors})
			return
		}
		ctx.JSON(400, gin.H{"error": "invalid request body"})
		return
	}

	svc := request.ToModel()

	if err := h.service.Create(&svc); err != nil {
		switch {
		case errors.Is(err, ErrInvalidPrice):
			ctx.JSON(400, gin.H{"error": err.Error()})
		case errors.Is(err, ErrServiceNameAlreadyExists):
			ctx.JSON(409, gin.H{"error": err.Error()})
		default:
			log.Printf("unexpected error creating service: %v", err)
			ctx.JSON(500, gin.H{"error": "internal server error"})
		}
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"data": ToServiceResponse(svc)})
}

func (h *HandlerService) GetServices(ctx *gin.Context) {
	services, err := h.service.GetAll()
	if err != nil {
		log.Printf("unexpected error listing services: %v", err)
		ctx.JSON(500, gin.H{"error": "internal server error"})
		return
	}
	ctx.JSON(200, gin.H{"data": ToServiceResponseList(services)})
}

func (h *HandlerService) GetServiceById(ctx *gin.Context) {
	id := ctx.Param("id")
	svc, err := h.service.GetByID(id)
	if err != nil {
		switch {
		case errors.Is(err, ErrServiceNotFound):
			ctx.JSON(404, gin.H{"error": err.Error()})
		default:
			log.Printf("unexpected error getting service: %v", err)
			ctx.JSON(500, gin.H{"error": "internal server error"})
		}
		return
	}
	ctx.JSON(200, gin.H{"data": ToServiceResponse(*svc)})
}

func (h *HandlerService) UpdateService(ctx *gin.Context) {
	id := ctx.Param("id")

	var request UpdateServiceRequest
	if err := ctx.ShouldBindJSON(&request); err != nil {
		if fieldErrors := httputil.FormatValidationError(err); fieldErrors != nil {
			ctx.JSON(400, gin.H{"errors": fieldErrors})
			return
		}
		ctx.JSON(400, gin.H{"error": "invalid request body"})
		return
	}

	svc, err := h.service.Update(id, request)
	if err != nil {
		switch {
		case errors.Is(err, ErrServiceNotFound):
			ctx.JSON(404, gin.H{"error": err.Error()})
		case errors.Is(err, ErrInvalidPrice):
			ctx.JSON(400, gin.H{"error": err.Error()})
		case errors.Is(err, ErrServiceNameAlreadyExists):
			ctx.JSON(409, gin.H{"error": err.Error()})
		default:
			log.Printf("unexpected error updating service: %v", err)
			ctx.JSON(500, gin.H{"error": "internal server error"})
		}
		return
	}

	ctx.JSON(200, gin.H{"data": ToServiceResponse(*svc)})
}

func (h *HandlerService) DeleteService(ctx *gin.Context) {
	id := ctx.Param("id")
	if err := h.service.Delete(id); err != nil {
		switch {
		case errors.Is(err, ErrServiceNotFound):
			ctx.JSON(404, gin.H{"error": err.Error()})
		case errors.Is(err, ErrServiceInUse):
			ctx.JSON(409, gin.H{"error": err.Error()})
		default:
			log.Printf("unexpected error deleting service: %v", err)
			ctx.JSON(500, gin.H{"error": "internal server error"})
		}
		return
	}
	ctx.Status(204)
}
