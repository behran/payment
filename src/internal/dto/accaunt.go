package dto

//Account ...
type Account struct {
	Name string `json:"name" validate:"required"`
}
