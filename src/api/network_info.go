package api

import "github.com/gin-gonic/gin"

func network_info(c *gin.Context) {
	c.JSON(200, gin.H{
		"network": "network",
	})
}
