package api

import (
	"encoding/json"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/heacat/heacat-api/src/config"
	"github.com/heacat/heacat-api/src/logger"
)

type registered_routes struct {
	Name string `json:"name"`
	Path string `json:"path"`
}

var routes []registered_routes

func InitAPI() {
	gin.SetMode(gin.ReleaseMode)

	router := gin.Default()

	router.GET("/api/v1/cpu", cpu_info)
	router.GET("/api/v1/disk", disk_info)
	router.GET("/api/v1/memory", memory_info)
	router.GET("/api/v1/network", network_info)
	router.GET("/api/v1/sysinfo", sys_info)
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
				Name: name[len(name)-1],
				Path: "http://" + config.Config.API.Host + ":" + strconv.Itoa(config.Config.API.Port) + item.Path,
			}
			routes = append(routes, new)
		}
	}
	logger.Log.Infof("API service started on address: " + config.Config.API.Host + ":" + strconv.Itoa(config.Config.API.Port))
	router.Run(config.Config.API.Host + ":" + strconv.Itoa(config.Config.API.Port))
}
