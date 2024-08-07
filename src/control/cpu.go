package control

import (
	"fmt"

	"github.com/heacat/heacat-api/src/config"
	"github.com/heacat/heacat-api/src/notifier"
	"github.com/heacat/heacat-api/src/utils"
)

func CheckCPUStatus() (any, string) {
	cpu := utils.GetCPUInfo()

	if cpu.LoadAvg >= float64(config.Config.Cpu.LoadLimit) {
		notifier.SendMessage(true, "cpu", fmt.Sprintf("[%s] [🔴 Alert] Anormal CPU status found (>%.2f%%)", config.Config.Alarm.ServerNickName, config.Config.Cpu.LoadLimit))
		return cpu, fmt.Sprintf("Anormal CPU status found (>%.2f%%)", config.Config.Cpu.LoadLimit)
	} else {
		notifier.SendMessage(false, "cpu", fmt.Sprintf("[%s] [🟢 Info] CPU status is normal (<%.2f%%)", config.Config.Alarm.ServerNickName, config.Config.Cpu.LoadLimit))
		return cpu, fmt.Sprintf("CPU status is normal (<%.2f%%)", config.Config.Cpu.LoadLimit)
	}
}
