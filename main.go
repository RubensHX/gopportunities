package main

import (
	"github.com/RubensHX/gopportunities/config"
	"github.com/RubensHX/gopportunities/router"
)

var (
	logger config.Logger
)

func main() {
	logger = *config.GetLogger("main")
	err := config.Init()
	if err != nil {
		logger.ErrorF("Error initializing config: %v", err)
		return
	}
	router.Initialize()
}
