package controller

import (
	"context"
	"net/http"

	order_service_model "booking-service/internal/model/order"
	"booking-service/internal/service"
	common_error "booking-service/pkg/error"
)

type IOrderController interface {
	CreateOrder(ctx context.Context, r *http.Request) (interface{}, error)
}

type OrderController struct {
	orderService service.IOrderService
}

func NewOrderController(orderService service.IOrderService) *OrderController {
	return &OrderController{
		orderService: orderService,
	}
}

func (oc *OrderController) CreateOrder(ctx context.Context, r *http.Request) (interface{}, error) {
	request, err := order_service_model.ParseAndValidateCreateOrderRequest(r)
	if err != nil {
		return nil, common_error.Annotate(err, "Error while parse and validate create order request")
	}

	if err = oc.orderService.CreateOrder(ctx, request.CreateOrderData); err != nil {
		return nil, common_error.Annotate(err, "Error while create order in service")
	}

	return &order_service_model.CreateOrderResponse{}, nil
}
