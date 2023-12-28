package domain

import (
	c "hotel-booking/customer/domain"
	r "hotel-booking/room/domain"
	"time"
)

type Reservation struct {
	ReservationId uint `gorm:"primaryKey"`
	CustomerId    uint
	Customer      c.Customer `gorm:"foreignKey:CustomerId"`
	RoomId        uint
	Room          r.Room `gorm:"foreignKey:RoomId"`
	CheckInDate   time.Time
	CheckOutDate  time.Time
	Occupancy     int
	Status        string
	CreatedAt     time.Time
	UpdatedAt     time.Time
}
