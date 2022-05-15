package routedef

import (
	"strconv"

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

func (routes Routes) RegisterRoutes(router *gin.Engine, logger logger.Logger, database dataaccess.Database) {
	router.GET("/products", func(ctx *gin.Context) {

		context := products.GinContext{
			GinProvidedContext: ctx,
		}

		byCategory := ctx.Query("category")

		byPriceLessThan := ctx.Query("ByPriceLessThan")

		if byCategory != emptyString {
			products.GetProductsByCategory(context, byCategory, logger, database)
			return
		}

		if byPriceLessThan != emptyString {
			price, err := strconv.Atoi(byPriceLessThan)
			if err == nil {
				products.GetProductsByPriceLessThan(context, price, logger, database)
			} else {
				logger.LogError("byPriceLessThan query parameter is not an integer")

				response := models.ErrorResponse{
					Title:       models.BadRequest,
					Description: models.ByPriceLessThanNotANumber,
					Code:        models.ER01,
				}
				ctx.JSON(400, response)
			}
			return
		}

		products.GetProducts(context, logger, database)
	})
}
