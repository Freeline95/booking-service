package storage

import (
	repo_model_room_availability "booking-service/internal/repository/room_availability/model"
	"sync"
)

type ILockableRoomAvailability interface {
	GetRoomAvailabilityModel() *repo_model_room_availability.RoomAvailability
	Lock()
	Unlock()
}

type LockableRoomAvailability struct {
	sync.Mutex
	value *repo_model_room_availability.RoomAvailability
}

func NewLockableRoomAvailability(value *repo_model_room_availability.RoomAvailability) *LockableRoomAvailability {
	return &LockableRoomAvailability{
		value: value,
	}
}

func (li *LockableRoomAvailability) Lock() {
	li.Mutex.Lock()
}

func (li *LockableRoomAvailability) Unlock() {
	li.Mutex.Unlock()
}

func (li *LockableRoomAvailability) GetRoomAvailabilityModel() *repo_model_room_availability.RoomAvailability {
	return li.value
}
