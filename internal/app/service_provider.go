package app

import (
	"booking-service/internal/config"
	"booking-service/internal/controller"
	"booking-service/internal/repository"
	room_availability_repository "booking-service/internal/repository/room_availability"
	room_availability_storage "booking-service/internal/repository/room_availability/storage"
	"booking-service/internal/service"
	order_service "booking-service/internal/service/order"
	common_log "booking-service/pkg/log"
	"github.com/gorilla/mux"
)

type ServiceProvider struct {
	config *config.Configuration

	orderController controller.IOrderController

	roomAvailabilityStorage room_availability_storage.IStorage

	roomAvailabilityRepository repository.IRoomAvailability

	orderService service.IOrderService

	router *mux.Router
}

func newServiceProvider(config *config.Configuration) *ServiceProvider {
	servicesProvider := &ServiceProvider{
		config: config,
	}

	return servicesProvider
}

func (sr *ServiceProvider) Shutdown() {
	common_log.Info("Shutting down...")

	common_log.Info("Shutdown completed")
}

func (sr *ServiceProvider) Router() *mux.Router {
	if sr.router == nil {
		sr.router = mux.NewRouter()
	}

	return sr.router
}

func (sr *ServiceProvider) RoomAvailabilityStorage() room_availability_storage.IStorage {
	if sr.roomAvailabilityStorage == nil {
		sr.roomAvailabilityStorage = room_availability_storage.NewStorage()
	}

	return sr.roomAvailabilityStorage
}

func (sr *ServiceProvider) RoomAvailabilityRepository() repository.IRoomAvailability {
	if sr.roomAvailabilityRepository == nil {
		sr.roomAvailabilityRepository = room_availability_repository.NewRoomAvailabilityRepository(
			sr.RoomAvailabilityStorage(),
		)
	}

	return sr.roomAvailabilityRepository
}

func (sr *ServiceProvider) OrderService() service.IOrderService {
	if sr.orderService == nil {
		sr.orderService = order_service.NewOrderService(
			sr.RoomAvailabilityRepository(),
		)
	}

	return sr.orderService
}

func (sr *ServiceProvider) OrderController() controller.IOrderController {
	if sr.orderController == nil {
		sr.orderController = controller.NewOrderController(
			sr.OrderService(),
		)
	}

	return sr.orderController
}
