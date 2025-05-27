package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/labstack/gommon/log"
	"github.com/one-project-one-month/Hotel-Booking-Management-System-Go/config"
	"github.com/one-project-one-month/Hotel-Booking-Management-System-Go/pkg/requestValidator"
)

type App struct {
	echo   *echo.Echo
	Logger echo.Logger
	wg     *sync.WaitGroup
	cfg    *config.Server
}

func NewApp(wg *sync.WaitGroup, cfg *config.Config) *App {
	app := echo.New()
	app.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
	}))
	app.Validator = &requestValidator.CustomValidator{Validator: validator.New()}

	if cfg.Environment == "development" {
		app.Logger.SetLevel(log.DEBUG)
		app.Logger.SetOutput(os.Stdout)
		app.Logger.SetHeader("${time_rfc3339} ${level} ${short_file}:${line} ${prefix} -")

	}

	return &App{
		echo:   app,
		wg:     wg,
		cfg:    &cfg.Server,
		Logger: app.Logger,
	}
}

func (app *App) start() {
	app.wg.Add(1)
	go func() {
		defer app.wg.Done()
		app.echo.Logger.Info("Server started on port: ", app.cfg.Port)
		app.Logger.Error(app.echo.Start(fmt.Sprintf("%s:%d", app.cfg.Host, app.cfg.Port)))
	}()

	app.shutdown()
}

func (app *App) shutdown() {
	signals := make(chan os.Signal, 1)
	signal.Notify(signals, os.Interrupt, syscall.SIGTERM)

	app.wg.Add(1)
	go func() {
		<-signals

		app.Logger.Info("received an interrupt, shutting down the server...")
		ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)

		if err := app.echo.Shutdown(ctx); err != nil {
			app.Logger.Fatal(err)
		}

		defer func() {
			app.Logger.Info("server gracefully stopped")
			cancel()
			app.wg.Done()
		}()
	}()
}
