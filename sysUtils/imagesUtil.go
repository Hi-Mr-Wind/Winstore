package sysUtils

import (
	"fmt"
	"image"
	"image/color"
	"image/draw"
	"image/gif"
	"image/jpeg"
	"image/png"
	"os"
	"path/filepath"
	"strings"

	_ "image/gif"  // GIF解码器
	_ "image/jpeg" // JPEG解码器
	_ "image/png"  // PNG解码器

	"golang.org/x/image/bmp"
	"golang.org/x/image/tiff"

	_ "golang.org/x/image/bmp"  // BMP解码器
	_ "golang.org/x/image/tiff" // TIFF解码器
	_ "golang.org/x/image/webp" // WebP解码器
)

// ImageConversionType 图片格式转换
// filePath 图片路径
// newFilePath 新图片保存路径
// typeCode 图片格式码
func ImageConversionType(filePath string, newFilePath string, typeCode int) error {
	typeName := getType(typeCode)
	file, err := os.Open(filePath)
	if err != nil {
		return err
	}
	defer file.Close()
	// 根据文件后缀选择解码器
	ext := strings.ToLower(filepath.Ext(filePath))
	img, format, err := decodeImage(file, ext)
	if err != nil {
		fmt.Println("图像解码错误:", err)
		return err
	}
	// 处理格式转换
	newImg, err := convertImage(img, format, typeName)
	if err != nil {
		fmt.Println("图像转换错误:", err)
		return err
	}
	name := extractName(filePath)
	newFilePath = filepath.Join(newFilePath, name+"."+typeName)
	outFile, err := os.Create(newFilePath)
	if err != nil {
		fmt.Println("输出文件创建失败:", err)
		return err
	}
	defer outFile.Close()

	return encodeImage(outFile, newImg, typeName)
}

// decodeImage 根据文件后缀选择解码器
func decodeImage(file *os.File, ext string) (image.Image, string, error) {
	// 尝试自动检测
	img, format, err := image.Decode(file)
	if err != nil {
		return nil, "", fmt.Errorf("unsupported image format: %s", ext)
	}
	return img, format, nil
}

// extractName 从文件路径中提取文件名
func extractName(imgPath string) string {
	filename := filepath.Base(imgPath)
	filenameWithoutExt := filename[:len(filename)-len(filepath.Ext(filename))]
	return filenameWithoutExt
}

// convertImage 处理图片格式转换的特殊逻辑
func convertImage(img image.Image, srcFormat string, dstFormat string) (image.Image, error) {
	// 需要移除透明通道的情况：源图像有透明通道且目标格式不支持透明
	hasAlpha := hasAlphaChannel(img)
	needsOpaque := hasAlpha && (dstFormat == "jpg" || dstFormat == "jpeg" || dstFormat == "bmp")
	if needsOpaque {
		return removeAlphaChannel(img), nil
	}
	// GIF特殊处理：保留原始GIF
	if srcFormat == "gif" && dstFormat == "gif" {
		return img, nil
	}
	return img, nil
}

// hasAlphaChannel 检测图像是否有透明通道
func hasAlphaChannel(img image.Image) bool {
	switch img.(type) {
	case *image.NRGBA, *image.NRGBA64, *image.RGBA, *image.RGBA64:
		return true
	case *image.Paletted:
		// 检查调色板中是否有透明颜色
		paletted := img.(*image.Paletted)
		for _, c := range paletted.Palette {
			if _, _, _, a := c.RGBA(); a != 0xffff {
				return true
			}
		}
	}
	return false
}

// removeAlphaChannel 移除透明通道（用白色背景替换透明区域）
func removeAlphaChannel(img image.Image) *image.RGBA {
	bounds := img.Bounds()
	rgba := image.NewRGBA(bounds)
	draw.Draw(rgba, bounds, image.NewUniform(color.White), image.Point{}, draw.Src)
	draw.Draw(rgba, bounds, img, bounds.Min, draw.Over)

	return rgba
}

// encodeImage 根据目标格式编码图片
func encodeImage(outFile *os.File, img image.Image, format string) error {
	switch strings.ToLower(format) {
	case "jpg", "jpeg":
		return jpeg.Encode(outFile, img, &jpeg.Options{Quality: 100})
	case "png":
		return png.Encode(outFile, img)
	case "gif":
		return gif.Encode(outFile, img, nil)
	case "bmp":
		return bmp.Encode(outFile, img)
	case "tiff", "tif":
		return tiff.Encode(outFile, img, &tiff.Options{Compression: tiff.Deflate})
	case "webp":
		return fmt.Errorf("webp encoding not implemented (requires external library)")
	default:
		return fmt.Errorf("unsupported target format: %s", format)
	}
}

// getType 根据类型码获取图片格式
func getType(typeCode int) string {
	switch typeCode {
	case 0:
		return "jpg"
	case 1:
		return "png"
	case 2:
		return "gif"
	case 3:
		return "bmp"
	case 4:
		return "tif"
	case 5:
		return "webp"
	default:
		return "unknown"
	}
}
