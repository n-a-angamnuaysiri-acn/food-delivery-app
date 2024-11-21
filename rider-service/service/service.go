package service

import (
	"net/http"

	"rider-service/model"
	"rider-service/repository"

	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
)

func GetRiders(ctx echo.Context) error {
	log.Info("Get All Riders Data")

	riders, err := repository.FindAllRider()

	if err != nil {
		data := map[string]interface{}{
			"message": err.Error(),
		}
		return ctx.JSON(http.StatusOK, data)
	}
	var response model.GetRidersResponse
	response.AddRiders(riders)
	return ctx.JSON(http.StatusOK, response)
}
