package api

import (
	"github.com/gin-gonic/gin"

	"github.com/shirou/gopsutil/v4/cpu"
)

func cpu_info(c *gin.Context) {
	info, _ := cpu.Info()
	percent, _ := cpu.Percent(0, false)

	c.JSON(200, gin.H{
		"info":    info,
		"percent": percent,
	})
}
