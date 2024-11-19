package restaurant

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
)

func GetRestaurants(ctx echo.Context) error {
	log.Info("Get All Restaurants Data")
	restaurants, err := FindAll()

	if err != nil {
		data := map[string]interface{}{
			"message": err.Error(),
		}
		return ctx.JSON(http.StatusOK, data)
	}
	
	var response  GetRestaurantsResponse
	response.AddRestaurants(restaurants)
	return ctx.JSON(http.StatusOK, response)
}