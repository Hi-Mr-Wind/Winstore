package sysUtils

import (
	"context"
	"fmt"
	"path/filepath"
	"sort"
	"strings"

	"github.com/wailsapp/wails/v2/pkg/runtime"
	"golang.org/x/sys/windows/registry"
)

// SoftwareInfo 存储软件信息
type softwareInfo struct {
	Name                 string
	Version              string
	Publisher            string
	InstallDate          string
	UninstallString      string
	QuietUninstallString string
}

// SoftwareList 软件列表
type SoftwareList struct {
	Name      string
	Version   string
	Publisher string
}

// ListInstalledSoftware 列出已安装的软件
func ListInstalledSoftware(ctx context.Context) ([]SoftwareList, error) {
	soft := make([]SoftwareList, 0)
	softwareList, err := getInstalledSoftware()
	if err != nil {
		runtime.LogErrorf(ctx, "获取软件列表失败: %v\n", err)
		return nil, err
	}

	// 按软件名称排序
	sort.Slice(softwareList, func(i, j int) bool {
		return softwareList[i].Name < softwareList[j].Name
	})

	for _, software := range softwareList {
		softwareList := SoftwareList{
			Name:      fmt.Sprintf("%-40s", truncateString(software.Name, 40)),
			Version:   fmt.Sprintf("%-15s", software.Version),
			Publisher: software.Publisher,
		}
		soft = append(soft, softwareList)
	}
	return soft, nil
}

// 获取已安装软件列表
func getInstalledSoftware() ([]softwareInfo, error) {
	var softwareList []softwareInfo

	// 注册表路径
	registryPaths := []string{
		`SOFTWARE\Microsoft\Windows\CurrentVersion\Uninstall`,
		`SOFTWARE\WOW6432Node\Microsoft\Windows\CurrentVersion\Uninstall`,
	}

	for _, path := range registryPaths {
		k, err := registry.OpenKey(registry.LOCAL_MACHINE, path, registry.ENUMERATE_SUB_KEYS|registry.QUERY_VALUE)
		if err != nil {
			continue // 跳过无法打开的路径
		}
		subkeys, err := k.ReadSubKeyNames(-1)
		if err != nil {
			continue
		}
		for _, subkey := range subkeys {
			subkeyPath := filepath.Join(path, subkey)
			sk, err := registry.OpenKey(registry.LOCAL_MACHINE, subkeyPath, registry.QUERY_VALUE)
			if err != nil {
				continue
			}

			displayName, _, _ := sk.GetStringValue("DisplayName")
			// 如果没有DisplayName，跳过此项
			if displayName == "" {
				err := sk.Close()
				if err != nil {
					return nil, err
				}
				continue
			}

			version, _, _ := sk.GetStringValue("DisplayVersion")
			publisher, _, _ := sk.GetStringValue("Publisher")
			installDate, _, _ := sk.GetStringValue("InstallDate")
			uninstallString, _, _ := sk.GetStringValue("UninstallString")
			quietUninstallString, _, _ := sk.GetStringValue("QuietUninstallString")
			// 过滤掉 windows 系统软件
			if strings.TrimSpace(publisher) == "Microsoft Corporation" ||
				strings.TrimSpace(publisher) == "Microsoft Corporations" {
				err := sk.Close()
				if err != nil {
					return nil, err
				}
				continue
			}
			software := softwareInfo{
				Name:                 displayName,
				Version:              version,
				Publisher:            publisher,
				InstallDate:          installDate,
				UninstallString:      uninstallString,
				QuietUninstallString: quietUninstallString,
			}

			softwareList = append(softwareList, software)
			e := sk.Close()
			if e != nil {
				return nil, e
			}
		}
		er := k.Close()
		if er != nil {
			return nil, er
		}
	}

	return softwareList, nil
}

// 截断字符串
func truncateString(s string, maxLen int) string {
	if len(s) <= maxLen {
		return s
	}
	return s[:maxLen-3] + "..."
}
