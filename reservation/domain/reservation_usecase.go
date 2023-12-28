package domain

import (
	"context"
)

type ReservationUsecase interface {
	GetReservation(ctx context.Context, id uint) (Reservation, error)
	CreateReservation(ctx context.Context, reservation *Reservation) (Reservation, error)
	UpdateReservation(ctx context.Context, reservation *Reservation) (Reservation, error)
	DeleteReservation(ctx context.Context, id uint) error

	SearchReservationByStatus(ctx context.Context, status string) ([]Reservation, error)
	SearchReservationByCustomer(ctx context.Context, customerId uint) ([]Reservation, error)
}

type ReservationRepository interface {
	GetReservation(ctx context.Context, id uint) (Reservation, error)
	CreateReservation(ctx context.Context, reservation *Reservation) (Reservation, error)
	UpdateReservation(ctx context.Context, reservation *Reservation) (Reservation, error)
	DeleteReservation(ctx context.Context, id uint) error

	SearchReservationByStatus(ctx context.Context, status string) ([]Reservation, error)
	SearchReservationByCustomer(ctx context.Context, customerId uint) ([]Reservation, error)
}
