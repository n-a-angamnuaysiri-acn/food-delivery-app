package rider

import "github.com/n-a-angamnuaysiri-acn/food-delivery-app/common"

type Rider struct {
	common.BaseData
}

type GetRidersResponse struct {
	Riders []Rider `json:"rider"`
}

func (resp *GetRidersResponse) AddRiders(riders []*Rider) {
	for _, r := range riders {
		resp.Riders = append(resp.Riders, *r)
	}
}