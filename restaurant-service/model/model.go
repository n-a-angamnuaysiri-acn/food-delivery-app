package model

type BaseData struct {
	Id   uint   `gorm:"primaryKey" json:"id"`
	Name string `json:"name"`
}

type RestaurantDB struct {
	BaseData
	Menu string `gorm:"default:[]"`
}

func (RestaurantDB) TableName() string {
	return "app_data.restaurants"
}

type Restaurant struct {
	BaseData
	Menu []Menu `json:"menu"`
}

type Menu struct {
	BaseData
	Price       float64 `json:"price"`
	Description string  `json:"description"`
}

type GetRestaurantsResponse struct {
	Restaurants []BaseData `json:"restaurant"`
}

func (resp *GetRestaurantsResponse) AddRestaurants(restaurants []Restaurant) {
	for _, r := range restaurants {
		resp.Restaurants = append(resp.Restaurants, r.BaseData)
	}
}

type GetMenuResponse struct {
	RestaurantId uint   `json:"restaurant_id"`
	Menu         []Menu `json:"menu"`
}