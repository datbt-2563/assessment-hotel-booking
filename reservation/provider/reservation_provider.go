package provider

import (
	"context"
	"hotel-booking/reservation/domain"

	"gorm.io/gorm"
)

type reservationProvider struct {
	db *gorm.DB
}

func NewReservationRepository(db *gorm.DB) domain.ReservationRepository {
	db.AutoMigrate(&domain.Reservation{})
	return reservationProvider{
		db: db,
	}
}

// CreateReservation implements domain.ReservationRepository.
func (p reservationProvider) CreateReservation(ctx context.Context, reservation *domain.Reservation) (domain.Reservation, error) {
	res := p.db.Create(reservation)
	return *reservation, res.Error
}

// DeleteReservation implements domain.ReservationRepository.
func (p reservationProvider) DeleteReservation(ctx context.Context, id uint) error {
	res := p.db.Delete(&domain.Reservation{}, id)
	return res.Error
}

// GetReservation implements domain.ReservationRepository.
func (p reservationProvider) GetReservation(ctx context.Context, id uint) (domain.Reservation, error) {
	var reservation domain.Reservation
	res := p.db.First(&reservation, id)
	return reservation, res.Error
}

// SearchReservationByCustomer implements domain.ReservationRepository.
func (p reservationProvider) SearchReservationByCustomer(ctx context.Context, customerId uint) ([]domain.Reservation, error) {
	var reservations []domain.Reservation
	res := p.db.Where("customer_id = ?", customerId).Find(&reservations)
	return reservations, res.Error
}

// SearchReservationByStatus implements domain.ReservationRepository.
func (p reservationProvider) SearchReservationByStatus(ctx context.Context, status string) ([]domain.Reservation, error) {
	var reservations []domain.Reservation
	res := p.db.Where("status LIKE ?", status).Find(&reservations)
	return reservations, res.Error
}

// UpdateReservation implements domain.ReservationRepository.
func (p reservationProvider) UpdateReservation(ctx context.Context, reservation *domain.Reservation) (domain.Reservation, error) {
	res := p.db.Model(&reservation).Updates(reservation)
	return *reservation, res.Error
}
