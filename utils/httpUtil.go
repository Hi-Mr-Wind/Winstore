package utils

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"os"
	"time"
	"winstore/model"

	"github.com/wailsapp/wails/v2/pkg/runtime"
)

func GetHttp(ctx context.Context, url string) (string, error) {
	resp, err := http.Get(url)
	if err != nil {
		return "", err
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			runtime.LogError(ctx, err.Error())
		}
	}(resp.Body)
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	return string(body), nil
}

// DownloadFile 下载文件并向事件发送进度
func DownloadFile(url string, filepath string, filename string, appContext context.Context) error {
	runtime.LogInfo(appContext, "开始下载文件"+url)
	sprintf := fmt.Sprintf("%s//%s", filepath, filename)
	// 创建输出文件
	out, err := os.Create(sprintf)
	if err != nil {
		return err
	}
	defer out.Close()

	// 发送HTTP请求
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// 检查响应状态
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("server returned non-200 status: %s", resp.Status)
	}

	// 获取文件大小
	totalSize := resp.ContentLength
	if totalSize == -1 {
		// 未知大小，不显示百分比
		totalSize = 0
	}
	buf := make([]byte, 32*1024) // 32KB缓冲区
	var downloaded int64
	var bytesRead int
	lastUpdate := time.Now()
	var lastDownloaded int64
	progress := new(model.DownloadProgress)
	for {
		// 读取数据
		bytesRead, err = resp.Body.Read(buf)
		if bytesRead > 0 {
			// 写入文件
			_, writeErr := out.Write(buf[:bytesRead])
			if writeErr != nil {
				return writeErr
			}

			// 更新下载量
			downloaded += int64(bytesRead)

			// 计算下载速度（每秒）
			now := time.Now()
			elapsed := now.Sub(lastUpdate).Seconds()
			if elapsed >= 0.1 { // 每0.5秒更新一次速度
				// 发送进度更新
				speed := float64(downloaded-lastDownloaded) / elapsed / 1024 // KB/s
				lastUpdate = now
				lastDownloaded = downloaded
				tot := float64(totalSize) / 1024
				//已下载大小
				down := float64(downloaded) / 1024
				//格式化已下载大小
				if down >= 1024 {
					down = down / 1024
					progress.Downloaded = fmt.Sprintf("%.2f MB", down)
				} else {
					progress.Downloaded = fmt.Sprintf("%.2f KB", down)
				}
				//格式化总大小
				if tot >= 1024 {
					tot = tot / 1024
					progress.Total = fmt.Sprintf("%.2f MB", tot)
				} else {
					progress.Total = fmt.Sprintf("%.2f KB", tot)
				}
				//格式化速度
				if speed >= 1024 {
					speed = speed / 1024
					progress.Speed = fmt.Sprintf("%.2f MB/s", speed)
				} else {
					progress.Speed = fmt.Sprintf("%.2f KB/s", speed)
				}

				if totalSize > 0 {
					progress.Percentage = float64(downloaded) / float64(totalSize) * 100
				}
				//发送下载进度
				baseModel := new(model.BaseModel)
				baseModel.Code = 0
				baseModel.Data = progress
				baseModel.Message = "success"
				runtime.EventsEmit(appContext, "download", baseModel.New())
			}
		}
		// 检查错误
		if err != nil {
			if err == io.EOF {
				break // 下载完成
			}
			return err
		}
	}
	progress.Percentage = 100
	progress.Downloaded = progress.Total
	baseModel := new(model.BaseModel)
	baseModel.Code = 0
	baseModel.Data = progress
	baseModel.Message = "success"
	runtime.EventsEmit(appContext, "download", baseModel.New())
	return nil
}
