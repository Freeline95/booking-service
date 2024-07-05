package business_errors

import pkg_error "booking-service/pkg/error"

const (
	HotelRoomUnavailableForSelectedDatesCode = "hotel_room_unavailable_for_selected_dates"
	InvalidEmailCode                         = "invalid_email_code"
)

var (
	HotelRoomUnavailableForChosenDatesBusinessError = pkg_error.NewBusinessError(HotelRoomUnavailableForSelectedDatesCode, "Hotel room is not available for selected dates")
	InvalidEmailBusinessError                       = pkg_error.NewBusinessError(InvalidEmailCode, "Email is invalid")
)
