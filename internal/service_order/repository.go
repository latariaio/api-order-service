package service_order

import "gorm.io/gorm"

type ServiceOrderRepository struct {
	db *gorm.DB
}

func NewServiceOrderRepository(db *gorm.DB) *ServiceOrderRepository {
	return &ServiceOrderRepository{db: db}
}

func (r *ServiceOrderRepository) Insert(serviceOrder *ServiceOrder) error {
	return r.db.Create(serviceOrder).Error
}

func (r *ServiceOrderRepository) FindByID(id string) (*ServiceOrder, error) {
	var serviceOrder ServiceOrder
	result := r.db.Where("id = ?", id).First(&serviceOrder)
	return &serviceOrder, result.Error
}

func (r *ServiceOrderRepository) FindAll() ([]ServiceOrder, error) {
	var serviceOrders []ServiceOrder
	result := r.db.Find(&serviceOrders)
	return serviceOrders, result.Error
}

func (r *ServiceOrderRepository) Update(id string, serviceOrder *ServiceOrder) error {
	return r.db.Model(&ServiceOrder{}).Where("id = ?", id).Updates(serviceOrder).Error
}

func (r *ServiceOrderRepository) Delete(id string) error {
	return r.db.Where("id = ?", id).Delete(&ServiceOrder{}).Error
}
