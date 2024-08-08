package utils

import (
	"github.com/heacat/heacat-api/src/config"
	"github.com/shirou/gopsutil/v4/disk"
)

type disk_info_t struct {
	Filesystem string  `json:"filesystem"`
	Type       string  `json:"type"`
	Size       string  `json:"size"`
	Used       string  `json:"used"`
	Available  string  `json:"available"`
	Percent    float64 `json:"percent"`
	Mountpoint string  `json:"mountpoint"`
}

func GetDiskInfo(unit string) []disk_info_t {
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
			disk_info = append(disk_info, disk_info_t{
				Filesystem: d.Device,
				Type:       d.Fstype,
				Size:       convertToUnit(usage.Total, unit),
				Used:       convertToUnit(usage.Used, unit),
				Available:  convertToUnit(usage.Free, unit),
				Percent:    usage.UsedPercent,
				Mountpoint: d.Mountpoint,
			})
		}
	}
	return disk_info
}
