package api

import (
	"os/exec"
	"regexp"
	"strings"
)

type gpu_info_t struct {
	Name          string `json:"name"`
	ID            string `json:"id"`
	KernelModules string `json:"kernel_modules"`
	UsedModules   string `json:"used_modules"`
}

func get_gpu_info() []gpu_info_t {
	var gpu_list []gpu_info_t
	gpu_name_r, _ := regexp.Compile(`(VGA\scompatible\scontroller:|3D\scontroller:)\s([\w\s\,\.\[\]\/\(\/\)\-])+\n`)
	gpu_used_module_r, _ := regexp.Compile(`(Kernel\sdriver\sin\suse:\s)([\w\s\,\.\[\]\/\(\/\)\-])+\n`)
	gpu_module_r, _ := regexp.Compile(`(Kernel\smodules:\s)([\w\s\,\.\[\]\/\(\/\)\-])+\n`)

	gpu_pci_ids, _ := exec.Command("sh", "-c", "lspci | grep -E 'VGA|3D|Display' | awk '{print $1}'").Output()
	trim := strings.Trim(string(gpu_pci_ids), "\n")

	for _, pci_id := range strings.Split(string(trim), "\n") {
		gpu_info, _ := exec.Command("lspci", "-ks", string(pci_id)).Output()
		gpu_list = append(gpu_list, gpu_info_t{
			Name:          strings.Trim(gpu_name_r.FindStringSubmatch(string(gpu_info))[0], "\n"),
			ID:            string(pci_id),
			KernelModules: strings.Trim(gpu_module_r.FindStringSubmatch(string(gpu_info))[0], "\n"),
			UsedModules:   strings.Trim(gpu_used_module_r.FindStringSubmatch(string(gpu_info))[0], "\n"),
		})
	}
	return gpu_list
}
