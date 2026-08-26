package customer

import (
	"gorm.io/gorm"
)

type CustomerRepository struct {
	db *gorm.DB
}

func NewCustomerRepository(db *gorm.DB) *CustomerRepository {
	return &CustomerRepository{db: db}
}

func (r *CustomerRepository) Insert(customer *Customer) error {
	return r.db.Create(&customer).Error
}

func (r *CustomerRepository) FindByDocument(document string) (*Customer, error) {
	var customer Customer
	if err := r.db.First(&customer, "document = ?", document).Error; err != nil {
		return nil, err
	}
	return &customer, nil
}

func (r *CustomerRepository) FindById(id string) (*Customer, error) {
	var customer Customer
	if err := r.db.Where("id = ?", id).First(&customer).Error; err != nil {
		return nil, err
	}
	return &customer, nil
}

func (r *CustomerRepository) FindByEmail(email string) (*Customer, error) {
	var customer Customer
	if err := r.db.Where("email = ?", email).First(&customer).Error; err != nil {
		return nil, err
	}
	return &customer, nil
}

func (r *CustomerRepository) FindAll() ([]Customer, error) {
	var customers []Customer
	if err := r.db.Find(&customers).Error; err != nil {
		return nil, err
	}
	return customers, nil
}

func (r *CustomerRepository) Update(customer *Customer) error {
	return r.db.Save(&customer).Error
}

func (r *CustomerRepository) Delete(customer *Customer) error {
	return r.db.Delete(customer).Error
}
