package control

import (
	"github.com/gin-gonic/gin"
	"github.com/heacat/heacat-api/src/config"
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
	anormal_disk_status_exists := "Disk usage is normal"

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
			anormal_disk_status_exists = "Anormal disk usage detected"
		}
	}
	c.JSON(200, gin.H{"disk": anormal_disk_status, "message": anormal_disk_status_exists})
}
