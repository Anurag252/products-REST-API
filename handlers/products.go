package products

import (
	"fmt"

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

func GetProductsByCategory(ctx Context, category string, logger logger.Logger, dataaccess dataaccess.Database) {
	logger.LogMessage(fmt.Sprint("Fetching products by category %s", category))
	productCollection := dataaccess.FilterByCategory(category, logger)
	ctx.JSON(200, productCollection)
}

func GetProductsByPriceLessThan(ctx Context, price int, logger logger.Logger, dataaccess dataaccess.Database) {
	logger.LogMessage(fmt.Sprint("Fetching all those products where price is less than %s", price))
	productCollection := dataaccess.FilterByPrice(price, logger)
	ctx.JSON(200, productCollection)
}

func GetProducts(ctx Context, logger logger.Logger, dataaccess dataaccess.Database) {
	logger.LogMessage("Fetching all products")
	productCollection := dataaccess.GetProducts(logger)
	ctx.JSON(200, productCollection)
}
