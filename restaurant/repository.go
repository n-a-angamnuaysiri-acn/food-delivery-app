package restaurant

import (
	"encoding/json"

	"github.com/n-a-angamnuaysiri-acn/food-delivery-app/config"
)

func FindAll() ([]Restaurant, error) {
	var db = config.Database()
	var restaurantsDB *[]RestaurantDB
	dbResponse := db.Find(&restaurantsDB)

	if dbResponse.Error != nil {
		return nil, dbResponse.Error
	}
	var restaurants []Restaurant
	for _, r := range *restaurantsDB {
		var menu []Menu
		err := json.Unmarshal([]byte(r.Menu), &menu)
		if err != nil {
			return nil, err
		}
		restaurants = append(restaurants, Restaurant{BaseData: r.BaseData, Menu: menu})
	}
	return restaurants, nil
}

func FindById(id uint) (*Restaurant, error) {
	var db = config.Database()
	var restaurant *RestaurantDB
	dbResponse := db.Find(&restaurant, "id = ?", id)
	if dbResponse.Error != nil {
		return nil, dbResponse.Error
	}
	var menu []Menu
	err := json.Unmarshal([]byte(restaurant.Menu), &menu)
	if err != nil {
		return nil, err
	}
	return &Restaurant{BaseData: restaurant.BaseData, Menu: menu}, nil
}
