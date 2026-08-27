package service

import "gorm.io/gorm"

type ServiceRepository struct {
	db *gorm.DB
}

func NewServiceRepository(db *gorm.DB) *ServiceRepository {
	return &ServiceRepository{db: db}
}

func (r *ServiceRepository) Insert(service *Service) error {
	return r.db.Create(service).Error
}

func (r *ServiceRepository) FindByID(id string) (*Service, error) {
	var service Service
	if err := r.db.Where("id = ?", id).First(&service).Error; err != nil {
		return nil, err
	}
	return &service, nil
}

func (r *ServiceRepository) FindByName(name string) (*Service, error) {
	var service Service
	if err := r.db.Where("name = ?", name).First(&service).Error; err != nil {
		return nil, err
	}
	return &service, nil
}

func (r *ServiceRepository) FindAll() ([]Service, error) {
	var services []Service
	if err := r.db.Find(&services).Error; err != nil {
		return nil, err
	}
	return services, nil
}

func (r *ServiceRepository) Update(service *Service) error {
	return r.db.Save(service).Error
}

func (r *ServiceRepository) Delete(id string) error {
	return r.db.Where("id = ?", id).Delete(&Service{}).Error
}

func (r *ServiceRepository) IsUsedInOrders(serviceID string) (bool, error) {
	var count int64
	err := r.db.Table("service_order_items").
		Where("service_id = ?", serviceID).
		Count(&count).Error
	if err != nil {
		return false, err
	}
	return count > 0, nil
}
