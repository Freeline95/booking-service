package order

import (
	"booking-service/pkg/business_errors"
	common_error "booking-service/pkg/error"
	common_validator "booking-service/pkg/validator"
	"encoding/json"
	"net/http"
	"time"
)

type CreateOrderRequest struct {
	CreateOrderData
}

type CreateOrderData struct {
	HotelID   string    `json:"hotel_id"`
	RoomID    string    `json:"room_id"`
	UserEmail string    `json:"email"`
	From      time.Time `json:"from"`
	To        time.Time `json:"to"`
}

type CreateOrderResponse struct {
}

func ParseAndValidateCreateOrderRequest(r *http.Request) (*CreateOrderRequest, error) {
	var request CreateOrderRequest

	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		return nil, common_error.Annotate(err, "Error while decode request to CreateOrderRequest")
	}

	if !common_validator.CheckIfEmail(request.UserEmail) {
		return nil, business_errors.InvalidEmailBusinessError
	}

	return &request, nil
}
