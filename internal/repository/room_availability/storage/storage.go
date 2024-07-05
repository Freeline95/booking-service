package storage

import (
	"context"
	"sort"
	"time"
)

var _ IStorage = &Storage{}

type IStorage interface {
	GetItemsByHotelIDAndRoomID(ctx context.Context, hotelID, roomID string) []ILockableRoomAvailability
	DecreaseQuota(ctx context.Context, hotelID, roomID string, date time.Time)
}

type Storage struct {
	indexByHotelIDAndRoomID map[string][]ILockableRoomAvailability
	indexByPrimaryKey       map[string]ILockableRoomAvailability
}

func NewStorage() *Storage {
	storage := &Storage{
		indexByHotelIDAndRoomID: make(map[string][]ILockableRoomAvailability),
		indexByPrimaryKey:       make(map[string]ILockableRoomAvailability),
	}
	storage.loadDataFromFixtures()

	return storage
}

func (s *Storage) loadDataFromFixtures() {
	sort.Slice(fixtures, func(first, second int) bool {
		return fixtures[first].Date.Before(fixtures[second].Date)
	})

	for _, item := range fixtures {
		newItem := item
		lockableRoomAvailability := NewLockableRoomAvailability(&newItem)

		keyByHotelAndRoom := item.HotelID + item.RoomID
		s.indexByHotelIDAndRoomID[keyByHotelAndRoom] = append(s.indexByHotelIDAndRoomID[keyByHotelAndRoom], lockableRoomAvailability)

		keyByHotelAndRoomAndDate := keyByHotelAndRoom + item.Date.Format("2006-01-02 15:04:05")
		s.indexByPrimaryKey[keyByHotelAndRoomAndDate] = lockableRoomAvailability
	}
}

func (s *Storage) GetItemsByHotelIDAndRoomID(ctx context.Context, hotelID, roomID string) []ILockableRoomAvailability {
	key := hotelID + roomID

	return s.indexByHotelIDAndRoomID[key]
}

func (s *Storage) DecreaseQuota(ctx context.Context, hotelID, roomID string, date time.Time) {
	key := hotelID + roomID + date.Format("2006-01-02 15:04:05")

	lockableRoomAvailable := s.indexByPrimaryKey[key]
	lockableRoomAvailable.Lock()
	defer lockableRoomAvailable.Unlock()

	s.indexByPrimaryKey[key].GetRoomAvailabilityModel().Quota--
}
