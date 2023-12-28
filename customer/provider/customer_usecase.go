package provider

import (
	"context"
	"hotel-booking/customer/domain"
)

type customerUsecase struct {
	customerRepo domain.CustomerRepository
}

// CreateCustomer implements domain.CustomerUsecase.
func (u *customerUsecase) CreateCustomer(ctx context.Context, customer *domain.Customer) (domain.Customer, error) {
	return u.customerRepo.CreateCustomer(ctx, customer)
}

// DeleteCustomer implements domain.CustomerUsecase.
func (u *customerUsecase) DeleteCustomer(ctx context.Context, id uint) error {
	return u.customerRepo.DeleteCustomer(ctx, id)
}

// GetCustomer implements domain.CustomerUsecase.
func (u *customerUsecase) GetCustomer(ctx context.Context, id uint) (domain.Customer, error) {
	return u.customerRepo.GetCustomer(ctx, id)
}

// SearchCustomerByName implements domain.CustomerUsecase.
func (u *customerUsecase) SearchCustomerByName(ctx context.Context, name string) ([]domain.Customer, error) {
	return u.customerRepo.SearchCustomerByName(ctx, name)
}

// SearchCustomerByPhone implements domain.CustomerUsecase.
func (u *customerUsecase) SearchCustomerByPhone(ctx context.Context, phoneNumber string) ([]domain.Customer, error) {
	return u.customerRepo.SearchCustomerByPhone(ctx, phoneNumber)
}

// UpdateCustomer implements domain.CustomerUsecase.
func (u *customerUsecase) UpdateCustomer(ctx context.Context, customer *domain.Customer) (domain.Customer, error) {
	return u.customerRepo.UpdateCustomer(ctx, customer)
}

func NewCustomerUsecase(r domain.CustomerRepository) domain.CustomerUsecase {
	return &customerUsecase{
		customerRepo: r,
	}
}
