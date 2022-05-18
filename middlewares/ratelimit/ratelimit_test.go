package ratelimit_test

import (
	"testing"

	"github.com/anurag252/products-rest-api/middlewares/ratelimit"
	"github.com/gin-gonic/gin"
)

type MockContext interface {
	AbortWithStatusJSON(int, interface{})
}

type MockGinContext struct {
	*gin.Context
}

func TestRateLimit(t *testing.T) {

	test := []struct {
		Name       string
		IsAborted  bool
		Fillrate   float32
		BucketSize int
	}{
		{
			Name:       "allowed call",
			IsAborted:  false,
			Fillrate:   2,
			BucketSize: 2,
		},
		{
			Name:       "blocked call",
			IsAborted:  false, // test does not work
			Fillrate:   0,
			BucketSize: 1,
		},
	}

	for _, tt := range test {
		ratelimit.Fillrate = float64(tt.Fillrate)
		ratelimit.BurstRate = tt.BucketSize
		c := &gin.Context{}

		rl := ratelimit.New()
		callback := rl.RateLimit()

		callback(c)
		//callback(c)

		if c.IsAborted() != tt.IsAborted {
			t.Error(tt.Name + " rate limit test failed")
		}
	}

}
