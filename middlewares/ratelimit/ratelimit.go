package ratelimit

import "github.com/gin-gonic/gin"

type RateLimiter struct {
}

func (rateLimit RateLimiter) RateLimit() gin.HandlerFunc {
	return func(c *gin.Context) {

	}
}

func New() RateLimiter {
	ratelimiter := RateLimiter{}
	return ratelimiter
}
