package model

type BaseData struct {
	Id   uint   `gorm:"primaryKey" json:"id"`
	Name string `json:"name"`
}

type Rider struct {
	BaseData
}

type GetRidersResponse struct {
	Riders []Rider `json:"rider"`
}

func (resp *GetRidersResponse) AddRiders(riders []*Rider) {
	for _, r := range riders {
		resp.Riders = append(resp.Riders, *r)
	}
}