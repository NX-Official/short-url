package middleware

import (
	"github.com/gin-gonic/gin"
	"short-url/internal/global/errs"
	"short-url/internal/global/jwt"
)

func Auth() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("token")
		if payload, valid := jwt.ParseToken(token); !valid {
			errs.Fail(c, errs.ErrTokenInvalid)
			c.Abort()
			return
		} else {
			c.Set("payload", payload)
		}
		c.Next()
	}
}
