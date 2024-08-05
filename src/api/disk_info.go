package api

import "github.com/gin-gonic/gin"

func disk_info(c *gin.Context) {
	c.JSON(200, gin.H{
		"disk": "disk",
	})
}
