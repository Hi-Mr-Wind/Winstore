package main

import (
	"archive/zip"
	"crypto/sha256"
	"encoding/json/v2"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"time"

	"github.com/hashicorp/go-version"
)

type Config struct {
	// 软件版本
	Version string `json:"version"`
	// 更新地址
	UpdateUrl string `json:"updateUrl"`
}

type Update struct {
	// 版本号
	Version string
	// 文件哈希
	Hash string
	// 下载地址
	downloadPath string
}

func main() {
	fmt.Println("Win Store 软件更新工具")
	fmt.Println("正在检查程序更新，请勿关闭此窗口...")
	checkUpdate()
}

func checkUpdate() {
	readFile, err := os.ReadFile("../config.json")
	if err != nil {
		fmt.Println("读取配置文件失败: ", err)
		return
	}
	var config Config
	err = json.Unmarshal(readFile, &config)
	if err != nil {
		fmt.Println("解析配置文件失败: ", err)
		return
	}
	get, err := http.Get(config.UpdateUrl)
	if err != nil {
		fmt.Printf("连接服务器获取更新数据失败：%s", err.Error())
		return
	}
	defer get.Body.Close()
	body, err := io.ReadAll(get.Body)
	if err != nil {
		fmt.Println("获取更新数据失败: ", err)
		return
	}
	var update Update
	err = json.Unmarshal(body, &update)
	if err != nil {
		fmt.Println("解析更新数据失败: ", err)
		return
	}
	currentVersion, _ := version.NewVersion(config.Version)
	newVersion, _ := version.NewVersion(update.Version)
	if currentVersion.LessThan(newVersion) {
		fmt.Println("检测到新版本，正在下载...")
		filePath := "../update.zip"
		err := downloadFile(update.downloadPath, filePath)
		if err != nil {
			fmt.Println("下载更新失败: ", err)
			return
		}
		fmt.Println("下载完成，正在验证更新...")
		fileSHA256, err := getFileSHA256(filePath)
		if err != nil {
			fmt.Println("获取文件哈希失败: ", err)
			return
		}
		if fileSHA256 != update.Hash {
			fmt.Println("文件哈希不匹配，请检查更新文件")
			return
		}
		err = unzip("../update.zip", "../")
		if err != nil {
			fmt.Println("解压更新失败: ", err)
			return
		}
		fmt.Println("更新完成，请重启win store 窗口将在5秒后重启")
		time.Sleep(5 * time.Second)
		os.Exit(0)
	} else {
		fmt.Println("当前版本为最新版本")
	}
}

// 下载文件
func downloadFile(url string, filepath string) error {
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	out, err := os.Create(filepath)
	if err != nil {
		return err
	}
	defer out.Close()
	_, err = io.Copy(out, resp.Body)
	return err
}

// 解压zip文件 并替换目录中的文件
func unzip(zipPath string, destPath string) error {
	zipReader, err := zip.OpenReader(zipPath)
	if err != nil {
		return err
	}
	defer zipReader.Close()

	for _, file := range zipReader.File {
		filePath := filepath.Join(destPath, file.Name)

		// 创建目录（如果是目录项）
		if file.FileInfo().IsDir() {
			if err := os.MkdirAll(filePath, os.ModePerm); err != nil {
				return err
			}
			continue
		}

		// 确保文件的父目录存在
		if err := os.MkdirAll(filepath.Dir(filePath), os.ModePerm); err != nil {
			return err
		}

		// 打开目标文件（如果存在则截断，不存在则创建）
		destFile, err := os.OpenFile(filePath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, file.Mode())
		if err != nil {
			return err
		}
		defer destFile.Close()

		// 打开 ZIP 中的文件
		zipFile, err := file.Open()
		if err != nil {
			return err
		}
		defer zipFile.Close()

		// 复制文件内容
		if _, err := io.Copy(destFile, zipFile); err != nil {
			return err
		}
	}

	return nil
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
