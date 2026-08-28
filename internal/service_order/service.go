package service_order

import (
	"errors"
	"time"

	"gorm.io/gorm"
)

type ServiceOrderService struct {
	repo *ServiceOrderRepository
}

func NewServiceOrderService(repo *ServiceOrderRepository) *ServiceOrderService {
	return &ServiceOrderService{repo: repo}
}

func (s *ServiceOrderService) Create(order *ServiceOrder) error {
	exists, err := s.repo.CustomerExists(order.CustomerID.String())
	if err != nil {
		return err
	}
	if !exists {
		return ErrCustomerNotFound
	}

	order.Status = StatusOpen
	order.OpenedAt = time.Now()

	return s.repo.Insert(order)
}

func (s *ServiceOrderService) GetAll() ([]ServiceOrder, error) {
	return s.repo.FindAll()
}

func (s *ServiceOrderService) GetByID(id string) (*ServiceOrder, error) {
	order, err := s.repo.FindByID(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ErrServiceOrderNotFound
		}
		return nil, err
	}
	return order, nil
}

func (s *ServiceOrderService) Update(id string, patch UpdateServiceOrderRequest) (*ServiceOrder, error) {
	order, err := s.GetByID(id)
	if err != nil {
		return nil, err
	}

	if order.Status.IsClosed() {
		return nil, ErrServiceOrderClosed
	}

	patch.ApplyTo(order)

	if err := s.repo.Update(order); err != nil {
		return nil, err
	}
	return order, nil
}

func (s *ServiceOrderService) UpdateStatus(id string, next Status) (*ServiceOrder, error) {
	order, err := s.GetByID(id)
	if err != nil {
		return nil, err
	}

	if !next.IsValid() || !order.Status.CanTransitionTo(next) {
		return nil, ErrInvalidStatusTransition
	}

	order.Status = next

	if next == StatusCompleted {
		now := time.Now()
		order.CompletedAt = &now
	}

	if err := s.repo.Update(order); err != nil {
		return nil, err
	}
	return order, nil
}

var deletableStatuses = map[Status]bool{
	StatusOpen:       true,
	StatusInAnalysis: true,
}

var cancellableStatuses = map[Status]bool{
	StatusAwaitingApproval: true,
	StatusInProgress:       true,
}

func (s *ServiceOrderService) Delete(id string) (*ServiceOrder, error) {
	order, err := s.GetByID(id)
	if err != nil {
		return nil, err
	}

	switch {
	case deletableStatuses[order.Status]:
		return nil, s.repo.Delete(id)
	case cancellableStatuses[order.Status]:
		order.Status = StatusCancelled
		if err := s.repo.Update(order); err != nil {
			return nil, err
		}
		return order, nil
	default:
		return nil, ErrServiceOrderCannotDelete
	}
}
