package controller

import "github.com/anurag252/products-rest-api/models"

type Discount interface {
	CalculateDiscount(*models.Product) models.Product
}

type ProductDiscount struct {
}

func (productDiscount ProductDiscount) CalculateDiscount(product *models.Product) models.Product {

	if product.Category == "boots" && product.Sku == "000003" {
		product.Price = (product.Price * 70) / 100
		return *product
	}

	if product.Category == "boots" {
		product.Price = (product.Price * 70) / 100
		return *product
	}

	if product.Sku == "000003" {
		product.Price = (product.Price * 85) / 100
		return *product
	}

	return *product

}
