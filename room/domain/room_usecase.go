package domain

import (
	"context"
)

type RoomUsecase interface {
	GetRoom(ctx context.Context, id uint) (Room, error)
	CreateRoom(ctx context.Context, Room *Room) (Room, error)
	UpdateRoom(ctx context.Context, Room *Room) (Room, error)
	DeleteRoom(ctx context.Context, id uint) error

	SearchRoomByStatus(ctx context.Context, status RoomStatus) ([]Room, error)
	SearchRoomByPrice(ctx context.Context, minPrice float64, maxPrice float64) ([]Room, error)
}

type RoomRepository interface {
	GetRoom(ctx context.Context, id uint) (Room, error)
	CreateRoom(ctx context.Context, Room *Room) (Room, error)
	UpdateRoom(ctx context.Context, Room *Room) (Room, error)
	DeleteRoom(ctx context.Context, id uint) error

	SearchRoomByStatus(ctx context.Context, status RoomStatus) ([]Room, error)
	SearchRoomByPrice(ctx context.Context, minPrice float64, maxPrice float64) ([]Room, error)
}
