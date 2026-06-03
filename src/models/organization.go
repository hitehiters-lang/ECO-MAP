package models

type Organization struct {
	ID         int    `json:"id"`
	Name       string `json:"name"`
	CategoryID int    `json:"category_id"`
	CityID     int    `json:"city_id"`
}
