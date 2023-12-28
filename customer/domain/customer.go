package domain

import (
	"time"
)

type Customer struct {
	CustomerId  uint `gorm:"primaryKey"`
	Name        string
	PhoneNumber string
	Email       string
	Address     string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
