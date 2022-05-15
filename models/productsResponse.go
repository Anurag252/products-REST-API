package models

type ProductResponse struct {
	Sku      string `json:"sku"`
	Name     string `json:"name"`
	Category string `json:"category"`
	Price    struct {
		Original           int    `json:"original"`
		Final              int    `json:"final"`
		DiscountPercentage string `json:"discount_percentage"`
		Currency           string `json:"currency"`
	} `json:"price"`
}
