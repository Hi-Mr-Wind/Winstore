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

// RestartSystem 重启系统
func (a *App) RestartSystem(mes string) error {
	return sysUtils.RestartSystem(mes)
}

// UninstallSoftware 卸载软件
// 优先尝试使用静默卸载程序如果不存在则使用普通卸载程序
// 如果执行卸载程序失败，则尝试删除整个软件目录，并尝试删除注册表项
// 在卸载软件完成之后，尝试删除注册表项
// 如果卸载流程失败，则返回错误
func (a *App) UninstallSoftware(softwareName string) error {
	return sysUtils.UninstallSoftware(a.ctx, softwareName)
}
