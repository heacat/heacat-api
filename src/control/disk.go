package control

import (
	"github.com/gin-gonic/gin"
	"github.com/heacat/heacat-api/src/config"
	"github.com/heacat/heacat-api/src/utils"
)

type disk_status struct {
	Usage      int    `json:"usage"`
	Available  int    `json:"available"`
	Total      int    `json:"total"`
	Percentage int    `json:"percentage"`
	Message    string `json:"message"`
}

func CheckDiskStatus(c *gin.Context) {
	disks := utils.GetDiskInfo()
	var anormal_disk_status []disk_status
	anormal_disk_status_exists := "Disk usage is normal"

	for _, disk := range disks {
		if disk.Percent >= float64(config.Config.Disk.PartUseLimit) {
			anormal_disk_status = append(anormal_disk_status, disk_status{
				Usage:      int(disk.Used),
				Available:  int(disk.Available),
				Total:      int(disk.Size),
				Percentage: int(disk.Percent),
			})
			anormal_disk_status_exists = "Anormal disk usage detected"
		}
	}
	c.JSON(200, gin.H{"disk": anormal_disk_status, "message": anormal_disk_status_exists})
}
