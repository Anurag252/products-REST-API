package mocks

import (
	"github.com/anurag252/products-rest-api/logger"
	"github.com/anurag252/products-rest-api/models"
)

type Mockdatabase struct {
}

func (mockdatabase Mockdatabase) FilterByPriceAndCategory(filterByCategory bool, filterByPrice bool, price int, category string, logger logger.Logger) models.ProductCollection {
	return models.ProductCollection{}
}
