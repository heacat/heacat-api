package api

import (
	"encoding/json"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/heacat/heacat-api/src/config"
	"github.com/heacat/heacat-api/src/control"
	"github.com/heacat/heacat-api/src/logger"
	"github.com/heacat/heacat-api/src/utils"
)

type registered_routes struct {
	Name string `json:"name"`
	Path string `json:"path"`
}

var routes []registered_routes

func InitAPI() {
	gin.SetMode(gin.ReleaseMode)

	router := gin.Default()

	router.GET("/api/v1/cpu", func(c *gin.Context) {
		cpu := utils.GetCPUInfo()
		c.JSON(200, gin.H{"cpu": cpu})
	})
	router.GET("/api/v1/gpu", func(ctx *gin.Context) {
		gpu := utils.GetGPUInfo()
		ctx.JSON(200, gin.H{"gpu": gpu})
	})
	router.GET("/api/v1/disk", func(ctx *gin.Context) {
		disk := utils.GetDiskInfo(config.Config.Disk.Unit)
		ctx.JSON(200, gin.H{"disk": disk})
	})
	router.GET("/api/v1/memory", func(c *gin.Context) {
		memory := utils.GetMemoryInfo(config.Config.Memory.Unit)
		c.JSON(200, gin.H{"memory": memory})
	})
	router.GET("/api/v1/network", func(ctx *gin.Context) {
		network := utils.GetNetworkInfo()
		ctx.JSON(200, gin.H{"network": network})
	})
	router.GET("/api/v1/sysinfo", func(ctx *gin.Context) {
		sysinfo := utils.GetSysInfo()
		ctx.JSON(200, gin.H{"sysinfo": sysinfo})
	})
	router.GET("/api/v1/check/disk", control.CheckDiskStatus)
	router.GET("/api/v1/check/memory", control.CheckMemoryStatus)
	router.GET("/api/v1/check/cpu", control.CheckCPUStatus)
	router.GET("/", func(c *gin.Context) {
		b, err := json.Marshal(routes)
		if err != nil {
			logger.Log.Error(err)
			c.JSON(500, gin.H{"error": "Internal server error"})
			return
		}
		c.JSON(200, gin.H{"routes": json.RawMessage(b)})
	})

	for _, item := range router.Routes() {
		name := strings.Split(item.Path, "/")
		if len(name) > 0 && name[len(name)-1] != "" {
			new := registered_routes{
				Name: name[len(name)-2] + "/" + name[len(name)-1],
				Path: "http://" + config.Config.API.Host + ":" + strconv.Itoa(config.Config.API.Port) + item.Path,
			}
			routes = append(routes, new)
		}
	}
	logger.Log.Infof("API service started on address: " + config.Config.API.Host + ":" + strconv.Itoa(config.Config.API.Port))
	router.Run(config.Config.API.Host + ":" + strconv.Itoa(config.Config.API.Port))
}
