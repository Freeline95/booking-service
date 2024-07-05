package app

import (
	"booking-service/internal/router/middleware"
	internalhttp "booking-service/internal/server/http"
	"github.com/gorilla/mux"
)

func AddRoutes(router *mux.Router, handler *internalhttp.Handler, serviceProvider *ServiceProvider) {
	router.Use(middleware.LogMiddleware)

	AddOrderRoutes(router, handler, serviceProvider)
}

func AddOrderRoutes(router *mux.Router, handler *internalhttp.Handler, serviceProvider *ServiceProvider) {
	ozonRoutes := router.PathPrefix("/orders").Subrouter()

	ozonRoutes.HandleFunc("", handler.JSONHandle(serviceProvider.OrderController().CreateOrder)).Methods("POST")
}
