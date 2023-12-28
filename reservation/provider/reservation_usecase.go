package provider

import (
	"context"
	"hotel-booking/reservation/domain"
)

type reservationUsecase struct {
	reservationRepo domain.ReservationRepository
}

// CreateReservation implements domain.ReservationUsecase.
func (u *reservationUsecase) CreateReservation(ctx context.Context, reservation *domain.Reservation) (domain.Reservation, error) {
	return u.reservationRepo.CreateReservation(ctx, reservation)
}

// DeleteReservation implements domain.ReservationUsecase.
func (u *reservationUsecase) DeleteReservation(ctx context.Context, id uint) error {
	return u.reservationRepo.DeleteReservation(ctx, id)
}

// GetReservation implements domain.ReservationUsecase.
func (u *reservationUsecase) GetReservation(ctx context.Context, id uint) (domain.Reservation, error) {
	return u.reservationRepo.GetReservation(ctx, id)
}

// SearchReservationByCustomer implements domain.ReservationUsecase.
func (u *reservationUsecase) SearchReservationByCustomer(ctx context.Context, customerId uint) ([]domain.Reservation, error) {
	return u.reservationRepo.SearchReservationByCustomer(ctx, customerId)
}

// SearchReservationByStatus implements domain.ReservationUsecase.
func (u *reservationUsecase) SearchReservationByStatus(ctx context.Context, status string) ([]domain.Reservation, error) {
	return u.reservationRepo.SearchReservationByStatus(ctx, status)
}

// UpdateReservation implements domain.ReservationUsecase.
func (u *reservationUsecase) UpdateReservation(ctx context.Context, reservation *domain.Reservation) (domain.Reservation, error) {
	return u.reservationRepo.UpdateReservation(ctx, reservation)
}

func NewReservationUsecase(r domain.ReservationRepository) domain.ReservationUsecase {
	return &reservationUsecase{
		reservationRepo: r,
	}
}
