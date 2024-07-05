package service

import (
	"context"

	order_service_model "booking-service/internal/model/order"
)

type IOrderService interface {
	CreateOrder(ctx context.Context, data order_service_model.CreateOrderData) error
}
