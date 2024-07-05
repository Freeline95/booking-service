package room_availability

import (
	"booking-service/internal/repository/room_availability/storage"
	"context"
	"time"

	def "booking-service/internal/repository"
)

var _ def.IRoomAvailability = &RoomAvailabilityRepository{}

type RoomAvailabilityRepository struct {
	storage storage.IStorage
}

func NewRoomAvailabilityRepository(storage storage.IStorage) *RoomAvailabilityRepository {
	repo := &RoomAvailabilityRepository{
		storage: storage,
	}

	return repo
}

func (ra *RoomAvailabilityRepository) GetItemsByHotelIDAndRoomID(ctx context.Context, hotelID, roomID string) []storage.ILockableRoomAvailability {
	return ra.storage.GetItemsByHotelIDAndRoomID(ctx, hotelID, roomID)
}

func (ra *RoomAvailabilityRepository) DecreaseQuota(ctx context.Context, hotelID, roomID string, date time.Time) {
	ra.storage.DecreaseQuota(ctx, hotelID, roomID, date)
}
