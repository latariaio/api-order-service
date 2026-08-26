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
	if customer.DocumentType != CustomerTypeIndividual && customer.DocumentType != CustomerTypeCompany {
		return ErrInvalidDocumentType
	}

	existCustomer, err := s.repo.FindByDocument(customer.Document)
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return err
	}

	if existCustomer != nil {
		return ErrCustomerAlreadyExists
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
	customer, err := s.repo.FindById(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ErrCustomerNotFound
		}
		return nil, err
	}

	if customer == nil {
		return nil, ErrCustomerNotFound
	}

	return customer, nil
}

func (s *CustomerService) Update(id string, patch UpdateCustomerRequest) (*Customer, error) {
	existing, err := s.repo.FindById(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ErrCustomerNotFound
		}
		return nil, err
	}
	if existing == nil {
		return nil, ErrCustomerNotFound
	}

	if patch.Email != nil && *patch.Email != existing.Email {
		other, err := s.repo.FindByEmail(*patch.Email)
		if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, err
		}
		if other != nil {
			return nil, ErrEmailAlreadyInUse
		}
	}

	patch.ApplyTo(existing)

	if err := s.repo.Update(existing); err != nil {
		return nil, err
	}

	return existing, nil
}

func (s *CustomerService) Delete(id string) error {
	customer, err := s.repo.FindById(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return ErrCustomerNotFound
		}
		return err
	}
	if customer == nil {
		return ErrCustomerNotFound
	}

	return s.repo.Delete(customer)
}
