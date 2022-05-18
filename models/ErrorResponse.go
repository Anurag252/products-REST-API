package models

type ErrorResponse struct {
	Title       string
	Description string
	Code        string
}

const (
	//titles
	BadRequest      = "Bad Request"
	TooManyRequests = "Too many Requests"

	//descriptions
	ByPriceLessThanNotANumber  = "priceLessThan query parameter is not a number"
	TooManyRequestsDescription = "You are sending too many requests too fast"

	//codes
	ER01 = "ER-01"
	ER02 = "ER-02"
)
