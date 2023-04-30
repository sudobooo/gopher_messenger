package main

import (
	"log"

	"github.com/sudobooo/gopher_messenger/user_service/internal/service"
)

func main() {
	s := service.New()
	if err := s.Start(); err != nil {
		log.Fatal(err)
	}
}
