package service

import (
	"fmt"
	"net/http"

	"customer-service/client"
	"customer-service/model"
	"customer-service/repository"

	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
)

func PlaceOrder(ctx echo.Context) error {
	var request model.PlaceOrderResquest
	err := ctx.Bind(&request)
	if err != nil {
		log.Error(err)
		return echo.ErrBadRequest
	}
	order, err := repository.CreateNewOrder(request)
	if err != nil {
		log.Error(err)
		data := map[string]interface{}{
			"message": err.Error(),
		}
		return ctx.JSON(http.StatusOK, data)
	}
	restaurant, err := repository.FindRestaurantById(request.RestaurantId)
	if err != nil {
		log.Error(err)
		data := map[string]interface{}{
			"message": err.Error(),
		}
		return ctx.JSON(http.StatusOK, data)
	}
	orderId := fmt.Sprintf("%d", order.Id)
	notificationRequest := model.NotificationRequest{
		Recipient: "restaurant",
		OrderId: orderId,
		Message: fmt.Sprintf("Order for restaurant %s is created with menus %v.", restaurant.Name, order.Items),
	}
	err = client.SendNotification(notificationRequest)
	if err != nil {
		log.Error(err)
		data := map[string]interface{}{
			"message": err.Error(),
		}
		return ctx.JSON(http.StatusOK, data)
	}
	return ctx.JSON(http.StatusOK, model.OrderStatusResponse{OrderId: orderId, Status: order.Status})
}