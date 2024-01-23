package url

import (
	"fmt"
	"short-url/config"
)

var (
	baseURL string
)

type ModuleUrl struct {
}

func (u *ModuleUrl) GetName() string {
	return "Url"
}

func (u *ModuleUrl) Init() {
	baseURL = config.Get().BaseURL
	fmt.Println("baseURL: ", baseURL)
}
