package utils

import (
	"os"
	"strconv"
	"strings"

	"github.com/shirou/gopsutil/v4/cpu"
)

type cpu_info_t struct {
	Name       string  `json:"name"`
	TotalCores int     `json:"totalcores"`
	Frequency  float64 `json:"frequency"`
	LoadAvg    float64 `json:"loadavg"`
}

func GetCPUInfo() cpu_info_t {
	load_avg_file, _ := os.ReadFile("/proc/loadavg")
	load_avg := strings.Split(string(load_avg_file), " ")
	load_avg_float, _ := strconv.ParseFloat(load_avg[1], 64)

	cpu, _ := cpu.Info()
	cpu_info := cpu_info_t{
		Name:       cpu[0].ModelName,
		TotalCores: len(cpu),
		Frequency:  cpu[0].Mhz,
		LoadAvg:    load_avg_float,
	}
	return cpu_info
}
