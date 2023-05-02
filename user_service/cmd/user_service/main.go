package main

import (
	"flag"
	"log"

	"github.com/BurntSushi/toml"

	"github.com/sudobooo/gopher_messenger/user_service/internal/service"
)

var (
	configPath string
)

func init() {
	flag.StringVar(&configPath, "config-path", "configs/user_service.toml", "path to config file")
}

func main() {
	flag.Parse()

	config := service.NewConfig()

	_, err := toml.DecodeFile(configPath, config)
	if err != nil {
		log.Fatal(err)
	}

	s := service.New(config)
	if err := s.Start(); err != nil {
		log.Fatal(err)
	}
}
