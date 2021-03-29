package models

//AccountAttributes ...
type AccountAttributes struct {
	ID      int     `json:"id" db:"account_id"`
	Name    string  `json:"name" db:"account_name"`
	Balance float64 `json:"balance" db:"balance"`
}
