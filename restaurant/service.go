package restaurant

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
	"github.com/n-a-angamnuaysiri-acn/food-delivery-app/config"
)

func GetRestaurants(ctx echo.Context) error {
	log.Info("Get All Restaurants Data")
	db := config.Database()

	var restaurants []*Restaurant 
	dbResponse := db.Find(&restaurants)

	if dbResponse.Error != nil {
		data := map[string]interface{}{
			"message": dbResponse.Error.Error(),
		}
		return ctx.JSON(http.StatusOK, data)
	}
	
	var response  GetRestaurantsResponse
	response.AddRestaurants(restaurants)
	return ctx.JSON(http.StatusOK, response)
}