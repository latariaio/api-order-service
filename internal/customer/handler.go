package customer

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/latariaio/api-order-service/internal/httputil"
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
	var request CreateCustomerRequest
	if err := ctx.ShouldBindJSON(&request); err != nil {
		if fieldErrors := httputil.FormatValidationError(err); fieldErrors != nil {
			ctx.JSON(400, gin.H{"errors": fieldErrors})
			return
		}

		ctx.JSON(400, gin.H{"error": "invalid request body"})
		return
	}

	customer := request.ToModel()

	if err := h.service.Create(&customer); err != nil {
		switch {
		case errors.Is(err, ErrInvalidDocumentType):
			ctx.JSON(400, gin.H{"error": err.Error()})
		case errors.Is(err, ErrCustomerAlreadyExists):
			ctx.JSON(409, gin.H{"error": err.Error()})
		case errors.Is(err, ErrCustomerNotFound):
			ctx.JSON(404, gin.H{"error": err.Error()})
		default:
			ctx.JSON(500, gin.H{"error": "internal server error"})
		}
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"data": ToCustomerResponse(customer)})
}

func (h *CustomerHandler) GetCustomer(ctx *gin.Context) {
	id := ctx.Param("id")

	customer, err := h.service.FindById(id)
	if err != nil {

		switch {
		case errors.Is(err, ErrCustomerNotFound):
			ctx.JSON(404, gin.H{"error": err.Error()})
		default:
			ctx.JSON(500, gin.H{"error": "internal server error"})
		}

		return
	}
	ctx.JSON(200, gin.H{"data": ToCustomerResponse(*customer)})
}

func (h *CustomerHandler) GetCustomers(ctx *gin.Context) {
	customers, err := h.service.FindAll()
	if err != nil {
		ctx.JSON(500, gin.H{"error": "internal server error"})
		return
	}
	ctx.JSON(200, gin.H{"data": ToCustomerResponseList(customers)})
}

func (h *CustomerHandler) UpdateCustomer(ctx *gin.Context) {
	id := ctx.Param("id")

	var request UpdateCustomerRequest
	if err := ctx.ShouldBindJSON(&request); err != nil {
		if fieldErrors := httputil.FormatValidationError(err); fieldErrors != nil {
			ctx.JSON(400, gin.H{"errors": fieldErrors})
			return
		}
		ctx.JSON(400, gin.H{"error": "invalid request body"})
		return
	}

	customer, err := h.service.Update(id, request)
	if err != nil {
		switch {
		case errors.Is(err, ErrCustomerNotFound):
			ctx.JSON(404, gin.H{"error": err.Error()})
		case errors.Is(err, ErrEmailAlreadyInUse):
			ctx.JSON(409, gin.H{"error": err.Error()})
		default:
			ctx.JSON(500, gin.H{"error": "internal server error"})
		}
		return
	}

	ctx.JSON(200, gin.H{"data": ToCustomerResponse(*customer)})
}

func (h *CustomerHandler) DeleteCustomer(ctx *gin.Context) {
	id := ctx.Param("id")

	if err := h.service.Delete(id); err != nil {
		switch {
		case errors.Is(err, ErrCustomerNotFound):
			ctx.JSON(404, gin.H{"error": err.Error()})
		default:
			ctx.JSON(500, gin.H{"error": "internal server error"})
		}
		return
	}

	ctx.Status(204)
}
