package repository

import (
	"encoding/json"

	"customer-service/config"
	"customer-service/model"

	"gorm.io/gorm"
)

func CreateNewOrder(request model.PlaceOrderResquest) (*model.Order, error) {
	var db = config.Database()

	itemsData, err := json.Marshal(request.Items)
	if err != nil {
		return &model.Order{}, err

	}
	order := model.Order{
		RestaurantId: request.RestaurantId, 
		Items: string(itemsData),
		Status: "created",
	}
	err = db.Transaction(func(tx *gorm.DB) error {
        tx.Create(&order)
        return nil
    })
	if err != nil {
		return &model.Order{}, err
	}
	return &order, nil
}


func FindRestaurantById(id string) (*model.Restaurant, error) {
	var db = config.Database()
	var restaurant *model.RestaurantDB
	dbResponse := db.Find(&restaurant, "id = ?", id)
	if dbResponse.Error != nil {
		return nil, dbResponse.Error
	}
	var menu []model.Menu
	err := json.Unmarshal([]byte(restaurant.Menu), &menu)
	if err != nil {
		return nil, err
	}
	return &model.Restaurant{BaseData: restaurant.BaseData, Menu: menu}, nil
}

