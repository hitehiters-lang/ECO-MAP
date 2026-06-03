package models

type LoginBody struct {
	Login    string `json:"login"`
	Password string `json:"password"`
}
