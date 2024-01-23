package module

import (
	"github.com/gin-gonic/gin"
	"short-url/internal/module/user"
)

type Module interface {
	GetName() string
	Init()
	InitRouter(r *gin.RouterGroup)
}

var Modules []Module

func RegisterModule(m Module) {
	Modules = append(Modules, m)
}

func init() {
	// Register your module here
	RegisterModule(&user.ModuleUser{})
}
