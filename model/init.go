package model

import (
	"context"
	"syscall"
	"unsafe"
	"winstore/log"

	"github.com/glebarez/sqlite"
	"github.com/wailsapp/wails/v2/pkg/runtime"
	"gorm.io/gorm"
)

var DB *gorm.DB

// DBinit 初始化数据库
func DBinit(c context.Context) {
	gormLogger := log.NewCustomGormLogger(c)
	db, err := gorm.Open(sqlite.Open("./lib/winstore.db"), &gorm.Config{
		Logger: gormLogger,
	})
	if err != nil {
		runtime.LogErrorf(c, "数据加载失败：%s", err.Error())
		user32 := syscall.NewLazyDLL("user32.dll")
		messageBox := user32.NewProc("MessageBoxW")
		fromString, err := syscall.UTF16PtrFromString("程序数据库加载失败！请检查程序安装目录下lib文件夹内是否存在数据文件")
		if err != nil {
			return
		}
		lpCaption, err := syscall.UTF16PtrFromString("错误！")
		if err != nil {
			return
		}
		ret, _, _ := messageBox.Call(
			0,
			uintptr(unsafe.Pointer(fromString)),
			uintptr(unsafe.Pointer(lpCaption)),
			0,
		)
		if ret == 0 {
			println("调用 MessageBox 失败")
		}
		runtime.Quit(c)
	}
	DB = db
}
