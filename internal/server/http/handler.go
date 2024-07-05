package http

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"time"

	"booking-service/internal/config"
	common_error "booking-service/pkg/error"
)

type Handler struct {
	config *config.Configuration
}

type ApiErrorResponse struct {
	Error string  `json:"business_errors"`
	Code  *string `json:"code,omitempty"`
}

type Action func(ctx context.Context, r *http.Request) (interface{}, error)

func NewHandler(config *config.Configuration) *Handler {
	return &Handler{config: config}
}

func (h *Handler) JSONHandle(action Action) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		ctx, cancel := context.WithTimeout(r.Context(), time.Second*5)
		defer cancel()

		r = r.WithContext(ctx)

		responseData, err := action(ctx, r)
		if err != nil {
			h.handleJsonError(w, err)

			return
		}

		err = json.NewEncoder(w).Encode(responseData)
		if err != nil {
			http.Error(w, "Failed to encode JSON", http.StatusInternalServerError)
			return
		}
	}
}

func (h *Handler) handleJsonError(w http.ResponseWriter, err error) {
	var apiErrorResponse *ApiErrorResponse
	var businessError common_error.BusinessError
	statusCode := http.StatusInternalServerError

	// Check and handle business business_errors
	if errors.As(err, &businessError) {
		statusCode = http.StatusBadRequest

		code := businessError.Code()

		apiErrorResponse = &ApiErrorResponse{
			Error: businessError.Error(),
			Code:  &code,
		}
	} else {
		apiErrorResponse = &ApiErrorResponse{
			Error: err.Error(),
		}
	}

	w.WriteHeader(statusCode)

	if apiErrorResponse != nil {
		responseErrorJSON, _ := json.Marshal(apiErrorResponse)
		_, _ = w.Write(responseErrorJSON)
	} else {
		http.Error(w, err.Error(), statusCode)
	}
}
