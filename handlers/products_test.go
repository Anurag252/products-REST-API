package products_test

import (
	"testing"

	products "github.com/anurag252/products-rest-api/handlers"
	"github.com/anurag252/products-rest-api/mocks"
)

func TestGetProductsByCategory(testing *testing.T) {

	context := mocks.MockContext{}
	logger := mocks.MockLogger{}
	database := mocks.Mockdatabase{}

	products.GetProductsByCategory(context, "", logger, database)
}
