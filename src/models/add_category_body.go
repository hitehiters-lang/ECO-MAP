package models

type AddCategoryBody struct {
	Name            string `json:"name"`
	IsContainerYard bool   `json:"is_container_yard" example:"false"`
}
