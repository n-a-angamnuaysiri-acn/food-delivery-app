package model

type Order struct {
	Id           uint   `gorm:"primaryKey" json:"id"`
	RestaurantId string `json:"restaurant_id"`
	RiderId      string `json:"rider_id"`
	Items        string `gorm:"default:[]"`
	Status       string `json:"status"`
}

type PlaceOrderResquest struct {
	RestaurantId string `json:"restaurant_id"`
	Items        []Item `json:"items"`
}

type OrderStatusResponse struct {
	OrderId string `json:"order_id"`
	Status  string `json:"status"`
}

type Item struct {
	MenuId   string `json:"menu_id"`
	Quantity uint   `json:"quantity"`
}

type NotificationRequest struct {
	Recipient string `json:"recipient"`
	OrderId string `json:"order_id"`
	Message string `json:"message"`
}

type NotificationResponse struct {
	Status string `json:"status"`
}

type BaseData struct {
	Id   string   `gorm:"primaryKey" json:"id"`
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