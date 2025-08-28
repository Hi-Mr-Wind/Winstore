package sysUtils

import (
	"context"
	"crypto/sha256"
	"encoding/json/v2"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"winstore/model"

	"github.com/hashicorp/go-version"
)

// CheckUpdate 检查更新
func CheckUpdate(ctx context.Context) (string, error) {
	readFile, err := os.ReadFile("../config.json")
	if err != nil {
		fmt.Println("读取配置文件失败: ", err)
		return "", err
	}
	var config model.Config
	err = json.Unmarshal(readFile, &config)
	if err != nil {
		fmt.Println("解析配置文件失败: ", err)
		return "", err
	}
	get, err := http.Get(config.UpdateUrl)
	if err != nil {
		fmt.Printf("连接服务器获取更新数据失败：%s", err.Error())
		return "", err
	}
	defer get.Body.Close()
	body, err := io.ReadAll(get.Body)
	if err != nil {
		fmt.Println("获取更新数据失败: ", err)
		return "", err
	}
	var update model.UpdateInfo
	err = json.Unmarshal(body, &update)
	if err != nil {
		fmt.Println("解析更新数据失败: ", err)
		return "", err
	}
	currentVersion, _ := version.NewVersion(config.Version)
	newVersion, _ := version.NewVersion(update.Version)
	if currentVersion.LessThan(newVersion) {
		return fmt.Sprintf("检测到新版本：%s", update.Version), nil
	}
	return "软件已是最新版本", nil
}

// UpdateApp 更新程序
func UpdateApp(ctx context.Context) error {
	// todo
	return nil
}

// 下载文件
func downloadFile(url string, filePath string) error {
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	dir := filepath.Dir(filePath)
	if err := os.MkdirAll(dir, os.ModePerm); err != nil {
		return err
	}
	out, err := os.Create(filePath)
	if err != nil {
		return err
	}
	defer out.Close()
	_, err = io.Copy(out, resp.Body)
	return err
}

// 获取文件哈希
func getFileSHA256(filePath string) (string, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return "", err
	}
	defer file.Close()

	hash := sha256.New()
	if _, err := io.Copy(hash, file); err != nil {
		return "", err
	}
	return fmt.Sprintf("%x", hash.Sum(nil)), nil
}
