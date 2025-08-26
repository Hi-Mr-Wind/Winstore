package apps

import (
	"winstore/sysUtils"
)

// GetSysInfo 获取系统信息
func (a *App) GetSysInfo() sysUtils.WinSysInfo {
	return sysUtils.GetSysInfo(a.ctx)
}

// GetSoftwareList 获取已安装的软件列表
func (a *App) GetSoftwareList() ([]sysUtils.SoftwareList, error) {
	return sysUtils.ListInstalledSoftware(a.ctx)
}
