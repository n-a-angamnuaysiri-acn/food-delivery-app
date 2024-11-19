package restaurant

import "github.com/n-a-angamnuaysiri-acn/food-delivery-app/config"

var db = config.Database()

func FindAll() ([]*Restaurant, error) {
	var restaurants []*Restaurant 
	dbResponse := db.Find(&restaurants)

	if dbResponse.Error != nil {
		return nil, dbResponse.Error
	}
	return restaurants, nil
}