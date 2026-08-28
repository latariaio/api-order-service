package service_order

import "gorm.io/gorm"

type ServiceOrderRepository struct {
	db *gorm.DB
}

func NewServiceOrderRepository(db *gorm.DB) *ServiceOrderRepository {
	return &ServiceOrderRepository{db: db}
}

func (r *ServiceOrderRepository) Insert(so *ServiceOrder) error {
	return r.db.Create(so).Error
}

func (r *ServiceOrderRepository) FindByID(id string) (*ServiceOrder, error) {
	var so ServiceOrder
	if err := r.db.Where("id = ?", id).First(&so).Error; err != nil {
		return nil, err
	}
	return &so, nil
}

func (r *ServiceOrderRepository) FindAll() ([]ServiceOrder, error) {
	var list []ServiceOrder
	if err := r.db.Find(&list).Error; err != nil {
		return nil, err
	}
	return list, nil
}

func (r *ServiceOrderRepository) Update(so *ServiceOrder) error {
	return r.db.Save(so).Error
}

func (r *ServiceOrderRepository) Delete(id string) error {
	return r.db.Where("id = ?", id).Delete(&ServiceOrder{}).Error
}

// CustomerExists confirma que o cliente informado realmente existe (e não está soft-deleted)
func (r *ServiceOrderRepository) CustomerExists(customerID string) (bool, error) {
	var count int64
	err := r.db.Table("customers").
		Where("id = ? AND deleted_at IS NULL", customerID).
		Count(&count).Error
	if err != nil {
		return false, err
	}
	return count > 0, nil
}
