package main

import (
	"time"

	"github.com/heacat/heacat-api/src/api"
	"github.com/heacat/heacat-api/src/config"
	"github.com/heacat/heacat-api/src/control"
	"github.com/heacat/heacat-api/src/logger"
	"github.com/heacat/heacat-api/src/notifier"
	"github.com/heacat/heacat-api/src/utils"
)

func main() {
	// Init logger
	logger.InitLogger()

	// Init config
	config.InitConfig()

	// Init telegram notifier
	notifier.InitTelegramNotifier()

	// Start checks
	go utils.SetInterval(time.Duration(config.Config.Disk.CheckInterval)*time.Minute, control.CheckDiskStatus)
	go utils.SetInterval(time.Duration(config.Config.Cpu.CheckInterval)*time.Minute, control.CheckCPUStatus)
	go utils.SetInterval(time.Duration(config.Config.Memory.CheckInterval)*time.Minute, control.CheckMemoryStatus)

	// Init API
	api.InitAPI()

}
