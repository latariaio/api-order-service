package service

import "github.com/gin-gonic/gin"

type HandlerService struct {
	service *Services
}

func NewHandlerService(service *Services) *HandlerService {
	return &HandlerService{
		service: service,
	}
}

type CreateServiceRequest struct {
	Name        string  `json:"name" binding:"required"`
	Description string  `json:"description"`
	Price       float64 `json:"price" binding:"required"`
}

func (h *HandlerService) Create(ctx *gin.Context) {
	var request CreateServiceRequest
	if err := ctx.ShouldBindJSON(&request); err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}
	service := Service{
		Name:        request.Name,
		Description: request.Description,
		Price:       request.Price,
	}
	if err := h.service.Create(&service); err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(200, service)
}

func (h *HandlerService) GetServices(ctx *gin.Context) {
	services, err := h.service.GetAll()
	if err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(200, gin.H{"data": services})
}

func (h *HandlerService) GetServiceById(ctx *gin.Context) {
	id := ctx.Param("id")
	service, err := h.service.GetByID(id)
	if err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(200, service)
}

func (h *HandlerService) UpdateService(ctx *gin.Context) {
	id := ctx.Param("id")
	var request CreateServiceRequest
	if err := ctx.ShouldBindJSON(&request); err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}
	serviceUpdate := Service{
		Name:        request.Name,
		Description: request.Description,
		Price:       request.Price,
	}
	if err := h.service.Update(id, &serviceUpdate); err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(200, serviceUpdate)
}

func (h *HandlerService) DeleteService(ctx *gin.Context) {
	id := ctx.Param("id")
	if err := h.service.Delete(id); err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(200, gin.H{"message": "service deleted"})
}
