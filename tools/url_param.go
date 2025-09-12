package tools

import (
	"fmt"
	"net/url"
	"strings"
)

// GetParam 获取Url的值
func URLParam(url *url.URL, name string) string {
	keys, ok := url.Query()[name]
	if !ok || len(keys[0]) < 1 {
		return ""
	}
	return keys[0]
}

// GetParams 获取Url的值
func URLParams(url *url.URL, name string) []string {
	return url.Query()[name]
}

// 检查、并修复 URL
func URLCheck(rawURL string) (string, error) {
	// 检查并补充协议
	if !strings.HasPrefix(rawURL, "http://") && !strings.HasPrefix(rawURL, "https://") {
		rawURL = "http://" + rawURL
	}

	// 解析 URL
	parsedURL, err := url.Parse(rawURL)
	if err != nil {
		return "", fmt.Errorf("error parsing URL: %v", err)
	}

	// 对查询参数进行编码
	queryParams := parsedURL.Query()
	parsedURL.RawQuery = queryParams.Encode()

	// 返回处理后的 URL 字符串
	return parsedURL.String(), nil
}
