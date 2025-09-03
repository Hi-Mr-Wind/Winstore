package comm

import (
	"crypto/md5"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"io"
	"os"
	"os/exec"
	"strings"
)

// 执行WMIC命令获取硬件信息
func getWmicResult(alias, prop string) string {
	cmd := exec.Command("wmic", alias, "get", prop)
	output, err := cmd.Output()
	if err != nil {
		return "ERROR"
	}

	// 提取有效数据（跳过标题行和空白行）
	lines := strings.Split(string(output), "\n")
	for _, line := range lines {
		clean := strings.TrimSpace(line)
		if clean != "" && clean != prop {
			return clean
		}
	}
	return "EMPTY"
}

// 生成MD5哈希
func generateMD5(input string) string {
	hasher := md5.New()
	hasher.Write([]byte(input))
	return hex.EncodeToString(hasher.Sum(nil))
}

// GetMachineCode 获取机器码
func GetMachineCode() string {
	components := []string{
		getWmicResult("diskdrive", "SerialNumber"), // 硬盘序列号
		getWmicResult("baseboard", "SerialNumber"), // 主板序列号
		getWmicResult("cpu", "ProcessorId"),        // CPU ID
		getWmicResult("bios", "SerialNumber"),      // BIOS序列号
	}
	combined := strings.Join(components, "|")

	return generateMD5(combined)
}

// GetFileSHA256 获取文件哈希
func GetFileSHA256(filePath string) (string, error) {
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
