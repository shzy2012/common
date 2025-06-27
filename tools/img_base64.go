package tools

import (
	"encoding/base64"
	"log"
	"os"
)

func Img2Base64(imagePath string) string {
	imageData, err := os.ReadFile(imagePath)
	if err != nil {
		log.Printf("[img2base64]=>%s\n", err.Error())
		return ""
	}

	// 编码为 base64
	imgData := base64.StdEncoding.EncodeToString(imageData)
	// 图像 URL 或 base64 编码图像
	return "data:image/jpeg;base64," + imgData
}
