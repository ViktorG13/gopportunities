package main

import (
	"github.com/ViktorG13/gopportunities/config"
	"github.com/ViktorG13/gopportunities/router"
)

var (
	logger config.Logger
)

func main() {
	// Start Logger
	logger = *config.GetLogger("Main")
	//Initialize Config.
	err := config.Init()
	if err != nil {
		logger.Errorf("Error in Initialization: %s", err.Error())
		return
	}

	logger.Info("Server Started!")
	//Initialize Router.
	router.Initialize()
}
