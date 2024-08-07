package control

import (
	"fmt"

	"github.com/heacat/heacat-api/src/config"
	"github.com/heacat/heacat-api/src/notifier"
	"github.com/heacat/heacat-api/src/utils"
)

type disk_status struct {
	Filesystem string  `json:"filesystem"`
	Type       string  `json:"type"`
	Size       string  `json:"size"`
	Used       string  `json:"used"`
	Available  string  `json:"available"`
	Percent    float64 `json:"percent"`
	Mountpoint string  `json:"mountpoint"`
}

func CheckDiskStatus() (any, string) {
	disks := utils.GetDiskInfo(config.Config.Disk.Unit)
	var anormal_disk_status []disk_status
	anormal_disk_status_exists := false

	for _, disk := range disks {
		if disk.Percent >= float64(config.Config.Disk.PartUseLimit) {
			anormal_disk_status = append(anormal_disk_status, disk_status{
				Filesystem: disk.Filesystem,
				Type:       disk.Type,
				Size:       disk.Size,
				Used:       disk.Used,
				Available:  disk.Available,
				Percent:    disk.Percent,
				Mountpoint: disk.Mountpoint,
			})
			anormal_disk_status_exists = true
		}
	}
	if !anormal_disk_status_exists {
		notifier.SendMessage(false, "disk", fmt.Sprintf("[%s] [🟢 Info] Disk status is normal (<%d%%)", config.Config.Alarm.ServerNickName, config.Config.Disk.PartUseLimit))
		return anormal_disk_status, fmt.Sprintf("Disk status is normal (<%d%%)", config.Config.Disk.PartUseLimit)
	} else {
		message := fmt.Sprintf("[%s] [🔴 Alert] The following partitions are above the threshold (>%d%%):\n--------------------\n", config.Config.Alarm.ServerNickName, config.Config.Disk.PartUseLimit)
		for _, disk := range anormal_disk_status {
			message += fmt.Sprintf("Partition: %s\nMountpoint: %s\nUsed: %s\nAvailable: %s\nPercent: %.2f%%\n--------------------\n", disk.Filesystem, disk.Mountpoint, disk.Used, disk.Available, disk.Percent)
		}
		notifier.SendMessage(true, "disk", message)
		return anormal_disk_status, fmt.Sprintf("Anormal disk status found (>%d%%)", config.Config.Disk.PartUseLimit)
	}
}
