package main

import (
	"flag"
	"log"

	"github.com/BurntSushi/toml"
	"github.com/sudobooo/gopher_messenger/user_service/internal/app/user"
	"github.com/sudobooo/gopher_messenger/user_service/internal/config"
)

var (
	configPath string
)

func init() {
	flag.StringVar(&configPath, "config-path", "configs/user_service.toml", "path to config file")
}

func main() {
	flag.Parse()

	currentConfig := config.NewConfig()

	_, err := toml.DecodeFile(configPath, currentConfig)
	if err != nil {
		log.Fatal(err)
	}

	s := user.New(currentConfig)
	if err := s.Start(); err != nil {
		log.Fatal(err)
	}
}
