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
	GetProducts(logger.Logger) models.ProductCollection
	FilterByPrice(int, logger.Logger) models.ProductCollection
	FilterByCategory(string, logger.Logger) models.ProductCollection
}

type InMemoryDatabase struct {
	productCollection models.ProductCollection
}

func New() Database {
	inMemoryDatabase := InMemoryDatabase{}
	data, err := os.ReadFile("products.json")
	if err == nil {
		err = json.Unmarshal(data, &inMemoryDatabase.productCollection)
		if err != nil {

			fmt.Printf("error while unmarshalling data %s", err)
		}
	} else {
		fmt.Printf("error while reading file %s", err)
	}
	return inMemoryDatabase
}

func (database InMemoryDatabase) FilterByPrice(price int, logger logger.Logger) models.ProductCollection {

	var products []models.Product
	var resultLength int = 0

	for _, product := range database.productCollection.Products {
		if product.Price < price {
			products = append(products, product)
		}
	}
	if MaxresultSize > len(products) {
		resultLength = len(products)
	} else {
		resultLength = MaxresultSize
	}

	filteredProductCollection := &models.ProductCollection{
		Products: products[0:resultLength],
	}
	return *filteredProductCollection
}

func (database InMemoryDatabase) FilterByCategory(category string, logger logger.Logger) models.ProductCollection {
	var products []models.Product
	var resultLength int = 0
	for _, product := range database.productCollection.Products {
		if product.Category == category {
			products = append(products, product)
		}
	}

	if MaxresultSize > len(products) {
		resultLength = len(products)
	} else {
		resultLength = MaxresultSize
	}

	filteredProductCollection := &models.ProductCollection{
		Products: products[0:resultLength],
	}
	return *filteredProductCollection
}

func (database InMemoryDatabase) GetProducts(logger logger.Logger) models.ProductCollection {
	logger.LogMessage("Getting all products")
	var resultLength int = 0

	if MaxresultSize > len(database.productCollection.Products) {
		resultLength = len(database.productCollection.Products)
	} else {
		resultLength = MaxresultSize
	}

	result := models.ProductCollection{
		Products: database.productCollection.Products[0:resultLength],
	}
	return result
}
