package api

import (
	"os"
	"os/exec"
	"regexp"
	"strings"
)

type memory_info_t struct {
	Physical struct {
		Free  string `json:"free"`
		Total string `json:"total"`
	} `json:"physical"`
	Virtual []struct {
		Free     string `json:"free"`
		Total    string `json:"total"`
		Path     string `json:"path"`
		Priority string `json:"priority"`
	} `json:"virtual"`
}

func get_mem_info() memory_info_t {
	memory_info := memory_info_t{}

	memory_free_r, _ := regexp.Compile(`MemFree:\s+([0-9]+)\skB`)
	memory_total_r, _ := regexp.Compile(`MemTotal:\s+([0-9]+)\skB`)

	phy_memory, _ := os.ReadFile("/proc/meminfo")

	virt_memory, _ := exec.Command("swapon", "--noheadings").Output()
	t_virt_memory := strings.Trim(string(virt_memory), "\n")
	s_virt_memory := strings.Split(t_virt_memory, "\n")

	memory_info.Physical.Free = memory_free_r.FindStringSubmatch(string(phy_memory))[1] + " kB"
	memory_info.Physical.Total = memory_total_r.FindStringSubmatch(string(phy_memory))[1] + " kB"

	for _, line := range s_virt_memory {
		columns := strings.Fields(line)
		if len(columns) > 1 {
			memory_info.Virtual = append(memory_info.Virtual, struct {
				Free     string `json:"free"`
				Total    string `json:"total"`
				Path     string `json:"path"`
				Priority string `json:"priority"`
			}{
				Free:     columns[3],
				Total:    columns[2],
				Path:     columns[0],
				Priority: columns[4],
			})
		}
	}

	return memory_info
}
