package common

type BaseData struct {
	Id   uint   `gorm:"primaryKey" json:"id"`
	Name string `json:"name"`
}