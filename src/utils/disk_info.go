package utils

import (
	"strconv"

	"github.com/heacat/heacat-api/src/config"
	"github.com/shirou/gopsutil/v4/disk"
)

type disk_info_t struct {
	Filesystem string  `json:"filesystem"`
	Type       string  `json:"type"`
	Size       float64 `json:"size"`
	Used       float64 `json:"used"`
	Available  float64 `json:"available"`
	Percent    float64 `json:"percent"`
	Mountpoint string  `json:"mountpoint"`
}

func GetDiskInfo() []disk_info_t {
	disks, _ := disk.Partitions(true)
	var disk_info []disk_info_t

	for _, d := range disks {
		found := false
		for _, fs := range config.Config.Disk.FileSystems {
			if fs == d.Fstype {
				found = true
				break
			}
		}
		if found {
			usage, _ := disk.Usage(d.Mountpoint)
			size, _ := strconv.ParseFloat(convertToUnit(usage.Total, "GB"), 64)
			used, _ := strconv.ParseFloat(convertToUnit(usage.Used, "GB"), 64)
			available, _ := strconv.ParseFloat(convertToUnit(usage.Free, "GB"), 64)
			disk_info = append(disk_info, disk_info_t{
				Filesystem: d.Device,
				Type:       d.Fstype,
				Size:       size,
				Used:       used,
				Available:  available,
				Percent:    usage.UsedPercent,
				Mountpoint: d.Mountpoint,
			})
		}
	}
	return disk_info
}
