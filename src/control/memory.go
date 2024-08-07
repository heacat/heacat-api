package control

import (
	"fmt"
	"strconv"

	"github.com/heacat/heacat-api/src/config"
	"github.com/heacat/heacat-api/src/notifier"
	"github.com/heacat/heacat-api/src/utils"
)

func CheckMemoryStatus() (any, string) {
	mem := utils.GetMemoryInfo("")

	free, _ := strconv.ParseFloat(mem.Physical.Free, 64)
	used, _ := strconv.ParseFloat(mem.Physical.Used, 64)
	used_percent := (free / used) * 100

	if int(used_percent) >= config.Config.Memory.UseLimit {
		notifier.SendMessage(true, "memory", fmt.Sprintf("[%s] [🔴 Alert] Anormal memory status found (>%d%%)", config.Config.Alarm.ServerNickName, config.Config.Memory.UseLimit))
		return mem, fmt.Sprintf("Anormal memory status found (>%d%%)", config.Config.Memory.UseLimit)
	} else {
		notifier.SendMessage(false, "memory", fmt.Sprintf("[%s] [🟢 Info] Memory status is normal (<%d%%)", config.Config.Alarm.ServerNickName, config.Config.Memory.UseLimit))
		return mem, fmt.Sprintf("Memory status is normal (<%d%%)", config.Config.Memory.UseLimit)
	}
}
