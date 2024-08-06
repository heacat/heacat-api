package utils

import (
	"fmt"
	"strconv"

	"github.com/shirou/gopsutil/v4/mem"
)

type mem_info_t struct {
	Physical struct {
		Total     string `json:"total"`
		Used      string `json:"used"`
		Free      string `json:"free"`
		Shared    string `json:"shared"`
		Buffers   string `json:"buffers"`
		Cached    string `json:"cached"`
		Available string `json:"available"`
	} `json:"physical"`
	Swap struct {
		Total string `json:"total"`
		Used  string `json:"used"`
		Free  string `json:"free"`
	} `json:"swap"`
}

func convertToUnit(value uint64, unit string) string {
	switch unit {
	case "GB":
		value = value / 1024 / 1024 / 1024
	case "GiB":
		value = value / 1000 / 1000 / 1000
	case "MB":
		value = value / 1024 / 1024
	case "MiB":
		value = value / 1000 / 1000
	case "KB":
		value = value / 1024
	case "KiB":
		value = value / 1000
	}
	return strconv.FormatUint(value, 10) + unit
}

func GetMemoryInfo(unit string) mem_info_t {
	vmemStat, err := mem.VirtualMemory()
	if err != nil {
		fmt.Printf("Failed to get virtual memory stats: %v\n", err)
	}

	swapStat, err := mem.SwapMemory()
	if err != nil {
		fmt.Printf("Failed to get swap memory stats: %v\n", err)
	}

	mem_info := mem_info_t{
		Physical: struct {
			Total     string `json:"total"`
			Used      string `json:"used"`
			Free      string `json:"free"`
			Shared    string `json:"shared"`
			Buffers   string `json:"buffers"`
			Cached    string `json:"cached"`
			Available string `json:"available"`
		}{
			Total:     convertToUnit(vmemStat.Total, unit),
			Used:      convertToUnit(vmemStat.Total-vmemStat.Available, unit),
			Free:      convertToUnit(vmemStat.Free, unit),
			Shared:    convertToUnit(vmemStat.Shared, unit),
			Buffers:   convertToUnit(vmemStat.Buffers, unit),
			Cached:    convertToUnit(vmemStat.Cached, unit),
			Available: convertToUnit(vmemStat.Available, unit),
		},
		Swap: struct {
			Total string `json:"total"`
			Used  string `json:"used"`
			Free  string `json:"free"`
		}{
			Total: convertToUnit(swapStat.Total, unit),
			Used:  convertToUnit(swapStat.Used, unit),
			Free:  convertToUnit(swapStat.Free, unit),
		},
	}
	return mem_info
}
