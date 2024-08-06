package api

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

func sys_info(c *gin.Context) {
	os_name_r, _ := regexp.Compile(`NAME="(.+)"`)
	os_name_pretty_r, _ := regexp.Compile(`PRETTY_NAME="(.+)"`)
	kernel_version_r, _ := regexp.Compile(`Linux\sversion\s([a-zA-Z0-9\.\-\_]+)\s`)

	hostname, _ := os.Hostname()
	os_release, _ := os.ReadFile("/etc/os-release")
	kernel, _ := os.ReadFile("/proc/version")
	uptime, _ := os.ReadFile("/proc/uptime")

	uptime_seconds, _ := strconv.ParseFloat(strings.Split(string(uptime), " ")[0], 64)
	uptime_days := int(uptime_seconds / 86400)
	uptime_hours := int(uptime_seconds/3600) % 24
	uptime_minutes := int(uptime_seconds/60) % 60

	c.JSON(200, gin.H{
		"hostname":         hostname,
		"os_name":          os_name_r.FindStringSubmatch(string(os_release))[1],
		"os_name_pretty":   os_name_pretty_r.FindStringSubmatch(string(os_release))[1],
		"kernel_version":   kernel_version_r.FindStringSubmatch(string(kernel))[1],
		"uptime":           fmt.Sprintf("%d days, %d hours, %d minutes", uptime_days, uptime_hours, uptime_minutes),
		"package_managers": get_pkg_info(),
		"cpu":              get_cpu_info(),
		"gpu":              get_gpu_info(),
		"memory":           get_mem_info("GB"),
	})
}
