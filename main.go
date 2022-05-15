package main

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
