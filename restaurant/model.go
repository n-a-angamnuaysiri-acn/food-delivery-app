package restaurant

type BaseData struct {
	Id   uint   `gorm:"primaryKey" json:"id"`
	Name string `json:"name"`
}

type Restaurant struct {
	BaseData
	Menu []Menu `json:"menu"`
}

var Restaurants = []Restaurant{

}

type Menu struct {
	BaseData
	Price       float64 `json:"price"`
	Description string  `json:"description"`
}

type GetRestaurantsResponse struct {
	Restaurants []BaseData `json:"restaurants"`
}