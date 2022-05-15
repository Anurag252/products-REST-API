package models

type ErrorResponse struct {
	Title       string
	Description string
	Code        string
}

const (
	//titles
	BadRequest = "Bad Request"

	//descriptions
	ByPriceLessThanNotANumber = "byPriceLessThan query parameter is not a number"

	//codes
	ER01 = "ER-01"
)
