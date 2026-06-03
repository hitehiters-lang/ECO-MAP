package models

type Category struct {
	ID              int    `json:"id"`
	Name            string `json:"name"`
	IsContainerYard bool   `json:"is_container_yard"`
}
