package service_order_item

import "gorm.io/gorm"

type ServiceOrderItemRepository struct {
	db *gorm.DB
}

func NewServiceOrderItemRepository(db *gorm.DB) *ServiceOrderItemRepository {
	return &ServiceOrderItemRepository{db: db}
}

func (r *ServiceOrderItemRepository) Insert(item *ServiceOrderItem) error {
	return r.db.Create(item).Error
}

func (r *ServiceOrderItemRepository) FindAll() ([]ServiceOrderItem, error) {
	var items []ServiceOrderItem
	return items, r.db.Find(&items).Error
}

func (r *ServiceOrderItemRepository) FindById(id string) (*ServiceOrderItem, error) {
	var item ServiceOrderItem
	return &item, r.db.First(&item, "id = ?", id).Error
}

func (r *ServiceOrderItemRepository) FindByServiceOrderId(serviceOrderId string) ([]ServiceOrderItem, error) {
	var items []ServiceOrderItem
	return items, r.db.Find(&items, "service_order_id = ?", serviceOrderId).Error
}

func (r *ServiceOrderItemRepository) Update(item *ServiceOrderItem) error {
	return r.db.Where("id = ?", item.ID).Save(item).Error
}

func (r *ServiceOrderItemRepository) Delete(id string) error {
	return r.db.Where("id = ?", id).Delete(&ServiceOrderItem{}).Error
}
