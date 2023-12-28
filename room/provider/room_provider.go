package provider

import (
	"context"
	"hotel-booking/room/domain"
	"log"

	"gorm.io/gorm"
)

type roomProvider struct {
	db *gorm.DB
}

func NewRoomRepository(db *gorm.DB) domain.RoomRepository {
	db.AutoMigrate(&domain.Room{})
	return roomProvider{
		db: db,
	}
}

// CreateRoom implements domain.RoomRepository.
func (p roomProvider) CreateRoom(ctx context.Context, room *domain.Room) (domain.Room, error) {
	res := p.db.Create(room)
	log.Println(res.Error)
	return *room, res.Error
}

// DeleteRoom implements domain.RoomRepository.
func (p roomProvider) DeleteRoom(ctx context.Context, id uint) error {
	res := p.db.Delete(&domain.Room{}, id)
	return res.Error
}

// GetRoom implements domain.RoomRepository.
func (p roomProvider) GetRoom(ctx context.Context, id uint) (domain.Room, error) {
	var room domain.Room
	res := p.db.First(&room, id)
	return room, res.Error
}

// UpdateRoom implements domain.RoomRepository.
func (p roomProvider) UpdateRoom(ctx context.Context, room *domain.Room) (domain.Room, error) {
	res := p.db.Model(&room).Updates(room)
	return *room, res.Error
}

// SearchRoomByPrice implements domain.RoomRepository.
func (p roomProvider) SearchRoomByPrice(ctx context.Context, minPrice float64, maxPrice float64) ([]domain.Room, error) {
	var rooms []domain.Room
	res := p.db.Where("price BETWEEN ? AND ?", minPrice, maxPrice).Find(&rooms)
	return rooms, res.Error
}

// SearchRoomByStatus implements domain.RoomRepository.
func (p roomProvider) SearchRoomByStatus(ctx context.Context, status domain.RoomStatus) ([]domain.Room, error) {
	var rooms []domain.Room
	res := p.db.Where("status LIKE ?", status).Find(&rooms)
	return rooms, res.Error
}
