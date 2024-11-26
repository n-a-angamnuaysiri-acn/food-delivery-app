package main

import (
	"context"
	"customer-service/config"
	"customer-service/client"
	"customer-service/service"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/labstack/gommon/log"
)

func main() {
	// Connect To Database
	go func() {
		config.DatabaseInit()
		gorm := config.Database()

		dbGorm, err := gorm.DB()
		if err != nil {
			panic(err)
		}
		err = dbGorm.Ping()
		if err != nil {
			panic(err)
		}
	}()

	// Set Up Apis
	app := echo.New()
	app.Use(middleware.Recover())
	app.Use(middleware.Logger())

	app.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})
	app.POST("/order", service.PlaceOrder)

	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt)
	defer stop()

	// Start server
	go func() {
		if err := app.Start(":8086"); err != nil && err != http.ErrServerClosed {
			log.Fatal("shutting down the server")
		}
	}()

	// start listening for Notifications
	doneListening := make(chan bool)
	go client.ListenForNotification(doneListening)

	// Wait for interrupt signal to gracefully shut down the server with a timeout of 10 seconds.
	<-ctx.Done()
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := app.Shutdown(ctx); err != nil {
		log.Fatal(err)
	}
}
