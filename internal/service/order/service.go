package order

import (
	order_service_model "booking-service/internal/model/order"
	"booking-service/internal/repository"
	"booking-service/internal/service"
	"booking-service/pkg/business_errors"
	common_log "booking-service/pkg/log"
	"context"
)

var _ service.IOrderService = &OrderService{}

type OrderService struct {
	roomAvailabilityRepository repository.IRoomAvailability
}

func NewOrderService(roomAvailabilityRepository repository.IRoomAvailability) *OrderService {
	return &OrderService{
		roomAvailabilityRepository: roomAvailabilityRepository,
	}
}

func (os *OrderService) CreateOrder(ctx context.Context, createOrderData order_service_model.CreateOrderData) error {
	// Noticed it is sorted availabilities by date
	availabilities := os.roomAvailabilityRepository.GetItemsByHotelIDAndRoomID(ctx, createOrderData.HotelID, createOrderData.RoomID)
	found := false

	// Could be solved by binary search, but the problem is we can not do it through availabilities contains Quota = 0
	// We could extract entities from repo without Quota = 0, but filtering takes Liniar time
	// Also the way - delete entity from storage when Quota reaches 0, then we could do binary search, but i was not sure i could delete it
	for _, availability := range availabilities {
		availabilityEntity := availability.GetRoomAvailabilityModel()
		availabilityDate := availabilityEntity.Date

		if availabilityDate.After(createOrderData.To) {
			break
		}

		if availabilityEntity.Quota != 0 && (availabilityDate.After(createOrderData.From) && availabilityDate.Before(createOrderData.To) ||
			availabilityDate.Equal(createOrderData.From) ||
			availabilityDate.Equal(createOrderData.To)) {

			os.roomAvailabilityRepository.DecreaseQuota(ctx, availabilityEntity.HotelID, availabilityEntity.RoomID, availabilityEntity.Date)

			found = true

			break
		}
	}

	if !found {
		common_log.Error("Hotel room is not available for selected dates: \n%v", createOrderData)

		return business_errors.HotelRoomUnavailableForChosenDatesBusinessError
	}

	return nil
}
