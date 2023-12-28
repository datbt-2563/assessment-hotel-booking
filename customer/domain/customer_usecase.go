package domain

import (
	"context"
)

type CustomerUsecase interface {
	GetCustomer(ctx context.Context, id uint) (Customer, error)
	CreateCustomer(ctx context.Context, customer *Customer) (Customer, error)
	UpdateCustomer(ctx context.Context, customer *Customer) (Customer, error)
	DeleteCustomer(ctx context.Context, id uint) error

	SearchCustomerByName(ctx context.Context, name string) ([]Customer, error)
	SearchCustomerByPhone(ctx context.Context, phoneNumber string) ([]Customer, error)
}

type CustomerRepository interface {
	GetCustomer(ctx context.Context, id uint) (Customer, error)
	CreateCustomer(ctx context.Context, customer *Customer) (Customer, error)
	UpdateCustomer(ctx context.Context, customer *Customer) (Customer, error)
	DeleteCustomer(ctx context.Context, id uint) error

	SearchCustomerByName(ctx context.Context, name string) ([]Customer, error)
	SearchCustomerByPhone(ctx context.Context, phoneNumber string) ([]Customer, error)
}
