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
	if config.Get().Port == "80" {
		baseURL = "http://" + config.Get().Host + config.Get().Prefix + "api/"
	} else {
		baseURL = "http://" + config.Get().Host + ":" + config.Get().Port + config.Get().Prefix + "api/"
	}
}
