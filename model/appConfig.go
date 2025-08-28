package model

import (
	"context"
	"encoding/json"
	"os"
	"winstore/comm"

	"github.com/wailsapp/wails/v2/pkg/runtime"
)

// Config 配置文件
type Config struct {
	// 软件版本
	Version string `json:"version"`
	// 更新地址
	UpdateUrl string `json:"updateUrl"`
}

// InitConfig 初始化配置文件 如果目录中不存在 config.json 文件则创建
func InitConfig(ctx context.Context) {
	_, err := os.Stat("./config.json")
	if err != nil {
		if os.IsNotExist(err) {
			// 创建配置文件
			file, err := os.Create("./config.json")
			if err != nil {
				runtime.LogErrorf(ctx, "创建配置文件失败：%s\n", err.Error())
				return
			}
			defer file.Close()

			// 创建默认配置
			config := Config{
				Version:   comm.AppVersion,
				UpdateUrl: comm.DefaultAppUpdateUrl,
			}

			// 将默认配置写入文件
			jsonData, err := json.Marshal(config)
			if err != nil {
				runtime.LogErrorf(ctx, "编译配置文件失败：%s\n", err.Error())
				return
			}
			_, err = file.Write(jsonData)
			if err != nil {
				runtime.LogErrorf(ctx, "写入配置文件失败：%s\n", err.Error())
				return
			}
			return
		}
		return
	}
}
