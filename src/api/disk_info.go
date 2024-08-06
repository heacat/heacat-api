package api

import (
	"github.com/heacat/heacat-api/src/config"
	"github.com/shirou/gopsutil/v4/disk"
)

type disk_info_t struct {
	Filesystem string `json:"filesystem"`
	Type       string `json:"type"`
	Size       string `json:"size"`
	Used       string `json:"used"`
	Available  string `json:"available"`
	Percent    string `json:"percent"`
	Mounted    string `json:"mounted"`
}

func get_disk_info() []disk_info_t {
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
				Size:       convert_to_unit(usage.Total, "GB"),
				Used:       convert_to_unit(usage.Used, "GB"),
				Available:  convert_to_unit(usage.Free, "GB"),
				Percent:    convert_to_unit(uint64(usage.UsedPercent), "%"),
				Mounted:    d.Mountpoint,
			})
		}
	}
	return disk_info
}
