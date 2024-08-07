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
		c.JSON(200, gin.H{"cpu": utils.GetCPUInfo()})
	})
	router.GET("/api/v1/gpu", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{"gpu": utils.GetGPUInfo()})
	})
	router.GET("/api/v1/disk", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{"disk": utils.GetDiskInfo(config.Config.Disk.Unit)})
	})
	router.GET("/api/v1/memory", func(c *gin.Context) {
		c.JSON(200, gin.H{"memory": utils.GetMemoryInfo(config.Config.Memory.Unit)})
	})
	router.GET("/api/v1/network", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{"network": utils.GetNetworkInfo()})
	})
	router.GET("/api/v1/sysinfo", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{"sysinfo": utils.GetSysInfo()})
	})
	router.GET("/api/v1/check/disk", func(ctx *gin.Context) {
		arr, message := control.CheckDiskStatus()
		ctx.JSON(200, gin.H{"disk": arr, "message": message})
	})
	router.GET("/api/v1/check/memory", func(ctx *gin.Context) {
		arr, message := control.CheckMemoryStatus()
		ctx.JSON(200, gin.H{"memory": arr, "message": message})
	})
	router.GET("/api/v1/check/cpu", func(ctx *gin.Context) {
		arr, message := control.CheckCPUStatus()
		ctx.JSON(200, gin.H{"cpu": arr, "message": message})
	})
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
