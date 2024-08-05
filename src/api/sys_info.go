package api

import (
	"os"
	"os/exec"
	"strings"

	"github.com/gin-gonic/gin"
)

func sys_info(c *gin.Context) {
	hostname, _ := os.Hostname()
	os_name, _ := exec.Command("bash", "-c", "cat /etc/os-release | grep ^NAME | cut -d= -f2 | tr -d '\"'").Output()
	os_name_pretty, _ := exec.Command("bash", "-c", "cat /etc/os-release | grep ^PRETTY_NAME | cut -d= -f2 | tr -d '\"'").Output()
	kernel, _ := exec.Command("uname", "-r").Output()
	uptime, _ := exec.Command("uptime", "-p").Output()
	package_manager, _ := exec.Command("bash", "-c", "which apt-get > /dev/null 2>&1 && echo apt-get || echo yum").Output()
	package_count, _ := exec.Command("bash", "-c", "which apt-get > /dev/null 2>&1 && dpkg -l | wc -l || rpm -qa | wc -l").Output()

	c.JSON(200, gin.H{
		"hostname":        strings.Trim(string(hostname), "\n"),
		"os_name":         strings.Trim(string(os_name), "\n"),
		"os_name_pretty":  strings.Trim(string(os_name_pretty), "\n"),
		"kernel":          strings.Trim(string(kernel), "\n"),
		"uptime":          strings.Trim(string(uptime), "\n"),
		"package_manager": strings.Trim(string(package_manager), "\n"),
		"package_count":   strings.Trim(string(package_count), "\n"),
	})
}
