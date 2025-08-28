package sysUtils

import (
	"context"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"sort"
	"strings"
	"winstore/model"

	"github.com/wailsapp/wails/v2/pkg/runtime"
	"golang.org/x/sys/windows/registry"
)

// SoftwareInfo 存储软件信息
type softwareInfo struct {
	// 软件名称
	Name string
	// 软件版本
	Version string
	// 软件供应商
	Publisher string
	// 安装时间
	InstallDate string
	// 卸载程序地址
	UninstallString string
	// 静默卸载程序地址
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
	// 遍历注册表路径
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

// UninstallSoftware 卸载软件
// 优先尝试使用静默卸载程序如果不存在则使用普通卸载程序
// 如果执行卸载程序失败，则尝试删除整个软件目录，并尝试删除注册表项
// 在卸载软件完成之后，尝试删除注册表项
// 如果卸载流程失败，则返回错误
func UninstallSoftware(ctx context.Context, softwareName string) error {
	softwareList, err := getInstalledSoftware()
	if err != nil {
		runtime.LogErrorf(ctx, "卸载软件失败！无法获取软件列表: %v\n", err)
		return err
	}
	var software softwareInfo
	for _, data := range softwareList {
		//找到软件信息后退出
		if strings.EqualFold(data.Name, softwareName) {
			software = data
			break
		}
	}
	if software.Name == "" {
		runtime.LogErrorf(ctx, "软件不存在！")
		return fmt.Errorf("软件不存在！")
	}
	fmt.Printf("#%#v\n", &software)
	// 尝试使用静默卸载程序
	var cmd string
	if software.QuietUninstallString != "" {
		cmd = software.QuietUninstallString
	} else { //如果没有静默卸载程序，则使用普通卸载程序
		cmd = software.UninstallString
	}
	cmdPath, args := parseUninstallString(cmd)

	// 执行卸载
	cmdRun := exec.Command(cmdPath, args...)

	if err := cmdRun.Start(); err != nil {
		runtime.LogErrorf(ctx, "卸载过程中出错: %v\n", err)
	} else {
		runtime.LogInfof(ctx, "\n%s 卸载程序执行开始。卸载进程PID：%v\n", softwareName, cmdRun.Process.Pid)
	}

	//异步监测卸载进程的运行情况
	go func() {
		err := cmdRun.Wait()
		if err != nil {
			runtime.EventsEmit(ctx, "UninstallSoftware_"+softwareName, fmt.Sprintf("%s卸载进程异常：%s", softwareName, err.Error()))
			runtime.LogErrorf(ctx, "卸载过程中出错: %v\n", err)
			err := cmdRun.Process.Kill()
			if err != nil {
				runtime.LogErrorf(ctx, "卸载进程关闭失败: %v\n", err)
				return
			}
			return
		}
		runtime.LogInfof(ctx, "\n%s 卸载程序执行完成,检查软件残留\n", softwareName)
		// 卸载后，检查软件目录是否还有该软件，防止用户操作错误导致未卸载
		installedSoftware, err := getInstalledSoftware()
		if err != nil {
			runtime.LogErrorf(ctx, "获取软件列表失败: %v\n", err)
			return
		}
		for _, data := range installedSoftware {
			if strings.EqualFold(data.Name, softwareName) {
				runtime.LogErrorf(ctx, "软件未卸载成功！请手动卸载！")
				runtime.EventsEmit(ctx, "UninstallSoftware_"+softwareName, fmt.Sprintf("%s 未卸载成功！请再次尝试！", softwareName))
				return
			}
		}

		// 检查软件目录是否有残留
		dir := filepath.Dir(cmdPath)
		_, osErr := os.Stat(dir)
		if osErr != nil {
			// 如果 卸载后的软件目录仍然存在，则插入数据
			if !os.IsNotExist(osErr) {
				residue := model.Residue{
					ResiduePath:    dir,
					ResidueAppName: softwareName,
				}
				residue.InsertData()
				// 标记删除残留目录
				err := MarkForDeletion(dir)
				if err != nil {
					runtime.LogErrorf(ctx, "删除失败: %v\n", err)
					return
				}
				runtime.EventsEmit(ctx, "UninstallSoftware_"+softwareName, fmt.Sprintf("%s 卸载完成！将在计算机重启后删除残留文件！", softwareName))
			}
			//else {
			//--暂时不再提供删除注册表项--

			//}
			return
		}

	}()
	return nil
}

// 截断字符串
func truncateString(s string, maxLen int) string {
	if len(s) <= maxLen {
		return s
	}
	return s[:maxLen-3] + "..."
}

// parseUninstallString 解析卸载字符串，返回可执行路径和参数
// 经过测试，发现软件卸载字符串格式并不统一，大致分为两种格式，一种是带引号的，另一种是不带引号的
// 卸载所需参数也不尽相同，有些参数在注册表中标注的形式为 -- 有些参数在注册表标为的格式为 /
func parseUninstallString(uninstallStr string) (exePath string, args []string) {
	uninstallStr = strings.TrimSpace(uninstallStr)
	if len(uninstallStr) == 0 {
		return "", nil
	}

	// 情况1: 整个字符串被引号包围
	if uninstallStr[0] == '"' {
		endQuote := strings.Index(uninstallStr[1:], `"`) + 1
		if endQuote > 0 {
			exePath = uninstallStr[1:endQuote]
			remaining := strings.TrimSpace(uninstallStr[endQuote+1:])
			if remaining != "" {
				args = splitArgs(remaining)
			}
			args = removeSpecialCharacters(args)
			return exePath, args
		}
	}

	// 情况2: 查找第一个参数开始的位置
	// 参数可能以 /, --, - 开头，或者是一个带引号的参数
	firstArgPos := -1
	inQuotes := false
	for i, r := range uninstallStr {
		switch {
		case r == '"':
			inQuotes = !inQuotes
		case !inQuotes && r == ' ':
			// 检查空格后的字符是否是参数开始
			if i+1 < len(uninstallStr) {
				nextChar := uninstallStr[i+1]
				if nextChar == '/' || nextChar == '-' {
					firstArgPos = i + 1
					break
				}
			}
		}
		if firstArgPos != -1 {
			break
		}
	}

	if firstArgPos > 0 {
		exePath = strings.TrimSpace(uninstallStr[:firstArgPos-1])
		args = splitArgs(uninstallStr[firstArgPos:])
	} else {
		exePath = uninstallStr
	}
	args = removeSpecialCharacters(args)
	return exePath, args
}

// splitArgs 分割参数字符串，处理带引号的参数
func splitArgs(s string) []string {
	var args []string
	var current strings.Builder
	inQuotes := false

	for _, r := range s {
		switch {
		case r == '"':
			inQuotes = !inQuotes
		case r == ' ' && !inQuotes:
			if current.Len() > 0 {
				args = append(args, current.String())
				current.Reset()
			}
		default:
			current.WriteRune(r)
		}
	}

	if current.Len() > 0 {
		args = append(args, current.String())
	}

	return args
}

// removeSpecialCharacters 参数去除特殊字符
func removeSpecialCharacters(data []string) []string {
	// 参数去除特殊字符
	argsReturn := make([]string, 0)
	for _, arg := range data {
		switch {
		case strings.Contains(arg, "/"):
			arg = strings.ReplaceAll(arg, "/", "")
			break
		case strings.Contains(arg, "\\"):
			arg = strings.ReplaceAll(arg, "\\", "")
			break
		case strings.Contains(arg, "-"):
			arg = strings.ReplaceAll(arg, "-", "")
			break
		}
		strings.TrimSpace(arg)
		argsReturn = append(argsReturn, arg)
	}
	return argsReturn
}
