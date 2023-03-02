package models

type Building struct {
	ID           int    `json:"id"`
	BuildingName string `json:"buildingName"`
	Status       string `json:"status"`
}
