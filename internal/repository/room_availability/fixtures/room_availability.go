package fixtures

import (
	repo_model_room_availability "booking-service/internal/repository/room_availability/model"
	common_date "booking-service/pkg/date"
)

var RoomsAvailabilities = []repo_model_room_availability.RoomAvailability{
	{"reddison", "lux", common_date.NewDate(2024, 1, 1), 1},
	{"reddison", "lux", common_date.NewDate(2024, 1, 2), 1},
	{"reddison", "lux", common_date.NewDate(2024, 1, 3), 1},
	{"reddison", "lux", common_date.NewDate(2024, 1, 4), 1},
	{"reddison", "lux", common_date.NewDate(2024, 1, 5), 0},
}
