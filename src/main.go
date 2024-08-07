package main

import (
	"github.com/heacat/heacat-api/src/api"
	"github.com/heacat/heacat-api/src/config"
	"github.com/heacat/heacat-api/src/logger"
	"github.com/heacat/heacat-api/src/notifier"
)

func main() {
	// Init logger
	logger.InitLogger()

	// Init config
	config.InitConfig()

	// Init telegram notifier
	notifier.InitTelegramNotifier()

	// Init API
	api.InitAPI()

}
