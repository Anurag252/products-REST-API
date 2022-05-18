package controller_test

import (
	"strconv"
	"testing"

	controller "github.com/anurag252/products-rest-api/controllers"
	"github.com/anurag252/products-rest-api/models"
)

func TestDiscounts(t *testing.T) {

	test := []struct {
		Product       models.Product
		ExpectedPrice int
	}{
		{
			Product: models.Product{
				Sku:      "000003",
				Name:     "Test1",
				Category: "boots",
				Price:    2000,
			}, ExpectedPrice: 1400,
		},
		{
			Product: models.Product{
				Sku:      "000003",
				Name:     "Test1",
				Category: "Category2",
				Price:    2000,
			}, ExpectedPrice: 1700,
		},
		{
			Product: models.Product{
				Sku:      "0002",
				Name:     "Test1",
				Category: "Category3",
				Price:    2000,
			}, ExpectedPrice: 2000,
		},
	}

	discount := &controller.ProductDiscount{}

	for _, p := range test {
		gotProduct := discount.CalculateDiscount(&p.Product)

		if gotProduct.Price != p.ExpectedPrice {
			t.Errorf("CalculateDiscount test failed , got price %s, wanted price %s", strconv.Itoa(gotProduct.Price), strconv.Itoa(p.ExpectedPrice))
		}
	}

}
