package service

import (
	"errors"

	"github.com/shopspring/decimal"
	"gorm.io/gorm"
)

type Services struct {
	repo *ServiceRepository
}

func NewServices(repo *ServiceRepository) *Services {
	return &Services{repo: repo}
}

func (s *Services) Create(service *Service) error {
	if service.Price.LessThan(decimal.Zero) {
		return ErrInvalidPrice
	}

	existing, err := s.repo.FindByName(service.Name)
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return err
	}
	if existing != nil {
		return ErrServiceNameAlreadyExists
	}

	return s.repo.Insert(service)
}

func (s *Services) GetAll() ([]Service, error) {
	return s.repo.FindAll()
}

func (s *Services) GetByID(id string) (*Service, error) {
	svc, err := s.repo.FindByID(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ErrServiceNotFound
		}
		return nil, err
	}
	return svc, nil
}

func (s *Services) Update(id string, patch UpdateServiceRequest) (*Service, error) {
	existing, err := s.repo.FindByID(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ErrServiceNotFound
		}
		return nil, err
	}

	if patch.Name != nil && *patch.Name != existing.Name {
		other, err := s.repo.FindByName(*patch.Name)
		if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, err
		}
		if other != nil {
			return nil, ErrServiceNameAlreadyExists
		}
	}

	patch.ApplyTo(existing)

	if existing.Price.LessThan(decimal.Zero) {
		return nil, ErrInvalidPrice
	}

	if err := s.repo.Update(existing); err != nil {
		return nil, err
	}
	return existing, nil
}

func (s *Services) Delete(id string) error {
	if _, err := s.GetByID(id); err != nil {
		return err // já vem ErrServiceNotFound se for o caso
	}

	inUse, err := s.repo.IsUsedInOrders(id)
	if err != nil {
		return err
	}
	if inUse {
		return ErrServiceInUse
	}

	return s.repo.Delete(id)
}
