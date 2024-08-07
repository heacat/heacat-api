package notifier

import (
	"strconv"
	"time"

	"github.com/heacat/heacat-api/src/config"
	"github.com/heacat/heacat-api/src/logger"
)

type alarm_t struct {
	Type bool `json:"type"`
	Date int  `json:"date"`
}

var disk_alarm alarm_t
var cpu_alarm alarm_t
var memory_alarm alarm_t

func SendMessage(anormalState bool, alarmFrom string, message string) {
	var alarm *alarm_t
	switch alarmFrom {
	case "disk":
		config.Config.Disk.PartUseLimit = 90
		alarm = &disk_alarm
	case "cpu":
		config.Config.Cpu.LoadLimit = 90
		alarm = &cpu_alarm
	case "memory":
		config.Config.Memory.UseLimit = 90
		alarm = &memory_alarm
	default:
		logger.Log.Error("Invalid alarm type:", alarmFrom)
		return
	}

	prev_disk_alarm_date := alarm.Date
	now, _ := strconv.Atoi(time.Now().Format("20060102"))
	if !anormalState && alarm.Date == 0 {
		return
	} else if (alarm.Type != anormalState) || alarm.Type != anormalState && (now-prev_disk_alarm_date >= 1) || alarm.Type == anormalState && (now-prev_disk_alarm_date) >= 1 {
		alarm.Type = anormalState
		alarm.Date = now
	} else {
		logger.Log.Info("Alarm already sent in this day, skipping")
		return
	}

	if config.Config.Alarm.Telegram.Enabled {
		TelegramNotifier(message)
	} else if config.Config.Alarm.Slack.Enabled {
		SlackNotifier(message)
	} else {
		logger.Log.Error("No notifier enabled")
	}
}
