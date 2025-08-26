package main

import (
	"embed"
	"fmt"
	"log"
	"winstore/apps"

	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/logger"
	"github.com/wailsapp/wails/v2/pkg/options"
	"github.com/wailsapp/wails/v2/pkg/options/assetserver"
)

//go:embed all:frontend/dist
var assets embed.FS

func main() {
	// 创建应用结构的实例
	app := apps.NewApp()
	// 重定向标准输出到 Wails 日志
	log.SetOutput(&WailsLogWriter{})
	// 创建带有选项的应用程序
	err := wails.Run(&options.App{
		Title:    "win store",
		Width:    1400,
		Height:   980,
		LogLevel: logger.DEBUG,
		//Logger:   logger.NewFileLogger("app.log"), // 可选：同时输出到文件
		AssetServer: &assetserver.Options{
			Assets: assets,
		},
		BackgroundColour: &options.RGBA{R: 255, G: 255, B: 255, A: 0},
		OnStartup:        app.Startup,
		OnShutdown:       app.Shutdown,
		SingleInstanceLock: &options.SingleInstanceLock{
			UniqueId:               "0bc864df-9fda-8d78-e52c-b9435aedd462",
			OnSecondInstanceLaunch: app.OnSecondInstanceLaunch,
		},
		Bind: []interface{}{
			app,
		},
	})
	if err != nil {
		println("Error:", err.Error())
	}
}

type WailsLogWriter struct{}

func (w *WailsLogWriter) Write(p []byte) (n int, err error) {
	// 这里可以通过前端显示日志，或者使用 wails 的日志系统
	fmt.Print(string(p))
	return len(p), nil
}
