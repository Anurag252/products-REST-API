package dataaccess_test

import (
	"strconv"
	"testing"

	dataccess "github.com/anurag252/products-rest-api/database"
	"github.com/anurag252/products-rest-api/mocks"
	"github.com/anurag252/products-rest-api/models"
)

func TestFilterByPriceAndCategory(testing *testing.T) {

	var database = dataccess.InMemoryDatabase{}
	database.ProductCollection = models.ProductCollection{
		Products: []models.Product{
			{"sku1", "Name1", "Category1", 5000},
			{"sku2", "Name2", "Category2", 6000},
			{"sku3", "Name3", "Category3", 7000},
			{"sku4", "Name4", "Category4", 8000},
			{"sku5", "Name5", "Category4", 9000},
			{"sku6", "Name6", "Category4", 10000},
			{"sku7", "Name7", "Category5", 11000},
		},
	}

	fakeLogger := mocks.MockLogger{}

	//only category filter
	resultset1 := models.ProductCollection{
		Products: []models.Product{
			{"sku4", "Name4", "Category4", 8000},
			{"sku5", "Name5", "Category4", 9000},
			{"sku6", "Name6", "Category4", 10000},
		},
	}

	//only price filter <= 8000
	resultset2 := models.ProductCollection{
		Products: []models.Product{
			models.Product{"sku1", "Name1", "Category1", 5000},
			models.Product{"sku2", "Name2", "Category2", 6000},
			models.Product{"sku3", "Name3", "Category3", 7000},
			models.Product{"sku4", "Name4", "Category4", 8000},
		},
	}

	//category 4 and less than 9000
	resultset3 := models.ProductCollection{
		Products: []models.Product{
			models.Product{"sku4", "Name4", "Category4", 8000},
			models.Product{"sku5", "Name5", "Category4", 9000},
		},
	}

	tests := []struct {
		FilterByPrice    bool
		FilterByCategory bool
		Price            int
		Category         string
		Logger           mocks.MockLogger
		Expected         models.ProductCollection
	}{
		{false, true, 0, "Category4", fakeLogger, resultset1},
		{true, false, 8000, "", fakeLogger, resultset2},
		{true, true, 9000, "Category4", fakeLogger, resultset3},
	}

	for _, tt := range tests {
		got := database.FilterByPriceAndCategory(tt.FilterByPrice, tt.FilterByCategory, tt.Price, tt.Category, tt.Logger)
		if len(got.Products) != len(tt.Expected.Products) {
			testing.Errorf("FilterByPriceAndCategory failed got %s want %s", strconv.Itoa(len(got.Products)), strconv.Itoa(len(tt.Expected.Products)))
		}
	}

}
