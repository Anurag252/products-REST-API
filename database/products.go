package dataaccess

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/anurag252/products-rest-api/logger"
	"github.com/anurag252/products-rest-api/models"
)

const (
	MaxresultSize = 5
)

type Database interface {
	FilterByPriceAndCategory(bool, bool, int, string, logger.Logger) models.ProductCollection
}

type InMemoryDatabase struct {
	ProductCollection models.ProductCollection
}

func New() Database {
	inMemoryDatabase := InMemoryDatabase{}
	data, err := os.ReadFile("products.json")
	if err == nil {
		err = json.Unmarshal(data, &inMemoryDatabase.ProductCollection)
		if err != nil {

			fmt.Printf("error while unmarshalling data %s", err)
		}
	} else {
		fmt.Printf("error while reading file %s", err)
	}
	return inMemoryDatabase
}

func (database InMemoryDatabase) FilterByPriceAndCategory(filterByPrice bool, filterByCategory bool, price int, category string, logger logger.Logger) models.ProductCollection {

	var productsByPrice []models.Product
	var resultLength int = 0
	var productsByCategory []models.Product

	if filterByCategory {
		for _, product := range database.ProductCollection.Products {
			if product.Category == category {
				productsByCategory = append(productsByCategory, product)
			}
		}
	} else {
		productsByCategory = database.ProductCollection.Products
	}

	if filterByPrice {
		for _, product := range productsByCategory {
			if product.Price <= price {
				productsByPrice = append(productsByPrice, product)
			}
		}
	} else {
		productsByPrice = productsByCategory
	}

	if MaxresultSize > len(productsByPrice) {
		resultLength = len(productsByPrice)
	} else {
		resultLength = MaxresultSize
	}

	filteredProductCollection := &models.ProductCollection{
		Products: productsByPrice[0:resultLength],
	}
	return *filteredProductCollection
}
