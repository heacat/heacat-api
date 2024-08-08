package utils

import (
	"os/exec"
	"strings"
)

type package_managers struct {
	Manager string `json:"manager"`
	Count   int    `json:"count"`
}

func GetPkgInfo() []package_managers {
	detected := []package_managers{}
	managers := []string{"apt", "dnf", "yum", "zypper", "pacman", "emerge", "apk", "pkg", "kpkg", "swupd", "flatpak", "snap", "nix", "brew"}

	var isRPM bool
	var isDEB bool

	for _, manager := range managers {
		_, err := exec.LookPath(manager)
		if err == nil {
			var count []byte

			switch manager {
			case "apt", "dpkg":
				if isDEB {
					continue
				}
				count, _ = exec.Command("dpkg", "-l").Output()
				isDEB = true
			case "dnf", "yum", "zypper", "rpm":
				if isRPM {
					continue
				}
				count, _ = exec.Command("rpm", "-qa").Output()
				isRPM = true
			case "pacman":
				count, _ = exec.Command("pacman", "-Q").Output()
			case "emerge":
				count, _ = exec.Command("eix", "-I").Output()
			case "apk":
				count, _ = exec.Command("apk", "info").Output()
			case "pkg":
				count, _ = exec.Command("pkg", "info").Output()
			case "kpkg":
				count, _ = exec.Command("kpkg", "list", "--installed").Output()
			case "swupd":
				count, _ = exec.Command("swupd", "bundle-list").Output()
			case "flatpak":
				count, _ = exec.Command("flatpak", "list").Output()
			case "snap":
				count, _ = exec.Command("snap", "list").Output()
			case "nix":
				count, _ = exec.Command("nix-env", "-q").Output()
			case "brew":
				count, _ = exec.Command("brew", "list").Output()
			default:
				count = []byte("0")
			}

			detected = append(detected, package_managers{
				Manager: manager,
				Count:   len(strings.Split(string(count), "\n")) - 1,
			})
		}
	}
	return detected
}
