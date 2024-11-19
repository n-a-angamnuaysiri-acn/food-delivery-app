package restaurant

import "github.com/n-a-angamnuaysiri-acn/food-delivery-app/common"

type Restaurant struct {
	common.BaseData
	Menu []Menu `gorm:"foreignKey:Id" json:"menu"`
}

type Menu struct {
	common.BaseData
	Price       float64 `json:"price"`
	Description string  `json:"description"`
}

type GetRestaurantsResponse struct {
	Restaurants []common.BaseData `json:"restaurant"`
}

func (resp *GetRestaurantsResponse) AddRestaurants(restaurants []*Restaurant) {
	for _, r := range restaurants {
		resp.Restaurants = append(resp.Restaurants, r.BaseData)
	}
}