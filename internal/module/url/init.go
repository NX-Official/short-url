package url

import (
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
	baseURL = "http://" + config.Get().Host + ":" + config.Get().Port + config.Get().Prefix + "api/"
}
