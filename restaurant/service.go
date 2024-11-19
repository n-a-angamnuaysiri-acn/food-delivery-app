package restaurant

import (
	"net/http"
	"strconv"

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

	var response GetRestaurantsResponse
	response.AddRestaurants(*restaurants)
	return ctx.JSON(http.StatusOK, response)
}

func GetMenu(ctx echo.Context) error {
	restaurantId := ctx.QueryParam("restaurant_id")
	if len(restaurantId) == 0 {
		data := map[string]interface{}{
			"message": "Not receive restaurant id",
		}
		return ctx.JSON(http.StatusOK, data)
	}
	idInt, err := strconv.Atoi(restaurantId)
	if err != nil {
		log.Error(err)
		return echo.ErrBadRequest
	}
	restaurant, err := FindById(uint(idInt))
	if err != nil {
		data := map[string]interface{}{
			"message": err.Error(),
		}
		return ctx.JSON(http.StatusOK, data)
	}

	response := GetMenuResponse{RestaurantId: restaurant.Id, Menu: restaurant.Menu}
	return ctx.JSON(http.StatusOK, response)
}
