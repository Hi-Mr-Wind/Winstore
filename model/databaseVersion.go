package model

import (
	"context"
	"encoding/json/v2"
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"
	"time"
	"winstore/comm"

	rtime "runtime"

	"github.com/wailsapp/wails/v2/pkg/runtime"
)

// DatabaseVersion 数据库版本
type DatabaseVersion struct {
	Version    int64  `gorm:"column:version;type:int;not null"                json:"version"`
	UpdateTime string `gorm:"column:update_time;type:text;not null"          json:"update_time"`
}

func (*DatabaseVersion) TableName() string {
	return "database_version"
}

// GetVersion 获取数据库版本
func (databaseVersion *DatabaseVersion) GetVersion() error {
	err := DB.Where("1=1").Find(databaseVersion).Error
	return err
}

// SelectUpdateDataBase 查询软件数据库更新
func (databaseVersion *DatabaseVersion) SelectUpdateDataBase(ctx context.Context) (bool, error) {
	config := new(Config)
	if err := config.GetConfig(); err != nil {
		return false, err
	}
	if err := databaseVersion.GetVersion(); err != nil {
		return false, err
	}
	client := &http.Client{
		Timeout: 30 * time.Second,
	}
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/database/version", config.UpdateUrl), nil)
	if err != nil {
		runtime.LogErrorf(ctx, "创建请求失败:%s/n", err)
	}
	req.Header.Set("Authorization", comm.GetMachineCode())
	// 发送请求
	resp, err := client.Do(req)
	if err != nil {
		runtime.LogErrorf(ctx, "请求数据失败: %v\n", err)
	}

	defer func() {
		if closeErr := resp.Body.Close(); closeErr != nil {
			fmt.Printf("关闭响应体时出错: %v\n", closeErr)
		}
	}()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return false, fmt.Errorf("读取响应体失败: %w", err)
	}

	var data BaseModel
	if err := json.Unmarshal(body, &data); err != nil {
		return false, fmt.Errorf("解析响应数据失败: %w", err)
	}
	if data.Code != 0 {
		return false, fmt.Errorf("获取数据库更新失败")
	}

	ver, err := strconv.ParseInt(data.Data.(string), 10, 64)
	if err != nil {
		return false, err
	}
	if ver <= databaseVersion.Version {
		return false, nil
	}
	return true, nil
}

// DownloadNewDataBase 下载数据库更新
func DownloadNewDataBase(ctx context.Context) error {
	config := new(Config)
	err := config.GetConfig()
	if err != nil {
		return err
	}
	// 创建一个httpclient 包含请求头和30秒超时时间
	client := &http.Client{
		Timeout: 30 * time.Second,
	}
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/database/download", config.UpdateUrl), nil)
	if err != nil {
		runtime.LogErrorf(ctx, "创建请求失败:%s/n", err)
	}
	req.Header.Set("Authorization", comm.GetMachineCode())
	// 发送请求
	resp, err := client.Do(req)
	if err != nil {
		runtime.LogErrorf(ctx, "请求数据失败: %v\n", err)
	}
	defer func() {
		if closeErr := resp.Body.Close(); closeErr != nil {
			runtime.LogErrorf(ctx, "关闭响应体时出错: %v\n", closeErr)
		}
	}()
	// 验证http返回码
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("下载软件数据库失败: %s", resp.Status)
	}
	// 获取body中的字符串地址
	downloadUrl, err := io.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("解析响应数据失败: %w", err)
	}
	// 根据downloadUrl请求并下载数据库文件
	request, err := http.NewRequest("GET", string(downloadUrl), nil)
	if err != nil {
		return err
	}
	request.Header.Set("Authorization", comm.GetMachineCode())
	get, err := client.Do(request)
	if err != nil {
		return fmt.Errorf("请求数据库文件失败: %w", err)
	}
	defer func() {
		if closeErr := get.Body.Close(); closeErr != nil {
			runtime.LogErrorf(ctx, "关闭响应时出错: %v\n", closeErr)
		}
	}()
	runtime.LogInfof(ctx, "开始下载数据库文件")
	// 验证http返回码
	if get.StatusCode != http.StatusOK {
		return fmt.Errorf("下载软件数据库失败: %s", get.Status)
	}
	//从请求中获取文件
	file, err := os.Create("./winstore.db")
	if err != nil {
		return fmt.Errorf("创建文件失败: %w", err)
	}
	_, err = io.Copy(file, get.Body)
	if err != nil {
		return fmt.Errorf("写入文件失败: %w", err)
	}
	// 关闭原数据库连接
	db, err := DB.DB()
	if err != nil {
		return err
	}
	if err := db.Close(); err != nil {
		return err
	}
	DB = nil
	rtime.GC()
	// 验证数据连接是否以及被关闭，如果没有则等待关闭，最大等待时间为60秒
	for i := 0; i < 60; i++ {
		if db.Stats().OpenConnections == 0 {
			break
		}
		time.Sleep(1 * time.Second)
	}

	//检查备份目录是否存在，如不存在则创建
	if _, err := os.Stat("./lib/backups"); os.IsNotExist(err) {
		err := os.Mkdir("./lib/backups", os.ModePerm)
		if err != nil {
			return err
		}
	}
	// 关闭文件句柄
	if closeErr := file.Close(); closeErr != nil {
		runtime.LogErrorf(ctx, "关闭文件时出错: %v\n", closeErr)
	}
	// 将原数据库文件移动到备份目录
	if err := os.Rename("./lib/winstore.db", "./lib/backups/winstore.db.bak"); err != nil {
		return err
	}
	//将新数据库文件移动到原数据库文件位置
	if err := os.Rename("./winstore.db", "./lib/winstore.db"); err != nil {
		return err
	}
	//重新连接数据库
	DBinit(ctx)
	//验证数据库连接
	d, err := DB.DB()
	if err != nil {
		return err
	}
	if err := d.Ping(); err != nil {
		return err
	}
	return nil
}
