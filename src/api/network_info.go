package api

type network_info_t struct {
	Interface string `json:"interface"`
	IP        string `json:"ip"`
	MAC       string `json:"mac"`
}

func get_network_info() []network_info_t {
	return []network_info_t{}
}
