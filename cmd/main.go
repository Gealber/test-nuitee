package main

import (
	"log"

	"github.com/Gealber/nuitee/config"
	"github.com/Gealber/nuitee/router"
)

func main() {
	// load config environment variables
	cfg := config.Config()

	// initialize router and all its routes
	r, err := router.Setup(cfg)
	if err != nil {
		log.Fatal(err)
	}

	// run server
	r.Run()
}
