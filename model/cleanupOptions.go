package model

// CleanupOptions 磁盘清理选项
type CleanupOptions struct {
	// 已下载的程序文件
	DownloadedProgramFiles uint32
	// 缓存文件
	InternetCacheFiles uint32
	// 临时文件
	TemporaryFiles uint32
	// 旧程序文件
	OldProgramFiles uint32
	// 缩略图缓存
	ThumbnailCache uint32
	// 回收站
	RecycleBin uint32
	// 临时安装文件
	TemporarySetupFiles uint32
}
