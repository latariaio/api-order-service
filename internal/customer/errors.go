package customer

import "errors"

var (
	ErrInvalidDocumentType   = errors.New("invalid document type")
	ErrCustomerAlreadyExists = errors.New("customer already exists")
	ErrCustomerNotFound      = errors.New("customer not found")
	ErrInvalidCustomerID     = errors.New("invalid customer id")
	ErrEmailAlreadyInUse     = errors.New("email already in use")
)
