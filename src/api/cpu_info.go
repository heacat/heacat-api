package api

import (
	"github.com/shirou/gopsutil/v4/cpu"
)

type cpu_info_t struct {
	Name       string  `json:"name"`
	TotalCores int     `json:"totalcores"`
	Frequency  float64 `json:"frequency"`
}

func get_cpu_info() cpu_info_t {
	cpu, _ := cpu.Info()
	cpu_info := cpu_info_t{
		Name:       cpu[0].ModelName,
		TotalCores: len(cpu),
		Frequency:  cpu[0].Mhz,
	}
	return cpu_info
}
