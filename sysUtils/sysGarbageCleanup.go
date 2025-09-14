package sysUtils

import (
	"context"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"time"
	"winstore/model"

	"github.com/wailsapp/wails/v2/pkg/runtime"
	"golang.org/x/sys/windows/registry"
)

// SysExpiredDocuments 系统临时文件清理
func SysExpiredDocuments(ctx context.Context, times time.Duration) error {
	tempDir := os.TempDir()                           // 获取当前用户临时目录
	cutoff := time.Now().Add(times * -24 * time.Hour) // times天前的临时文件
	err := filepath.Walk(tempDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		// 跳过目录和最近修改的文件
		if !info.IsDir() && info.ModTime().Before(cutoff) {
			if err := os.Remove(path); err == nil {
				runtime.EventsEmit(ctx, "ExpiredDocuments", "已清理临时文件:"+path)
			}
		}
		return nil
	})
	if err != nil {
		runtime.EventsEmit(ctx, "ExpiredDocuments", "清理垃圾文件失败！")
		runtime.EventsOff(ctx, "ExpiredDocuments")
	}
	runtime.EventsOff(ctx, "ExpiredDocuments")
	return nil
}

// SysGarbageCleanup 系统垃圾清理
// 0: 默认垃圾清理配置
func SysGarbageCleanup(ctx context.Context, config int) error {
	// 使用配置集执行自动清理
	cleanCmd := exec.Command("cleanmgr", "/sagerun:1")
	if err := cleanCmd.Start(); err != nil {
		runtime.LogErrorf(ctx, "清理失败: %v\n", err)
	}
	return nil
}

func configureCleanupOptions(options model.CleanupOptions) {
	// 注册表路径：HKEY_LOCAL_MACHINE\SOFTWARE\Microsoft\Windows\CurrentVersion\Explorer\VolumeCaches
	basePath := `SOFTWARE\Microsoft\Windows\CurrentVersion\Explorer\VolumeCaches\`

	// 要启用的清理项及其状态值
	targets := map[string]uint32{
		`Downloaded Program Files`: 0x2, // 下载的程序文件
		`Internet Cache Files`:     0x2, // 互联网缓存文件
		`Temporary Files`:          0x2, // 临时文件
		`Thumbnail Cache`:          0x2, // 缩略图缓存
		`Old ChkDsk Files`:         0x2, // 旧版磁盘检查文件
		`Recycle Bin`:              0x2, // 回收站
		`Temporary Setup Files`:    0x2, // 临时安装文件
	}
	for subkey, value := range targets {
		keyPath := basePath + subkey
		key, err := registry.OpenKey(registry.LOCAL_MACHINE, keyPath, registry.SET_VALUE)
		if err != nil {
			fmt.Printf("无法打开注册表项 %s: %v\n", subkey, err)
			continue
		}
		if err := key.SetDWordValue("StateFlags0001", value); err != nil {
			fmt.Printf("设置 %s 失败: %v\n", subkey, err)
		} else {
			fmt.Printf("已配置: %s = %d\n", subkey, value)
		}
		key.Close()
	}
}
