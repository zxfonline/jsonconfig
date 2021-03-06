package main

import (
	"github.com/zxfonline/jsonconfig"
)

func main() {
	// setup the structure
	config := struct {
		Url string `json:"url"`

		Methods []string `json:"methods"`

		AlwaysLoad bool `json:"always_load"`

		Module struct {
			Name  string `json:"name"`
			Route string `json:"route"`
			Port  int    `json:"port"`
		} `json:"module"`
	}{}

	// parse and load json config
	err := jsonconfig.Load("config.json", &config)

	if err == nil {
		println("The url is", config.Url)
		println("Supported methods are", config.Methods[0], config.Methods[1])

		println("The module is", config.Module.Name, "on route", config.Url, ":", config.Module.Port, config.Module.Route)
	} else {
		println(err.Error())
	}
}
