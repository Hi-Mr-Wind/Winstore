package apps

import (
	"context"
	"fmt"
	"strings"
	"winstore/model"
	"winstore/utils"

	"github.com/wailsapp/wails/v2/pkg/options"
	"github.com/wailsapp/wails/v2/pkg/runtime"
)

var secondInstanceArgs []string

// App struct
type App struct {
	ctx context.Context
}

// NewApp 创建新的 App 应用程序结构
func NewApp() *App {
	return &App{}
}

// Startup 启动时调用启动。保存上下文
// 因此，我们可以调用运行时方法
func (a *App) Startup(ctx context.Context) {
	a.ctx = ctx
	runtime.LogSetLogLevel(a.ctx, 2)
	model.DBinit(ctx)
}

// Shutdown 退出时调用
func (a *App) Shutdown(ctx context.Context) {
	runtime.LogInfo(a.ctx, "App shutdown")
}

func (a *App) OnSecondInstanceLaunch(secondInstanceData options.SecondInstanceData) {
	secondInstanceArgs = secondInstanceData.Args
	runtime.LogWarningf(a.ctx, "user opened second instance %s", strings.Join(secondInstanceData.Args, ","))
	runtime.LogWarningf(a.ctx, "user opened second from %s", secondInstanceData.WorkingDirectory)
	runtime.WindowUnminimise(a.ctx)
	runtime.Show(a.ctx)
	go runtime.EventsEmit(a.ctx, "launchArgs", secondInstanceArgs)
}

// Greet 返回给定名称的问候语
func (a *App) Greet(name string) string {
	return fmt.Sprintf("Hello %s,表演时间到了！", name)
}

// DownloadFile 下载文件
func (a *App) DownloadFile(url string, filepath string, filename string) error {
	return utils.DownloadFile(url, filepath, filename, a.ctx)
}
