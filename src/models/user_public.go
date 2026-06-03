package models

type UserPublic struct {
	ID          int    `json:"id"`
	Login       string `json:"login"`
	CityID      int    `json:"city_id"`
	IsSuperuser bool   `json:"is_superuser"`
}
