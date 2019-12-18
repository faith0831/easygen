package app

import (
	"github.com/gin-gonic/gin"
)

// GIN GIN
type GIN struct {
	*gin.Context
}

// Success Success
func (c *GIN) Success(msg string, data interface{}) {
	c.JSON(200, gin.H{
		"code":   1,
		"msg":    msg,
		"result": data,
	})
}

// Ok Ok
func (c *GIN) Ok(data interface{}) {
	c.Success("ok", data)
}

// Error Error
func (c *GIN) Error(msg string) {
	c.JSON(200, gin.H{
		"code": 0,
		"msg":  msg,
	})
}
