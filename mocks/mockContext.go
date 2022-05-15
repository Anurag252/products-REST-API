package mocks

import "github.com/gin-gonic/gin"

type MockContext struct {
	*gin.Context
}

func (mockContext MockContext) JSON(statuscode int, anyValue interface{}) {

}
