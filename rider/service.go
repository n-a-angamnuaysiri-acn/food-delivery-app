package rider

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
)

func GetRiders(ctx echo.Context) error {
	log.Info("Get All Riders Data")

	riders, err := FindAll()

	if err != nil {
		data := map[string]interface{}{
			"message": err.Error(),
		}
		return ctx.JSON(http.StatusOK, data)
	}
	var response GetRidersResponse
	response.AddRiders(riders)
	return ctx.JSON(http.StatusOK, response)
}