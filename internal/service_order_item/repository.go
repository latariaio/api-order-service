package service_order_item

import (
	"github.com/shopspring/decimal"
	"gorm.io/gorm"
)

type ServiceOrderItemRepository struct {
	db *gorm.DB
}

func NewServiceOrderItemRepository(db *gorm.DB) *ServiceOrderItemRepository {
	return &ServiceOrderItemRepository{db: db}
}

func (r *ServiceOrderItemRepository) Insert(item *ServiceOrderItem) error {
	return r.db.Create(item).Error
}

func (r *ServiceOrderItemRepository) FindByID(id string) (*ServiceOrderItem, error) {
	var item ServiceOrderItem
	if err := r.db.Where("id = ?", id).First(&item).Error; err != nil {
		return nil, err
	}
	return &item, nil
}

func (r *ServiceOrderItemRepository) FindByServiceOrderID(orderID string) ([]ServiceOrderItem, error) {
	var items []ServiceOrderItem
	if err := r.db.Where("service_order_id = ?", orderID).Find(&items).Error; err != nil {
		return nil, err
	}
	return items, nil
}

func (r *ServiceOrderItemRepository) Update(item *ServiceOrderItem) error {
	return r.db.Save(item).Error
}

func (r *ServiceOrderItemRepository) Delete(id string) error {
	return r.db.Where("id = ?", id).Delete(&ServiceOrderItem{}).Error
}

// SumTotalByServiceOrderID soma o total_price de todos os itens ativos de uma OS
func (r *ServiceOrderItemRepository) SumTotalByServiceOrderID(orderID string) (decimal.Decimal, error) {
	items, err := r.FindByServiceOrderID(orderID)
	if err != nil {
		return decimal.Zero, err
	}

	total := decimal.Zero
	for _, item := range items {
		total = total.Add(item.TotalPrice)
	}
	return total, nil
}
