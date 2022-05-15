package server

import (
	dataaccess "github.com/anurag252/products-rest-api/database"
	"github.com/anurag252/products-rest-api/logger"
	"github.com/anurag252/products-rest-api/middlewares/ratelimit"
	routedef "github.com/anurag252/products-rest-api/routes"
	"github.com/gin-gonic/gin"
)

type DefaultServer struct {
	Gindefault gin.Engine
}

type Server interface {
	Serve(string)
}

func New(router routedef.Routes, database dataaccess.Database, ratelimit ratelimit.RateLimiter, logger logger.Logger) Server {
	ginDefault := gin.Default()
	router.RegisterRoutes(ginDefault, logger, database)
	server := DefaultServer{
		Gindefault: *ginDefault,
	}
	ginDefault.Use(ratelimit.RateLimit())
	return server
}

func (server DefaultServer) Serve(address string) {
	server.Gindefault.Run(address)
}
