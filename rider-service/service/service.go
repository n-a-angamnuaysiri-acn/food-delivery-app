package service

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
	"rider-service/model"
	"rider-service/repository"
)

func GetRiders(ctx echo.Context) error {
	log.Info("Get All Riders Data")

	riders, err := repository.FindAll()

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