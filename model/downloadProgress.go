package model

type DownloadProgress struct {
	//进度百分比
	Percentage float64 `json:"percentage"`
	//已下载
	Downloaded string `json:"downloaded"`
	//总大小 MB或者KB
	Total string `json:"total"`
	//速度 kb/s或mb/s
	Speed string `json:"speed"`
}
