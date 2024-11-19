package rider

import "github.com/n-a-angamnuaysiri-acn/food-delivery-app/config"

func FindAll() ([]*Rider, error) {
	var db = config.Database()
	var riders []*Rider
	dbResponse := db.Find(&riders)

	if dbResponse.Error != nil {
		return nil, dbResponse.Error
	}
	return riders, nil
}