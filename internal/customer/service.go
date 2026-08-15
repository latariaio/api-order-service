package customer

import (
	"errors"

	"gorm.io/gorm"
)

type CustomerService struct {
	repo *CustomerRepository
}

func NewCustomerService(repo *CustomerRepository) *CustomerService {
	return &CustomerService{
		repo: repo,
	}
}

func (s *CustomerService) Create(customer *Customer) error {
	existCustomer, err := s.repo.FindByDocument(customer.Document)
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return err
	}
	if existCustomer != nil {
		return errors.New("customer already exists")
	}
	return s.repo.Insert(customer)
}

func (s *CustomerService) FindAll() ([]Customer, error) {
	customers, err := s.repo.FindAll()
	if err != nil {
		return nil, err
	}

	return customers, nil
}

func (s *CustomerService) FindById(id string) (*Customer, error) {
	return s.repo.FindById(id)
}

func (s *CustomerService) Update(id string, customer *Customer) error {
	existCustomer, err := s.repo.FindById(id)
	if err != nil {
		return err
	}
	if existCustomer == nil {
		return errors.New("customer not found")
	}
	customer.ID = existCustomer.ID
	return s.repo.Update(customer)
}

func (s *CustomerService) Delete(id string) error {
	customer, err := s.repo.FindById(id)
	if err != nil {
		return err
	}
	if customer == nil {
		return errors.New("customer not found")
	}

	return s.repo.Delete(id)
}
