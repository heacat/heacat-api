package main

import (
	"github.com/heacat/heacat-api/src/api"
	"github.com/heacat/heacat-api/src/config"
	"github.com/heacat/heacat-api/src/logger"
)

func main() {
	// Init logger
	logger.InitLogger()

	// Init config
	config.InitConfig()

	// Init API
	api.InitAPI()
}
