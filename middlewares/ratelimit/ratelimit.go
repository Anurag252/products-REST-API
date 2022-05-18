package ratelimit

import (
	"fmt"

	"github.com/anurag252/products-rest-api/models"
	"github.com/gin-gonic/gin"
	"golang.org/x/time/rate"
)

var (
	//number of requests per second
	Fillrate = 0.00001157407 //1 token per 24 hrs

	//maximum bucket size
	BurstRate = 20
)

type RateLimiter interface {
	RateLimit() gin.HandlerFunc
}

type TokenRateLimiter struct {
	ratelimit rate.Limiter
}

func (rateLimit TokenRateLimiter) RateLimit() gin.HandlerFunc {
	return func(c *gin.Context) {
		if !rateLimit.ratelimit.Allow() {
			errorResponse := models.ErrorResponse{
				Title:       models.TooManyRequests,
				Description: models.TooManyRequestsDescription,
				Code:        models.ER02,
			}
			fmt.Printf(errorResponse.Code)
			c.AbortWithStatusJSON(429, errorResponse)
		}
	}
}

func New() RateLimiter {
	ratelimiter := TokenRateLimiter{
		ratelimit: *rate.NewLimiter(rate.Limit(Fillrate), BurstRate),
	}
	return ratelimiter
}
