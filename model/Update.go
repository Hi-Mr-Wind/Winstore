package model

type UpdateInfo struct {
	// 版本号
	Version string `json:"version"`
	// 下载地址
	DownloadURL string `json:"download_url"`
	// 描述
	Description string `json:"description"`
	// 文件大小
	Size int64 `json:"size"`
	// 哈希
	Checksum string `json:"checksum"`
}
