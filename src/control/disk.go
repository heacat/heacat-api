package control

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/heacat/heacat-api/src/config"
	"github.com/heacat/heacat-api/src/notifier"
	"github.com/heacat/heacat-api/src/utils"
)

type disk_status struct {
	Partition  string `json:"partition"`
	Usage      string `json:"usage"`
	Available  string `json:"available"`
	Total      string `json:"total"`
	Percentage int    `json:"percentage"`
	Mountpoint string `json:"mountpoint"`
}

func CheckDiskStatus(c *gin.Context) {
	disks := utils.GetDiskInfo(config.Config.Disk.Unit)
	var anormal_disk_status []disk_status
	anormal_disk_status_exists := false

	for _, disk := range disks {
		if disk.Percent >= float64(config.Config.Disk.PartUseLimit) {
			anormal_disk_status = append(anormal_disk_status, disk_status{
				Partition:  disk.Filesystem,
				Usage:      disk.Used,
				Available:  disk.Available,
				Total:      disk.Size,
				Mountpoint: disk.Mountpoint,
				Percentage: int(disk.Percent),
			})
			anormal_disk_status_exists = true
		}
	}
	if !anormal_disk_status_exists {
		c.JSON(200, gin.H{"disk": anormal_disk_status, "message": "Disk status is normal"})
		message := fmt.Sprintf("[%s] [🟢 Info] Disk status is normal", config.Config.Alarm.ServerNickName)
		notifier.SendMessage(false, "disk", message)
		return
	}
	c.JSON(200, gin.H{"disk": anormal_disk_status, "message": "Anormal disk status found"})

	message := fmt.Sprintf("[%s] [🔴 Alert] The following partitions are above the threshold (%d%%):\n--------------------\n", config.Config.Alarm.ServerNickName, config.Config.Disk.PartUseLimit)
	for _, disk := range anormal_disk_status {
		message += fmt.Sprintf("Partition: %s\nUsage: %s\nAvailable: %s\nTotal: %s\nPercentage: %d\nMountpoint: %s\n--------------------\n", disk.Partition, disk.Usage, disk.Available, disk.Total, disk.Percentage, disk.Mountpoint)
	}
	notifier.SendMessage(true, "disk", message)
}
