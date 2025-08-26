package sysInfo

import (
	"context"
	"time"

	"github.com/StackExchange/wmi"
	"github.com/wailsapp/wails/v2/pkg/runtime"
)

// WMI 结构体定义
type win32OperatingSystem struct {
	Caption        string
	Version        string
	BuildNumber    string
	OSArchitecture string
	LastBootUpTime time.Time
}

type win32ComputerSystem struct {
	Name                string
	TotalPhysicalMemory uint64
}

type win32Processor struct {
	Name                      string
	NumberOfCores             uint32
	NumberOfLogicalProcessors int
}

type win32PhysicalMemory struct {
	Capacity uint64
}

type WinSysInfo struct {
	// 系统类型
	SysType string
	// 系统版本
	SysVersion string
	// 系统架构
	SysArch string
	// 计算机名
	ComputerName string
	// 总内存
	TotalMemory float64
	// 处理器型号
	CpuType string
	// 处理器核心数
	CpuCoreNum uint32
	// 逻辑处理器数
	CpuLogicNum int
	// 运行时间
	Uptime string
	// 最后启动时间
	LastBootUpTime string
}

func GetSysInfo(ctx context.Context) WinSysInfo {
	// 获取操作系统信息
	var osInfo []win32OperatingSystem
	query := "SELECT Caption, Version, BuildNumber, OSArchitecture, LastBootUpTime FROM Win32_OperatingSystem"
	if err := wmi.Query(query, &osInfo); err != nil {
		runtime.LogErrorf(ctx, "查询操作系统信息失败：%s\n", err.Error())
	}

	// 获取计算机系统信息
	var csInfo []win32ComputerSystem
	if err := wmi.Query("SELECT Name, TotalPhysicalMemory FROM Win32_ComputerSystem", &csInfo); err != nil {
		runtime.LogErrorf(ctx, "查询获取计算机系统信息失败：%s\n", err.Error())
	}

	// 获取处理器信息
	var procInfo []win32Processor
	if err := wmi.Query("SELECT Name, NumberOfCores, NumberOfLogicalProcessors FROM Win32_Processor", &procInfo); err != nil {
		runtime.LogErrorf(ctx, "查询处理器信息失败：%s\n", err.Error())
	}

	// 获取物理内存信息
	var memInfo []win32PhysicalMemory
	if err := wmi.Query("SELECT Capacity FROM Win32_PhysicalMemory", &memInfo); err != nil {
		runtime.LogErrorf(ctx, "获取物理内存信息失败：%s\n", err.Error())
	}

	// 计算总内存
	var totalMemGB float64
	for _, m := range memInfo {
		totalMemGB += float64(m.Capacity) / (1024 * 1024 * 1024)
	}
	// 计算系统运行时间
	uptime := time.Since(osInfo[0].LastBootUpTime).Round(time.Second)

	return WinSysInfo{
		SysType:        osInfo[0].Caption,
		SysVersion:     osInfo[0].Version,
		SysArch:        osInfo[0].OSArchitecture,
		ComputerName:   csInfo[0].Name,
		TotalMemory:    totalMemGB,
		CpuType:        procInfo[0].Name,
		CpuCoreNum:     procInfo[0].NumberOfCores,
		CpuLogicNum:    procInfo[0].NumberOfLogicalProcessors,
		Uptime:         uptime.String(),
		LastBootUpTime: osInfo[0].LastBootUpTime.Local().Format("2006-01-02 15:04:05"),
	}
}
