package main

import (
	"gorm.io/gen"
	"short-url/internal/model"
)

func main() {
	g := gen.NewGenerator(gen.Config{
		OutPath: "./internal/global/query",
		Mode:    gen.WithoutContext, // generate mode
	})

	//g.ApplyBasic(database.Models...)
	g.ApplyBasic(
		model.Url{},
	)

	g.Execute()
}
