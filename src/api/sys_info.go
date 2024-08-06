package api

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type sys_info_t struct {
	Hostname        string             `json:"hostname"`
	OsName          string             `json:"os_name"`
	OsNamePretty    string             `json:"os_name_pretty"`
	KernelVersion   string             `json:"kernel_version"`
	Uptime          string             `json:"uptime"`
	PackageManagers []package_managers `json:"package_managers"`
	Cpu             cpu_info_t         `json:"cpu"`
	Gpu             []gpu_info_t       `json:"gpu"`
	Memory          mem_info_t         `json:"memory"`
	Network         []network_info_t   `json:"network"`
	Disks           []disk_info_t      `json:"disks"`
}

func sys_info() sys_info_t {
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

	return sys_info_t{
		Hostname:        hostname,
		OsName:          os_name_r.FindStringSubmatch(string(os_release))[1],
		OsNamePretty:    os_name_pretty_r.FindStringSubmatch(string(os_release))[1],
		KernelVersion:   kernel_version_r.FindStringSubmatch(string(kernel))[1],
		Uptime:          fmt.Sprintf("%d days, %d hours, %d minutes", uptime_days, uptime_hours, uptime_minutes),
		PackageManagers: get_pkg_info(),
		Cpu:             get_cpu_info(),
		Gpu:             get_gpu_info(),
		Memory:          get_mem_info("GB"),
		Network:         get_network_info(),
		Disks:           get_disk_info(),
	}
}
