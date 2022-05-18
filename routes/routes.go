package routedef

import (
	"fmt"
	"strconv"

	controller "github.com/anurag252/products-rest-api/controllers"
	dataaccess "github.com/anurag252/products-rest-api/database"
	products "github.com/anurag252/products-rest-api/handlers"
	"github.com/anurag252/products-rest-api/logger"
	"github.com/anurag252/products-rest-api/models"
	"github.com/gin-gonic/gin"
)

type Routes struct {
}

const emptyString string = ""

func New() Routes {
	routes := Routes{}
	return routes
}

func (routes Routes) RegisterRoutes(router *gin.Engine, logger logger.Logger, database dataaccess.Database, discountService controller.Discount) {

	router.GET("/products", func(ctx *gin.Context) {
		context := products.GinContext{
			GinProvidedContext: ctx,
		}
		var filterByPrice, filterByCategory bool = false, false

		byCategory := ctx.Query("category")

		byPriceLessThan := ctx.Query("priceLessThan")

		if byCategory != emptyString {
			filterByCategory = true
		}

		if byPriceLessThan != emptyString {
			filterByPrice = true
		}

		price, err := strconv.Atoi(byPriceLessThan)
		// making priceLessThan as optional
		if err == nil || !filterByPrice {
			products.GetProductsByCategoryAndPrice(context, byCategory, filterByPrice, filterByCategory, price, logger, database, discountService)
		} else {
			logger.LogError(fmt.Sprintf("byPriceLessThan query parameter is not an integer:- %s", err))

			response := models.ErrorResponse{
				Title:       models.BadRequest,
				Description: models.ByPriceLessThanNotANumber,
				Code:        models.ER01,
			}
			ctx.JSON(400, response)
		}
		return
	})
}
