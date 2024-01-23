package middleware

import (
	"github.com/gin-gonic/gin"
	"short-url/internal/global/errs"
)

func Recovery() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer errs.Recovery(c)
		c.Next()
	}
}
