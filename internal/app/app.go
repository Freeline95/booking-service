package app

import (
	"booking-service/internal/config"
	internal_http "booking-service/internal/server/http"
	common_error "booking-service/pkg/error"
	common_log "booking-service/pkg/log"
	"context"
	"fmt"
	"net/http"
)

type App struct {
	config          *config.Configuration
	serviceProvider *ServiceProvider
	HTTPServer      *http.Server
}

func NewApp(ctx context.Context) (*App, error) {
	app := &App{}

	err := app.init(ctx)
	if err != nil {
		return nil, common_error.Annotate(err, "Error while init dependencies")
	}

	return app, err
}

func (a *App) init(ctx context.Context) error {
	initFunctions := []struct {
		name string
		fn   func(context.Context) error
	}{
		{"Config", a.initConfig},
		{"ServiceProvider", a.initServiceProvider},
		{"HTTPServer", a.initHTTPServer},
	}

	for _, f := range initFunctions {
		err := f.fn(ctx)
		if err != nil {
			return common_error.Annotate(err, fmt.Sprintf("Error while initializing %s", f.name))
		}
	}

	return nil
}

func (a *App) Run() error {
	if err := a.runHTTPServer(); err != nil {
		return common_error.Annotate(err, "Error while running HTTP server")
	}

	return nil
}

func (a *App) Shutdown() {
	a.serviceProvider.Shutdown()

	if err := a.HTTPServer.Shutdown(nil); err != nil {
		common_log.Error("Error while shutdown http server: %s", err.Error())
	}

	common_log.Info("HTTP server stopped")
}

func (a *App) initServiceProvider(_ context.Context) error {
	a.serviceProvider = newServiceProvider(a.config)

	return nil
}

func (a *App) initConfig(_ context.Context) error {
	a.config = &config.Configuration{}

	a.config.HTTPServerConfiguration.Port = "8080"

	return nil
}

func (a *App) initHTTPServer(_ context.Context) error {
	handler := internal_http.NewHandler(a.config)

	AddRoutes(a.serviceProvider.Router(), handler, a.serviceProvider)

	a.HTTPServer = &http.Server{
		Addr:    ":" + a.config.HTTPServerConfiguration.Port,
		Handler: a.serviceProvider.Router(),
	}

	return nil
}

func (a *App) runHTTPServer() error {
	common_log.Info("HTTP server is running on port 80")

	if err := a.HTTPServer.ListenAndServe(); err != nil {
		common_log.Error("Error starting the server: %s\n", err)

		return err
	}

	return nil
}
