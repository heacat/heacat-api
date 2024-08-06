package api

import (
	"fmt"
	"os"
	"os/exec"
	"regexp"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

type package_managers struct {
	Manager string `json:"manager"`
	Count   int    `json:"count"`
}

func detect_package_managers() []package_managers {
	detected := []package_managers{}
	managers := []string{"apt", "dnf", "yum", "zypper", "pacman", "emerge", "apk", "pkg", "kpkg", "swupd", "flatpak", "snap", "nix", "brew"}

	var isRPM bool
	var isDEB bool

	for _, manager := range managers {
		_, err := exec.LookPath(manager)
		if err == nil {
			var count []byte

			switch manager {
			case "apt", "dpkg":
				if isDEB {
					continue
				}
				count, _ = exec.Command("dpkg", "-l").Output()
				isDEB = true
			case "dnf", "yum", "zypper", "rpm":
				if isRPM {
					continue
				}
				count, _ = exec.Command("rpm", "-qa").Output()
				isRPM = true
			case "pacman":
				count, _ = exec.Command("pacman", "-Q").Output()
			case "emerge":
				count, _ = exec.Command("eix", "-I").Output()
			case "apk":
				count, _ = exec.Command("apk", "info").Output()
			case "pkg":
				count, _ = exec.Command("pkg", "info").Output()
			case "kpkg":
				count, _ = exec.Command("kpkg", "list", "--installed").Output()
			case "swupd":
				count, _ = exec.Command("swupd", "bundle-list").Output()
			case "flatpak":
				count, _ = exec.Command("flatpak", "list").Output()
			case "snap":
				count, _ = exec.Command("snap", "list").Output()
			case "nix":
				count, _ = exec.Command("nix-env", "-q").Output()
			case "brew":
				count, _ = exec.Command("brew", "list").Output()
			default:
				count = []byte("0")
			}

			detected = append(detected, package_managers{
				Manager: manager,
				Count:   len(strings.Split(string(count), "\n")) - 1,
			})
		}
	}
	return detected
}

func sys_info(c *gin.Context) {
	os_name_r, _ := regexp.Compile(`NAME="(.+)"`)
	os_name_pretty_r, _ := regexp.Compile(`PRETTY_NAME="(.+)"`)
	kernel_version_r, _ := regexp.Compile(`Linux\sversion\s([a-zA-Z0-9\.\-\_]+)\s`)
	gpu_name_r, _ := regexp.Compile(`((VGA compatible|3D)\scontroller:\s)([\w\s\,\.\[\]\/\(\/\)\-])+\s`)
	memory_free_r, _ := regexp.Compile(`MemFree:\s+([0-9]+)\skB`)
	memory_total_r, _ := regexp.Compile(`MemTotal:\s+([0-9]+)\skB`)

	hostname, _ := os.Hostname()
	os_release, _ := os.ReadFile("/etc/os-release")
	kernel, _ := os.ReadFile("/proc/version")
	uptime, _ := os.ReadFile("/proc/uptime")
	gpu, _ := exec.Command("lspci").Output()
	memory, _ := os.ReadFile("/proc/meminfo")

	uptime_seconds, _ := strconv.ParseFloat(strings.Split(string(uptime), " ")[0], 64)
	uptime_days := int(uptime_seconds / 86400)
	uptime_hours := int(uptime_seconds/3600) % 24
	uptime_minutes := int(uptime_seconds/60) % 60

	gpu_match := []string{}
	for _, element := range gpu_name_r.FindAllStringSubmatch(string(gpu), -1) {
		gpu_match = append(gpu_match, element[0])
	}

	c.JSON(200, gin.H{
		"hostname":         string(hostname),
		"os_name":          os_name_r.FindStringSubmatch(string(os_release))[1],
		"os_name_pretty":   os_name_pretty_r.FindStringSubmatch(string(os_release))[1],
		"kernel_version":   kernel_version_r.FindStringSubmatch(string(kernel))[1],
		"uptime":           fmt.Sprintf("%d days, %d hours, %d minutes", uptime_days, uptime_hours, uptime_minutes),
		"package_managers": detect_package_managers(),
		"cpu":              get_cpu_info().Name,
		"gpu":              gpu_match,
		"memory":           fmt.Sprintf("%s kB / %s kB", memory_free_r.FindStringSubmatch(string(memory))[1], memory_total_r.FindStringSubmatch(string(memory))[1]),
	})
}
