package api

import "github.com/shirou/gopsutil/v4/net"

type network_info_t struct {
	Interface string                `json:"interface"`
	MTU       int                   `json:"mtu"`
	IPs       net.InterfaceAddrList `json:"ips"`
	MAC       string                `json:"mac"`
	Flags     []string              `json:"flags"`
}

func get_network_info() []network_info_t {
	networks, _ := net.Interfaces()
	var network_info []network_info_t

	for _, n := range networks {
		network_info = append(network_info, network_info_t{
			Interface: n.Name,
			MTU:       n.MTU,
			IPs:       n.Addrs,
			MAC:       n.HardwareAddr,
			Flags:     n.Flags,
		})
	}
	return network_info
}
