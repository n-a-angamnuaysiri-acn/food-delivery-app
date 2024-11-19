package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/labstack/gommon/log"
	"github.com/n-a-angamnuaysiri-acn/food-delivery-app/config"
	"github.com/n-a-angamnuaysiri-acn/food-delivery-app/restaurant"
	"github.com/n-a-angamnuaysiri-acn/food-delivery-app/rider"
)

func main() {
	app := echo.New()
	app.Use(middleware.Recover())
	app.Use(middleware.Logger())

	app.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})
	app.GET("/restaurant", restaurant.GetRestaurants)
	app.GET("/rider", rider.GetRiders)

	// Connect To Database
	config.DatabaseInit()
	gorm := config.Database()

	dbGorm, err := gorm.DB()
	if err != nil {
		panic(err)
	}

	dbGorm.Ping()

	log.Fatal(app.Start(":8080"))
}