package models

type AddSuperuserBody struct {
	Login       string `json:"login"`
	Password    string `json:"password"`
	CityID      int    `json:"city_id" example:"1"`
	Name        string `json:"name,omitempty"`
	IsSuperuser bool   `json:"is_superuser"`
}
