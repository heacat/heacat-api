package api

import "github.com/gin-gonic/gin"

func memory_info(c *gin.Context) {
	c.JSON(200, gin.H{
		"memory": "memory",
	})
}
