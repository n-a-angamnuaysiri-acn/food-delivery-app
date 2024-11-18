package restaurant

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
)

func GetRestaurants(ctx echo.Context) error {
	log.Info("Get All Restaurants Data")
	restaurants := Restaurants
	var response  GetRestaurantsResponse
	
	for _, r := range restaurants {
		response.Restaurants = append(response.Restaurants, r.BaseData)
	}
	return ctx.JSON(http.StatusOK, response)
}