package main

import (
	"github.com/pedropassos06/gopportunities/config"
	"github.com/pedropassos06/gopportunities/router"
)

var (
	logger *config.Logger
)

func main() {
	// Initialize Logger
	logger = config.GetLogger("main")

	// Initialize configs
	err := config.Init()
	if err != nil {
		logger.Err(err)
		return
	}

	// Initialize Router
	router.Initialize()
}
