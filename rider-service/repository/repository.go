package repository

import (
	"rider-service/config"
	"rider-service/model"
)

func FindAllRider() ([]*model.Rider, error) {
	var db = config.Database()
	var riders []*model.Rider
	dbResponse := db.Find(&riders)

	if dbResponse.Error != nil {
		return nil, dbResponse.Error
	}
	return riders, nil
}
