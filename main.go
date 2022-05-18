package main

// ref: https://swaggo.github.io/swaggo.io/declarative_comments_format/general_api_info.html
// @title Products API
// @version 1.0
// @description This is a sample API that returns a list of Products.
// @termsOfService http://swagger.io/terms/

// @contact.name anurag Dubey
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:1234
// @BasePath /
// @schemes http

import (
	dataaccess "github.com/anurag252/products-rest-api/database"
	"github.com/anurag252/products-rest-api/logger"
	"github.com/anurag252/products-rest-api/middlewares/ratelimit"
	routedef "github.com/anurag252/products-rest-api/routes"
	"github.com/anurag252/products-rest-api/server"
)

func main() {
	router := routedef.New()

	database := dataaccess.New()

	ratelimit := ratelimit.New()

	logger := logger.New()

	server := server.New(router, database, ratelimit, logger)

	server.Serve(":1234")
}
