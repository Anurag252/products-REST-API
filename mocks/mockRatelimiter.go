package mocks

import "github.com/gin-gonic/gin"

type MockRateLimited struct {
}

func (mockRateLimit MockRateLimited) RateLimit() gin.HandlerFunc {
	return nil
}
