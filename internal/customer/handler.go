package customer

import (
	"log"

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

type CreateCustomerRequest struct {
	Name     string `json:"name"`
	Document string `json:"document"`
	Phone    string `json:"phone"`
	Email    string `json:"email"`
	Address  string `json:"address"`
}

func (h *CustomerHandler) CreateCustomer(ctx *gin.Context) {
	var request CreateCustomerRequest
	if err := ctx.BindJSON(&request); err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}

	customer := Customer{
		Name:     request.Name,
		Document: request.Document,
		Phone:    request.Phone,
		Email:    request.Email,
		Address:  request.Address,
	}

	log.Default().Println("creating customer", customer)

	if err := h.service.Create(&customer); err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(200, customer)
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
	ctx.JSON(200, customers)
}

func (h *CustomerHandler) UpdateCustomer(ctx *gin.Context) {
	id := ctx.Param("id")
	var request CreateCustomerRequest
	if err := ctx.BindJSON(&request); err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}
	customer := Customer{
		Name:     request.Name,
		Document: request.Document,
		Phone:    request.Phone,
		Email:    request.Email,
		Address:  request.Address,
	}
	if err := h.service.Update(id, &customer); err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(200, customer)
}

func (h *CustomerHandler) DeleteCustomer(ctx *gin.Context) {
	id := ctx.Param("id")
	if err := h.service.Delete(id); err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(200, gin.H{"message": "customer deleted"})
}
