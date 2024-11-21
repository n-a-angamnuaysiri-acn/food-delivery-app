package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/labstack/gommon/log"
	"restaurant-service/config"
	"restaurant-service/service"
)

func main() {
	app := echo.New()
	app.Use(middleware.Recover())
	app.Use(middleware.Logger())

	app.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})
	app.GET("/restaurant", service.GetRestaurants)
	app.GET("/menu", service.GetMenu)
	app.POST("/restaurant/order/accept", service.AcceptingOrder)
	

	// Connect To Database
	config.DatabaseInit()
	gorm := config.Database()

	dbGorm, err := gorm.DB()
	if err != nil {
		panic(err)
	}

	dbGorm.Ping()

	log.Fatal(app.Start(":8082"))
}

