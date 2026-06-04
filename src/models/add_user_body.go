package models

type AddUserBody struct {
	Login    string `json:"login"`
	Password string `json:"password"`
	CityID   int    `json:"city_id" example:"1"`
	Name     string `json:"name,omitempty"`
}
