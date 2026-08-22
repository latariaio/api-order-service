package customer

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type CustomerHandler struct {
	service *CustomerService
}

func NewCustomerHandler(service *CustomerService) *CustomerHandler {
	return &CustomerHandler{
		service: service,
	}
}

func (h *CustomerHandler) CreateCustomer(ctx *gin.Context) {
	var request CustomerResponse
	if err := ctx.ShouldBindJSON(&request); err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}

	customer := ToModel(request)

	if err := h.service.Create(&customer); err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, ToCustomerResponse(customer))
}

func (h *CustomerHandler) GetCustomer(ctx *gin.Context) {
	id := ctx.Param("id")
	log.Default().Println("getting customer", id)
	customer, err := h.service.FindById(id)
	if err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(200, customer)
}

func (h *CustomerHandler) GetCustomers(ctx *gin.Context) {
	customers, err := h.service.FindAll()
	if err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(200, ToCustomerResponseList(customers))
}

func (h *CustomerHandler) UpdateCustomer(ctx *gin.Context) {
	id := ctx.Param("id")
	var request CustomerResponse
	if err := ctx.ShouldBindJSON(&request); err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}
	customer := ToModel(request)

	if err := h.service.Update(id, &customer); err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(200, ToCustomerResponse(customer))
}

func (h *CustomerHandler) DeleteCustomer(ctx *gin.Context) {
	id := ctx.Param("id")
	if err := h.service.Delete(id); err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(200, gin.H{"message": "customer deleted"})
}
