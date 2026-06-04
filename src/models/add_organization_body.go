package models

type AddOrganizationBody struct {
	Name        string  `json:"name"`
	CategoryID  int     `json:"category_id"`
	CityID      int     `json:"city_id" example:"1"`
	Description string  `json:"description"`
	CoordX      float64 `json:"coord_x" example:"47.11638889"`
	CoordY      float64 `json:"coord_y" example:"51.93421389"`
}
