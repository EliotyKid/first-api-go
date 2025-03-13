package main

import (
	"github.com/EliotyKid/first-api-go.git/config"
	"github.com/EliotyKid/first-api-go.git/router"
)

var (
	logger *config.Logger
)

func main() {
	logger = config.GetLogger("main")

	//Initialize configs
	err := config.Init()
	if err != nil {
		logger.Errorf("config initialization error: %v", err)
		return
	}

	//Initialize Routes
	router.Initialize()
}
