package sysUtils

import (
	"fmt"
	"syscall"
	"unsafe"
)

const (
	MovefileDelayUntilReboot = 0x4
)

var (
	modkernel32                  = syscall.NewLazyDLL("kernel32.dll")
	procMoveFileEx               = modkernel32.NewProc("MoveFileExW")
	advapi32                     = syscall.NewLazyDLL("advapi32.dll")
	procInitiateSystemShutdownEx = advapi32.NewProc("InitiateSystemShutdownExW")
)

// MarkForDeletion 标记删除指定文件夹，在删除时，会等待系统重启后才会删除
func MarkForDeletion(path string) error {
	pathPtr, err := syscall.UTF16PtrFromString(path)
	if err != nil {
		return err
	}

	ret, _, err := procMoveFileEx.Call(
		uintptr(unsafe.Pointer(pathPtr)),
		0,
		MovefileDelayUntilReboot,
	)
	if ret == 0 {
		return fmt.Errorf("MoveFileEx failed: %v", err)
	}
	return nil
}

// RestartSystem 带有提示的重启计算机
// 例如 RestartSystem("软件清理即将重启您的计算机")
func RestartSystem(reason string) error {
	reasonPtr, _ := syscall.UTF16PtrFromString(reason)
	ret, _, err := procInitiateSystemShutdownEx.Call(
		uintptr(0),
		uintptr(0),
		uintptr(0),
		uintptr(1),
		uintptr(1),
		uintptr(unsafe.Pointer(reasonPtr)),
	)
	if ret == 0 {
		return fmt.Errorf("InitiateSystemShutdownEx failed: %v", err)
	}
	return nil
}
