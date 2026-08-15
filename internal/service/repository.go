package service

import "gorm.io/gorm"

type ServiceRepository struct {
	db *gorm.DB
}

func NewServiceRepository(db *gorm.DB) *ServiceRepository {
	return &ServiceRepository{db: db}
}

func (r *ServiceRepository) Insert(service *Service) error {
	return r.db.Create(&service).Error
}

func (r *ServiceRepository) FindByID(id string) (*Service, error) {
	var service Service
	if err := r.db.Where("id = ?", id).First(&service).Error; err != nil {
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

func (r *ServiceRepository) Update(id string, service *Service) error {
	return r.db.Model(&Service{}).Where("id = ?", id).Updates(service).Error
}

func (r *ServiceRepository) Delete(id string) error {
	return r.db.Where("id = ?", id).Delete(&Service{}).Error
}
