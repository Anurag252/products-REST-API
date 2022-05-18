package mocks

import "github.com/gin-gonic/gin"

type MockContext struct {
	*gin.Context
	StatusCode int
	Value      interface{}
}

func (mockContext *MockContext) JSON(statuscode int, anyValue interface{}) {
	mockContext.StatusCode = statuscode
	mockContext.Value = anyValue
}
