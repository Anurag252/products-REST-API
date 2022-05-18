package products

import (
	"fmt"
	"strconv"

	controller "github.com/anurag252/products-rest-api/controllers"
	dataaccess "github.com/anurag252/products-rest-api/database"
	"github.com/anurag252/products-rest-api/logger"
	"github.com/gin-gonic/gin"
)

type Context interface {
	JSON(int, interface{})
}

type GinContext struct {
	GinProvidedContext *gin.Context
}

func (ctx GinContext) JSON(statusCode int, value interface{}) {
	ctx.GinProvidedContext.JSON(statusCode, value)
}

// @BasePath /
// Products API godoc
// @Summary Gets all products
// @Schemes http
// @Description Gets all products
// @Tags products
// @Accept json
// @Produce json
// @Success 200
// @Param priceLessThan query int false "price less than"
// @Param category query string false "category"
// @Router /products [get]
func GetProductsByCategoryAndPrice(ctx Context, category string, filterByPrice bool, filterByCategory bool, price int, logger logger.Logger, dataaccess dataaccess.Database, discountService controller.Discount) {
	logger.LogMessage(fmt.Sprintf("Fetching all those products where price is less than %s", strconv.Itoa(price)))
	productCollection := dataaccess.FilterByPriceAndCategory(filterByPrice, filterByCategory, price, category, logger)
	for _, p := range productCollection.Products {
		discountService.CalculateDiscount(&p)
	}
	ctx.JSON(200, productCollection)
}
