package notifier

import (
	"strconv"
	"time"

	"github.com/heacat/heacat-api/src/config"
	"github.com/heacat/heacat-api/src/logger"
)

type alarm_t struct {
	Type bool      `json:"type"`
	Date time.Time `json:"date"`
}

var disk_alarm alarm_t
var cpu_alarm alarm_t
var memory_alarm alarm_t

func SendMessage(alarm_type bool, alarm_from string, message string) {
	config.Config.Disk.PartUseLimit = 90
	now := time.Now()
	curr_date_atoi, _ := strconv.Atoi(now.Format("20060102"))
	switch alarm_from {
	case "disk":
		prev_disk_alarm_date, _ := strconv.Atoi(disk_alarm.Date.Format("20060102"))
		if (!disk_alarm.Type) || (disk_alarm.Type != alarm_type) || (disk_alarm.Type == alarm_type && (curr_date_atoi-prev_disk_alarm_date) >= 1) {
			disk_alarm.Type = alarm_type
			disk_alarm.Date = now
		} else {
			logger.Log.Info("Alarm already sent in the last 24 hours. Skipping...")
			return
		}
	case "cpu":
		prev_cpu_alarm_date, _ := strconv.Atoi(cpu_alarm.Date.Format("20060102"))
		if (!cpu_alarm.Type) || (cpu_alarm.Type != alarm_type && (curr_date_atoi-prev_cpu_alarm_date) >= 1) {
			cpu_alarm.Type = alarm_type
			cpu_alarm.Date = now
		} else {
			logger.Log.Info("Alarm already sent in the last 24 hours. Skipping...")
			return
		}
	case "memory":
		prev_memory_alarm_date, _ := strconv.Atoi(memory_alarm.Date.Format("20060102"))
		if (!memory_alarm.Type) || (memory_alarm.Type != alarm_type && (curr_date_atoi-prev_memory_alarm_date) >= 1) {
			memory_alarm.Type = alarm_type
			memory_alarm.Date = now
		} else {
			logger.Log.Info("Alarm already sent in the last 24 hours. Skipping...")
			return
		}
	}

	if config.Config.Alarm.Telegram.Enabled {
		TelegramNotifier(message)
	} else if config.Config.Alarm.Slack.Enabled {
		SlackNotifier(message)
	} else {
		logger.Log.Error("No notifier enabled, please enable at least one notifier. Skipping...")
	}
}
