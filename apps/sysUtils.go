package apps

import (
	"time"
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

// UpdateApp 更新软件自身
func (a *App) UpdateApp() error {

	return nil
}

// SysExpiredDocuments 清理X天之前的临时文件 默认清理七天前的临时文件
func (a *App) SysExpiredDocuments(times int64) error {
	if times <= 0 {
		times = 7
	}
	return sysUtils.SysExpiredDocuments(a.ctx, time.Duration(times))
}

// SysGarbageCleanup 调用windows系统垃圾清理
func (a *App) SysGarbageCleanup(config int) error {
	return sysUtils.SysGarbageCleanup(a.ctx, config)
}
