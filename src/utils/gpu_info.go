package utils

import (
	"os/exec"
	"strings"

	"github.com/dlclark/regexp2"
)

type gpu_info_t struct {
	Name          string `json:"name"`
	ID            string `json:"id"`
	KernelModules string `json:"kernel_modules"`
	UsedModules   string `json:"used_modules"`
}

func GetGPUInfo() []gpu_info_t {
	var gpu_list []gpu_info_t
	gpu_name_r := regexp2.MustCompile(`(?<=(VGA\scompatible\scontroller:|3D\scontroller:)\s)([\w\s\,\.\[\]\/\(\/\)\-])+\n`, 0)
	gpu_used_module_r := regexp2.MustCompile(`(?<=(Kernel\sdriver\sin\suse:)\s)([\w\s\,\.\[\]\/\(\/\)\-])+\n`, 0)
	gpu_module_r := regexp2.MustCompile(`(?<=(Kernel\smodules:)\s)([\w\s\,\.\[\]\/\(\/\)\-])+`, 0)

	gpu_pci_ids, _ := exec.Command("sh", "-c", "lspci | grep -E 'VGA|3D|Display' | awk '{print $1}'").Output()
	trim := strings.Trim(string(gpu_pci_ids), "\n")

	for _, pci_id := range strings.Split(string(trim), "\n") {
		gpu_info, _ := exec.Command("lspci", "-ks", string(pci_id)).Output()

		gpu_name, _ := gpu_name_r.FindStringMatch(string(gpu_info))
		gpu_module, _ := gpu_module_r.FindStringMatch(string(gpu_info))
		gpu_used_module, _ := gpu_used_module_r.FindStringMatch(string(gpu_info))
		gpu_list = append(gpu_list, gpu_info_t{
			Name:          strings.Trim(gpu_name.String(), "\n"),
			ID:            string(pci_id),
			KernelModules: strings.Trim(gpu_module.String(), "\n"),
			UsedModules:   strings.Trim(gpu_used_module.String(), "\n"),
		})
	}
	return gpu_list
}
