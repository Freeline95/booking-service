package middleware

import (
	"net/http"

	common_log "booking-service/pkg/log"
)

const LogTemplateMessage = "Log middleware. Incoming request: %v\n"

func LogMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		common_log.Info(LogTemplateMessage, *r)

		next.ServeHTTP(w, r)
	})
}
