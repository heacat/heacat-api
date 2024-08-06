package api

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

func convert_to_unit(value uint64, unit string) string {
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

func get_mem_info(unit string) mem_info_t {
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
			Total:     convert_to_unit(vmemStat.Total, unit),
			Used:      convert_to_unit(vmemStat.Total-vmemStat.Available, unit),
			Free:      convert_to_unit(vmemStat.Free, unit),
			Shared:    convert_to_unit(vmemStat.Shared, unit),
			Buffers:   convert_to_unit(vmemStat.Buffers, unit),
			Cached:    convert_to_unit(vmemStat.Cached, unit),
			Available: convert_to_unit(vmemStat.Available, unit),
		},
		Swap: struct {
			Total string `json:"total"`
			Used  string `json:"used"`
			Free  string `json:"free"`
		}{
			Total: convert_to_unit(swapStat.Total, unit),
			Used:  convert_to_unit(swapStat.Used, unit),
			Free:  convert_to_unit(swapStat.Free, unit),
		},
	}
	return mem_info
}
