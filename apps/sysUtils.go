package apps

import "winstore/sysInfo"

// GetSysInfo 获取系统信息
func (a *App) GetSysInfo() sysInfo.WinSysInfo {
	return sysInfo.GetSysInfo(a.ctx)
}
