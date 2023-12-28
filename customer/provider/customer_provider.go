package provider

import (
	"context"
	"hotel-booking/customer/domain"
	"log"

	"gorm.io/gorm"
)

type customerProvider struct {
	db *gorm.DB
}

func NewCustomerRepository(db *gorm.DB) domain.CustomerRepository {
	db.AutoMigrate(&domain.Customer{})
	return customerProvider{
		db: db,
	}
}

// CreateCustomer implements domain.CustomerRepository.
func (p customerProvider) CreateCustomer(ctx context.Context, customer *domain.Customer) (domain.Customer, error) {
	res := p.db.Create(customer)
	log.Println(res.Error)
	return *customer, res.Error
}

// DeleteCustomer implements domain.CustomerRepository.
func (p customerProvider) DeleteCustomer(ctx context.Context, id uint) error {
	res := p.db.Delete(&domain.Customer{}, id)
	return res.Error
}

// GetCustomer implements domain.CustomerRepository.
func (p customerProvider) GetCustomer(ctx context.Context, id uint) (domain.Customer, error) {
	var customer domain.Customer
	res := p.db.First(&customer, id)
	return customer, res.Error
}

// SearchCustomerByName implements domain.CustomerRepository.
func (p customerProvider) SearchCustomerByName(ctx context.Context, name string) ([]domain.Customer, error) {
	var customers []domain.Customer
	res := p.db.Where("name LIKE ?", "%"+name+"%").Find(&customers)
	return customers, res.Error
}

// SearchCustomerByPhone implements domain.CustomerRepository.
func (p customerProvider) SearchCustomerByPhone(ctx context.Context, phoneNumber string) ([]domain.Customer, error) {
	var customers []domain.Customer
	res := p.db.Where("phone_number LIKE ?", "%"+phoneNumber+"%").Find(&customers)
	return customers, res.Error
}

// UpdateCustomer implements domain.CustomerRepository.
func (p customerProvider) UpdateCustomer(ctx context.Context, customer *domain.Customer) (domain.Customer, error) {
	res := p.db.Model(&customer).Updates(customer)
	return *customer, res.Error
}
