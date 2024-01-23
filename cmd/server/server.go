package server

import (
	"github.com/gin-gonic/gin"
	"log"
	"short-url/config"
	"short-url/internal/global/database"
	"short-url/internal/global/middleware"
	"short-url/internal/module"
)

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
	for _, m := range module.Modules {
		log.Println("InitRouter: " + m.GetName())
		m.InitRouter(r.Group("/" + config.Get().Prefix))
	}
	err := r.Run(config.Get().Host + ":" + config.Get().Port)
	if err != nil {
		panic(err)
	}
}
