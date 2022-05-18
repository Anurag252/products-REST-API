package products_test

import (
	"strconv"
	"testing"

	controller "github.com/anurag252/products-rest-api/controllers"
	products "github.com/anurag252/products-rest-api/handlers"
	"github.com/anurag252/products-rest-api/mocks"
)

func TestGetProductsByCategory(testing *testing.T) {

	context := &mocks.MockContext{}
	logger := mocks.MockLogger{}
	database := mocks.Mockdatabase{}
	discount := &controller.ProductDiscount{}

	products.GetProductsByCategoryAndPrice(context, "", true, true, 7000, logger, database, discount)

	if context.StatusCode != 200 {
		testing.Error("TestGetProductsByCategory failed" + strconv.Itoa(context.StatusCode))
	}

}
