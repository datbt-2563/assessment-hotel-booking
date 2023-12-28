package provider

import (
	"context"
	"hotel-booking/room/domain"
)

type roomUsecase struct {
	roomRepo domain.RoomRepository
}

// SearchRoomByPrice implements domain.RoomUsecase.
func (u *roomUsecase) SearchRoomByPrice(ctx context.Context, minPrice float64, maxPrice float64) ([]domain.Room, error) {
	return u.roomRepo.SearchRoomByPrice(ctx, minPrice, maxPrice)
}

// SearchRoomByStatus implements domain.RoomUsecase.
func (u *roomUsecase) SearchRoomByStatus(ctx context.Context, status domain.RoomStatus) ([]domain.Room, error) {
	return u.roomRepo.SearchRoomByStatus(ctx, status)
}

// CreateRoom implements domain.RoomUsecase.
func (u *roomUsecase) CreateRoom(ctx context.Context, room *domain.Room) (domain.Room, error) {
	return u.roomRepo.CreateRoom(ctx, room)
}

// DeleteRoom implements domain.RoomUsecase.
func (u *roomUsecase) DeleteRoom(ctx context.Context, id uint) error {
	return u.roomRepo.DeleteRoom(ctx, id)
}

// GetRoom implements domain.RoomUsecase.
func (u *roomUsecase) GetRoom(ctx context.Context, id uint) (domain.Room, error) {
	return u.roomRepo.GetRoom(ctx, id)
}

// UpdateRoom implements domain.RoomUsecase.
func (u *roomUsecase) UpdateRoom(ctx context.Context, room *domain.Room) (domain.Room, error) {
	return u.roomRepo.UpdateRoom(ctx, room)
}

func NewRoomUsecase(r domain.RoomRepository) domain.RoomUsecase {
	return &roomUsecase{
		roomRepo: r,
	}
}
