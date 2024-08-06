package api

type disk_info_t struct {
	Device string `json:"device"`
	Size   string `json:"size"`
	Free   string `json:"free"`
}

func get_disk_info() []disk_info_t {
	return []disk_info_t{}
}
