package main

import (
	"flag"
	"log"

	"github.com/BurntSushi/toml"
	"github.com/paramonies/ozon-link-shortener/internal/app/apiserver"
)

var configPath string

func init() {
	flag.StringVar(&configPath, "config-path", "configs/apiserver.toml", "path to config file")
}

// @title Link Shorter Rest Service
// @version 1.0
// @description Cервис для сокращения ссылок

// @host localhost:8080
// @BasePath /
func main() {
	flag.Parse()
	config := apiserver.NewConfig()
	_, err := toml.DecodeFile(configPath, config)
	if err != nil {
		log.Fatal(err)
	}

	if err := apiserver.Start(config); err != nil {
		log.Fatal(err)
	}
}
