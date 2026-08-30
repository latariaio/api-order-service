package service_order_item

import (
	"errors"

	"github.com/latariaio/api-order-service/internal/service"
	"github.com/latariaio/api-order-service/internal/service_order"
	"github.com/shopspring/decimal"
	"gorm.io/gorm"
)

type ServiceOrderItemService struct {
	repo        *ServiceOrderItemRepository
	orderRepo   *service_order.ServiceOrderRepository
	serviceRepo *service.ServiceRepository
}

func NewServiceOrderItemService(
	repo *ServiceOrderItemRepository,
	orderRepo *service_order.ServiceOrderRepository,
	serviceRepo *service.ServiceRepository,
) *ServiceOrderItemService {
	return &ServiceOrderItemService{repo: repo, orderRepo: orderRepo, serviceRepo: serviceRepo}
}

func (s *ServiceOrderItemService) AddItem(orderID string, request AddServiceOrderItemRequest) (*ServiceOrderItem, error) {
	order, err := s.orderRepo.FindByID(orderID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ErrServiceOrderNotFound
		}
		return nil, err
	}

	if order.Status.IsClosed() {
		return nil, ErrServiceOrderClosed
	}

	svc, err := s.serviceRepo.FindByID(request.ServiceID.String())
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ErrServiceNotFound
		}
		return nil, err
	}

	item := &ServiceOrderItem{
		ServiceOrderID: order.ID,
		ServiceID:      svc.ID,
		Quantity:       request.Quantity,
		UnitPrice:      svc.Price, // preço atual do catálogo — nunca o que o frontend mandar
	}
	item.TotalPrice = item.UnitPrice.Mul(decimal.NewFromInt(int64(item.Quantity)))

	if err := s.repo.Insert(item); err != nil {
		return nil, err
	}

	if err := s.recalculateOrderTotal(order); err != nil {
		return nil, err
	}

	return item, nil
}

func (s *ServiceOrderItemService) ListByOrder(orderID string) ([]ServiceOrderItem, error) {
	return s.repo.FindByServiceOrderID(orderID)
}

func (s *ServiceOrderItemService) UpdateItem(orderID, itemID string, request UpdateServiceOrderItemRequest) (*ServiceOrderItem, error) {
	order, err := s.orderRepo.FindByID(orderID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ErrServiceOrderNotFound
		}
		return nil, err
	}

	if order.Status.IsClosed() {
		return nil, ErrServiceOrderClosed
	}

	item, err := s.repo.FindByID(itemID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ErrServiceOrderItemNotFound
		}
		return nil, err
	}

	if item.ServiceOrderID != order.ID {
		return nil, ErrItemDoesNotBelongToOrder
	}

	item.Quantity = request.Quantity
	item.TotalPrice = item.UnitPrice.Mul(decimal.NewFromInt(int64(item.Quantity)))

	if err := s.repo.Update(item); err != nil {
		return nil, err
	}

	if err := s.recalculateOrderTotal(order); err != nil {
		return nil, err
	}

	return item, nil
}

func (s *ServiceOrderItemService) RemoveItem(orderID, itemID string) error {
	order, err := s.orderRepo.FindByID(orderID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return ErrServiceOrderNotFound
		}
		return err
	}

	if order.Status.IsClosed() {
		return ErrServiceOrderClosed
	}

	item, err := s.repo.FindByID(itemID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return ErrServiceOrderItemNotFound
		}
		return err
	}

	if item.ServiceOrderID != order.ID {
		return ErrItemDoesNotBelongToOrder
	}

	if err := s.repo.Delete(itemID); err != nil {
		return err
	}

	return s.recalculateOrderTotal(order)
}

func (s *ServiceOrderItemService) recalculateOrderTotal(order *service_order.ServiceOrder) error {
	total, err := s.repo.SumTotalByServiceOrderID(order.ID.String())
	if err != nil {
		return err
	}
	order.TotalPrice = total
	return s.orderRepo.Update(order)
}
