package server

import (
	controller "github.com/anurag252/products-rest-api/controllers"
	dataaccess "github.com/anurag252/products-rest-api/database"
	"github.com/anurag252/products-rest-api/docs"
	"github.com/anurag252/products-rest-api/logger"
	"github.com/anurag252/products-rest-api/middlewares/ratelimit"
	routedef "github.com/anurag252/products-rest-api/routes"
	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

type DefaultServer struct {
	Gindefault gin.Engine
}

type Server interface {
	Serve(string)
}

func New(router routedef.Routes, database dataaccess.Database, ratelimit ratelimit.RateLimiter, logger logger.Logger) Server {
	ginDefault := gin.Default()
	docs.SwaggerInfo.BasePath = "/"
	ginDefault.Use(ratelimit.RateLimit())
	discounts := &controller.ProductDiscount{}
	router.RegisterRoutes(ginDefault, logger, database, discounts)
	server := DefaultServer{
		Gindefault: *ginDefault,
	}
	ginDefault.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	return server
}

func (server DefaultServer) Serve(address string) {
	server.Gindefault.Run(address)
}
