package control

import "github.com/gin-gonic/gin"

func CheckMemoryStatus(c *gin.Context) {
	c.JSON(200, gin.H{"status": "ok"})
}
