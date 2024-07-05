package repository

import (
	"booking-service/internal/repository/room_availability/storage"
	"context"
	"time"
)

type IRoomAvailability interface {
	GetItemsByHotelIDAndRoomID(ctx context.Context, hotelID, roomID string) []storage.ILockableRoomAvailability
	DecreaseQuota(ctx context.Context, hotelID, roomID string, date time.Time)
}
