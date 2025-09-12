package tools

import (
	"fmt"
	"net/url"
	"testing"
)

func Test_GetParam(t *testing.T) {

	url, _ := url.Parse("https://example.org/x?a=1&a=2&b=3&=3&&&&")
	fmt.Printf("[a]=%s\n", URLParams(url, "a"))
	fmt.Printf("[b]=%s\n", URLParam(url, "b"))
	fmt.Printf("[c]=%s\n", URLParam(url, "c"))
}

func Test_URLCheck(t *testing.T) {
	rawURL := "https://www.baidu.com/s?wd=中文"
	processedURL, err := URLCheck(rawURL)
	if err != nil {
		fmt.Println("Error processing URL:", err)
		return
	}

	fmt.Println("Processed URL:", processedURL)
}
