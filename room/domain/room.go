package domain

import "time"

type Room struct {
	RoomId       uint `gorm:"primaryKey"`
	RoomType     string
	Price        float64
	MaxOccupancy int
	Status       RoomStatus
	Email        string
	Address      string
	CreatedAt    time.Time
	UpdatedAt    time.Time
}

type RoomStatus string

const (
	RoomAvailable   = "Available"
	RoomOccupied    = "Occupied"
	RoomCleaning    = "Cleaning"
	RoomMaintenance = "Maintenance"
	RoomReserved    = "Reserved"
	RoomHold        = "Hold"
	RoomCanceled    = "Canceled"
)
