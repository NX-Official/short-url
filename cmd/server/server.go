package server

import (
	"github.com/gin-gonic/gin"
	"io/fs"
	"log"
	"net/http"
	"short-url/config"
	"short-url/internal/global/database"
	"short-url/internal/global/middleware"
	"short-url/internal/module"
	"strings"
)

var FrontFiles fs.FS

func Init() {
	config.Read()
	database.Init()

	for _, m := range module.Modules {
		log.Println("Init Module: " + m.GetName())
		m.Init()
	}
}

func Run() {
	r := gin.New()
	r.Use(gin.Logger(), middleware.Recovery())

	r.Use(func(c *gin.Context) {
		if !strings.HasPrefix(c.Request.URL.Path, "/api") {
			c.FileFromFS(c.Request.URL.Path, http.FS(FrontFiles))
		}
	})

	apiGroup := r.Group(config.Get().Prefix + "/api")
	for _, m := range module.Modules {
		log.Println("InitRouter: " + m.GetName())
		m.InitRouter(apiGroup)
	}

	err := r.Run(config.Get().Host + ":" + config.Get().Port)
	if err != nil {
		panic(err)
	}
}
