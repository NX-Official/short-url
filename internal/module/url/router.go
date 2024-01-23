package url

import (
	"github.com/gin-gonic/gin"
)

func (u *ModuleUrl) InitRouter(r *gin.RouterGroup) {
	r.POST("/shorten", Shorten)
	r.GET("/:shortUrl", Jump)
}
