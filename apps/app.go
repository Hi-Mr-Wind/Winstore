package apps

import (
	"context"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"winstore/comm"
	"winstore/model"
	"winstore/sysUtils"

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
	//获取用户缓存目录
	userCacheDir, oserr := os.UserCacheDir()
	if oserr != nil {
		dir := "./cache"
		// 如果没有获取到用户目录则安装路径下创建缓存目录
		err := os.MkdirAll(dir, os.ModePerm)
		if err != nil {
			panic(err)
		}
		comm.UserDir = dir
	} else {
		cacheDir := filepath.Join(userCacheDir, "winstore")
		err := os.MkdirAll(cacheDir, os.ModePerm)
		if err != nil {
			panic(err)
		}
		comm.UserDir = cacheDir
	}
	// 初始化缓存
	comm.NewCache()
	return &App{}
}

// Startup 启动时调用启动。保存上下文
// 因此，我们可以调用运行时方法
func (a *App) Startup(ctx context.Context) {
	a.ctx = ctx
	runtime.LogSetLogLevel(a.ctx, 2)
	// 初始化数据库
	model.DBinit(ctx)
	// 初始化软件配置
	model.InitConfig(ctx)
	residue := new(model.Residue)
	all := residue.SelectAll()
	if all != nil {
		// 遍历检查残余数据
		for _, r := range *all {
			// 判断残余目录是否存在，如果已经不存在，则删除残余数据
			_, err := os.Stat(r.ResiduePath)
			if err != nil {
				if os.IsNotExist(err) {
					err := r.DeleteByName()
					if err != nil {
						runtime.LogError(a.ctx, "删除软件残留数据失败！")
						return
					}
				}
				return
			}
		}
	}
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
	return sysUtils.DownloadFile(url, filepath, filename, a.ctx)
}
