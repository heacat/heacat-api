package control

import "github.com/gin-gonic/gin"

func CheckCPUStatus(c *gin.Context) {
	c.JSON(200, gin.H{"status": "ok"})
}