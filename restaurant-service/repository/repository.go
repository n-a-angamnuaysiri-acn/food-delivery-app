package repository

import (
	"encoding/json"

	"restaurant-service/config"
	"restaurant-service/model"

	"github.com/labstack/gommon/log"
)

func FindAllRestaurant() (*[]model.Restaurant, error) {
	var db = config.Database()
	var restaurantsDB *[]model.RestaurantDB
	dbResponse := db.Find(&restaurantsDB)

	if dbResponse.Error != nil {
		return nil, dbResponse.Error
	}
	var restaurants []model.Restaurant
	for _, r := range *restaurantsDB {
		var menu []model.Menu
		err := json.Unmarshal([]byte(r.Menu), &menu)
		if err != nil {
			return nil, err
		}
		restaurants = append(restaurants, model.Restaurant{BaseData: r.BaseData, Menu: menu})
	}
	return &restaurants, nil
}

func FindRestaurantById(id uint) (*model.Restaurant, error) {
	var db = config.Database()
	var restaurant *model.RestaurantDB
	dbResponse := db.Find(&restaurant, "id = ?", id)
	if dbResponse.Error != nil {
		return nil, dbResponse.Error
	}
	var menu []model.Menu
	err := json.Unmarshal([]byte(restaurant.Menu), &menu)
	if err != nil {
		log.Error(err)
		return nil, err
	}
	return &model.Restaurant{BaseData: restaurant.BaseData, Menu: menu}, nil
}
